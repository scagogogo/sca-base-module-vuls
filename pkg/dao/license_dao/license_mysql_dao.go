package license_dao

import (
	"context"
	"github.com/scagogogo/sca-base-module-dao/mysql"
	"github.com/scagogogo/sca-base-module-vuls/pkg/domain"
)

type LicenseMysqlDao struct {
}

var _ LicenseDao = &LicenseMysqlDao{}

func NewLicenseMysqlDao() *LicenseMysqlDao {
	return &LicenseMysqlDao{}
}

func (x *LicenseMysqlDao) Create(ctx context.Context, license *domain.License) error {
	//TODO implement me
	panic("implement me")
}

func (x *LicenseMysqlDao) Update(ctx context.Context, license *domain.License) error {
	//TODO implement me
	panic("implement me")
}

func (x *LicenseMysqlDao) Upsert(ctx context.Context, license *domain.License) error {
	//TODO implement me
	panic("implement me")
}

func (x *LicenseMysqlDao) Find(ctx context.Context, identifier string) (*domain.License, error) {
	//TODO implement me
	panic("implement me")
}

func (x *LicenseMysqlDao) Delete(ctx context.Context, identifier string) error {
	//TODO implement me
	panic("implement me")
}

func (x *LicenseMysqlDao) LoadAllLicenses(ctx context.Context) ([]*domain.License, error) {
	//TODO implement me
	panic("implement me")
}

// FindLicense 根据license名称查询相关设置
func FindLicense(ctx context.Context, identifier string) (*domain.License, error) {
	var r *domain.License
	err := mysql.Gorm.Model(&r).Where("identifier = ?", identifier).Scan(&r).Error
	if err != nil {
		return nil, err
	}
	return r, nil
}

// SaveLicense 保存license
func SaveLicense(ctx context.Context, license *domain.License) error {
	return mysql.Gorm.Model(&license).Save(license).Error
}

// LoadAllLicenses 加载所有的license
func LoadAllLicenses(ctx context.Context) ([]*domain.License, error) {
	var slice []*domain.License
	err := mysql.Gorm.Model(&domain.License{}).Scan(&slice).Error
	if err != nil {
		return nil, err
	}
	return slice, nil
}
