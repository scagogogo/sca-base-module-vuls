package license_dao

import (
	"context"
	"github.com/scagogogo/sca-base-module-dao/pkg/mysql"

	"github.com/scagogogo/sca-base-module-vuls/pkg/models"
)

type LicenseMysqlDao struct {
}

var _ LicenseDao = &LicenseMysqlDao{}

func NewLicenseMysqlDao() *LicenseMysqlDao {
	return &LicenseMysqlDao{}
}

func (x *LicenseMysqlDao) Create(ctx context.Context, license *models.License) error {
	//TODO implement me
	panic("implement me")
}

func (x *LicenseMysqlDao) Update(ctx context.Context, license *models.License) error {
	//TODO implement me
	panic("implement me")
}

func (x *LicenseMysqlDao) Upsert(ctx context.Context, license *models.License) error {
	//TODO implement me
	panic("implement me")
}

func (x *LicenseMysqlDao) Find(ctx context.Context, identifier string) (*models.License, error) {
	//TODO implement me
	panic("implement me")
}

func (x *LicenseMysqlDao) Delete(ctx context.Context, identifier string) error {
	//TODO implement me
	panic("implement me")
}

func (x *LicenseMysqlDao) LoadAllLicenses(ctx context.Context) ([]*models.License, error) {
	//TODO implement me
	panic("implement me")
}

// FindLicense 根据license名称查询相关设置
func FindLicense(ctx context.Context, identifier string) (*models.License, error) {
	var r *models.License
	err := mysql.Gorm.Model(&r).Where("identifier = ?", identifier).Scan(&r).Error
	if err != nil {
		return nil, err
	}
	return r, nil
}

// SaveLicense 保存license
func SaveLicense(ctx context.Context, license *models.License) error {
	return mysql.Gorm.Model(&license).Save(license).Error
}

// LoadAllLicenses 加载所有的license
func LoadAllLicenses(ctx context.Context) ([]*models.License, error) {
	var slice []*models.License
	err := mysql.Gorm.Model(&models.License{}).Scan(&slice).Error
	if err != nil {
		return nil, err
	}
	return slice, nil
}
