package license_dao

import (
	"context"
	"github.com/scagogogo/sca-base-module-vuls/pkg/domain"
)

// LicenseDao license的dao
type LicenseDao interface {

	// Create 新建一条license
	Create(ctx context.Context, license *domain.License) error

	// Update 更新license的信息
	Update(ctx context.Context, license *domain.License) error

	// Upsert 如果license存在则更新，否则新建
	Upsert(ctx context.Context, license *domain.License) error

	// Find 查询给定id的license的信息
	Find(ctx context.Context, identifier string) (*domain.License, error)

	// Delete 根据给定的ID删除license
	Delete(ctx context.Context, identifier string) error

	// LoadAllLicenses 加载所有的license
	LoadAllLicenses(ctx context.Context) ([]*domain.License, error)
}
