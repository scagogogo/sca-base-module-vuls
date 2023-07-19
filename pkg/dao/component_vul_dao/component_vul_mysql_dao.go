package component_vul_dao

import (
	"context"
	"github.com/scagogogo/sca-base-module-vuls/pkg/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// ComponentVulMysqlDao mysql实现
type ComponentVulMysqlDao struct {
	gorm *gorm.DB
}

var _ ComponentVulDao = &ComponentVulMysqlDao{}

func NewComponentVulMysqlDao(gorm *gorm.DB) *ComponentVulMysqlDao {
	return &ComponentVulMysqlDao{
		gorm: gorm,
	}
}

func (x *ComponentVulMysqlDao) Create(ctx context.Context, cv *models.ComponentVul) error {
	return x.gorm.WithContext(ctx).Model(&cv).Create(cv).Error
}

func (x *ComponentVulMysqlDao) Update(ctx context.Context, cv *models.ComponentVul) error {
	return x.gorm.WithContext(ctx).Model(&cv).Where("name = ? AND version = ?", cv.Name, cv.Version).Save(cv).Error
}

func (x *ComponentVulMysqlDao) Upsert(ctx context.Context, cv *models.ComponentVul) error {
	return x.gorm.WithContext(ctx).Model(&cv).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "cve"}},
		DoUpdates: clause.Assignments(map[string]any{
			"update_time": cv.UpdateTime,
		}),
	}).Create(&cv).Error
}

func (x *ComponentVulMysqlDao) FindByComponentName(ctx context.Context, componentName string) ([]*models.ComponentVul, error) {
	var slice []*models.ComponentVul
	err := x.gorm.WithContext(ctx).Model(&models.ComponentVul{}).Where("name = ?", componentName).Scan(&slice).Error
	if err != nil {
		return nil, err
	} else {
		return slice, nil
	}
}

func (x *ComponentVulMysqlDao) Find(ctx context.Context, componentName, componentVersion string) ([]*models.ComponentVul, error) {
	var slice []*models.ComponentVul
	err := x.gorm.WithContext(ctx).Model(&models.ComponentVul{}).Where("name = ? AND version = ?", componentName, componentVersion).Scan(&slice).Error
	if err != nil {
		return nil, err
	} else {
		return slice, nil
	}
}

func (x *ComponentVulMysqlDao) DeleteByVulId(ctx context.Context, vulId string) (int64, error) {
	tx := x.gorm.WithContext(ctx).Where("vul_id = ?", vulId).Delete(&models.ComponentVul{})
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func (x *ComponentVulMysqlDao) DeleteByComponentName(ctx context.Context, componentName string) (int64, error) {
	tx := x.gorm.WithContext(ctx).Where("name = ?", componentName).Delete(&models.ComponentVul{})
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func (x *ComponentVulMysqlDao) DeleteByComponentNameAndVersion(ctx context.Context, componentName, componentVersion string) (int64, error) {
	tx := x.gorm.WithContext(ctx).Where("name = ? AND version = ?", componentName, componentVersion).Delete(&models.ComponentVul{})
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func (x *ComponentVulMysqlDao) LoadAll(ctx context.Context) ([]*models.ComponentVul, error) {
	var slice []*models.ComponentVul
	err := x.gorm.WithContext(ctx).Model(&models.ComponentVul{}).Scan(&slice).Error
	if err != nil {
		return nil, err
	}
	return slice, nil
}
