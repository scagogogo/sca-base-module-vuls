package component_vul_dao

import (
	"context"
	"github.com/scagogogo/sca-base-module-vuls/pkg/models"
)

// ComponentVulDao 组件漏洞dao
type ComponentVulDao interface {

	// Create 保存组件漏洞信息
	Create(ctx context.Context, cv *models.ComponentVul) error

	// Update 更新组件漏洞
	Update(ctx context.Context, cv *models.ComponentVul) error

	// Upsert 组件漏洞存在的话则更新，否则插入
	Upsert(ctx context.Context, cv *models.ComponentVul) error

	// FindByComponentName 根据组件名称查询上面的所有漏洞
	FindByComponentName(ctx context.Context, componentName string) ([]*models.ComponentVul, error)

	// Find 根据组件名字和版本查询漏洞
	Find(ctx context.Context, componentName, componentVersion string) ([]*models.ComponentVul, error)

	// DeleteByVulId 根据漏洞ID删除组件漏洞
	DeleteByVulId(ctx context.Context, vulId string) (int64, error)

	// DeleteByComponentName 根据组件的名字删除组件漏洞
	DeleteByComponentName(ctx context.Context, componentName string) (int64, error)

	// DeleteByComponentNameAndVersion 根据组件的名字和版本删除组件漏洞
	// return:
	// int64: 被删除的数据条数
	// error: if has error
	DeleteByComponentNameAndVersion(ctx context.Context, componentName, componentVersion string) (int64, error)

	// LoadAll 加载所有的组件漏洞
	LoadAll(ctx context.Context) ([]*models.ComponentVul, error)
}
