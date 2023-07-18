package component_vul_dao

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/scagogogo/sca-base-module-vuls/pkg/domain"
)

type ComponentVulMemoryDao struct {
	// map[组件name]map[组件版本]map[漏洞id]*domain.ComponentVul
	db map[string]map[string]map[string]*domain.ComponentVul
}

var _ ComponentVulDao = &ComponentVulMemoryDao{}

// NewComponentVulMemoryDaoFromJsonLine 从jsonline文件中创建
func NewComponentVulMemoryDaoFromJsonLine(ctx context.Context, jsonLineBytes []byte) (*ComponentVulMemoryDao, error) {
	x := &ComponentVulMemoryDao{
		db: make(map[string]map[string]map[string]*domain.ComponentVul, 0),
	}
	split := bytes.Split(jsonLineBytes, []byte("\n"))
	for _, lineBytes := range split {
		// 空行忽略
		if len(lineBytes) == 0 {
			continue
		}
		r := &domain.ComponentVul{}
		err := json.Unmarshal(lineBytes, &r)
		if err != nil {
			return nil, err
		}
		err = x.Upsert(ctx, r)
		if err != nil {
			return nil, err
		}
	}
	return x, nil
}

func (x *ComponentVulMemoryDao) ensureNameVulExists(name string) map[string]map[string]*domain.ComponentVul {
	nameVuls, exists := x.db[name]
	if !exists {
		nameVuls = make(map[string]map[string]*domain.ComponentVul)
		x.db[name] = nameVuls
	}
	return nameVuls
}

func (x *ComponentVulMemoryDao) ensureVersionVulExists(name, version string) map[string]*domain.ComponentVul {
	nameVuls := x.ensureNameVulExists(name)
	versionVuls, exists := nameVuls[version]
	if !exists {
		versionVuls = make(map[string]*domain.ComponentVul, 0)
		nameVuls[version] = versionVuls
	}
	return versionVuls
}

func (x *ComponentVulMemoryDao) Create(ctx context.Context, cv *domain.ComponentVul) error {
	versionVuls := x.ensureVersionVulExists(cv.Name, cv.Version)
	_, exists := versionVuls[cv.VulId]
	if exists {
		return fmt.Errorf("%s vul id %s already exists", cv.BuildComponentId(), cv.VulId)
	}
	versionVuls[cv.VulId] = cv
	return nil
}

func (x *ComponentVulMemoryDao) Update(ctx context.Context, cv *domain.ComponentVul) error {
	versionVuls := x.ensureVersionVulExists(cv.Name, cv.Version)
	_, exists := versionVuls[cv.VulId]
	if !exists {
		//return fmt.Errorf("%s vul id %s not exists, can not update it", cv.BuildComponentId(), cv.VulId)
		return nil
	}
	versionVuls[cv.VulId] = cv
	return nil
}

func (x *ComponentVulMemoryDao) Upsert(ctx context.Context, cv *domain.ComponentVul) error {
	versionVuls := x.ensureVersionVulExists(cv.Name, cv.Version)
	versionVuls[cv.VulId] = cv
	return nil
}

func (x *ComponentVulMemoryDao) FindByComponentName(ctx context.Context, componentName string) ([]*domain.ComponentVul, error) {
	nameVuls := x.ensureNameVulExists(componentName)
	slice := make([]*domain.ComponentVul, 0)
	for _, versionVuls := range nameVuls {
		for _, vul := range versionVuls {
			slice = append(slice, vul)
		}
	}
	return slice, nil
}

func (x *ComponentVulMemoryDao) Find(ctx context.Context, componentName, componentVersion string) ([]*domain.ComponentVul, error) {
	versionVuls := x.ensureVersionVulExists(componentName, componentVersion)
	slice := make([]*domain.ComponentVul, 0)
	for _, vul := range versionVuls {
		slice = append(slice, vul)
	}
	return slice, nil
}

func (x *ComponentVulMemoryDao) DeleteByVulId(ctx context.Context, vulId string) (int64, error) {
	c := 0
	for _, nameVuls := range x.db {
		for _, versionVuls := range nameVuls {
			_, exists := versionVuls[vulId]
			if exists {
				c++
			}
			delete(versionVuls, vulId)
		}
	}
	return int64(c), nil
}

func (x *ComponentVulMemoryDao) DeleteByComponentName(ctx context.Context, componentName string) (int64, error) {
	nameVuls := x.ensureNameVulExists(componentName)

	c := 0
	for _, versionVuls := range nameVuls {
		c += len(versionVuls)
	}

	delete(x.db, componentName)

	return int64(c), nil
}

func (x *ComponentVulMemoryDao) DeleteByComponentNameAndVersion(ctx context.Context, componentName, componentVersion string) (int64, error) {

	versionVuls := x.ensureVersionVulExists(componentName, componentVersion)

	c := len(versionVuls)

	delete(x.db[componentName], componentVersion)

	return int64(c), nil
}

func (x *ComponentVulMemoryDao) LoadAll(ctx context.Context) ([]*domain.ComponentVul, error) {
	slice := make([]*domain.ComponentVul, 0)
	for _, nameVul := range x.db {
		for _, versionVul := range nameVul {
			for _, vul := range versionVul {
				slice = append(slice, vul)
			}
		}
	}
	return slice, nil
}
