package vul_dao

import (
	"context"
	"github.com/scagogogo/sca-base-module-vuls/pkg/models"
)

type VulMemoryDao struct {

	// 漏洞
	db map[string]*models.Vul

	// map[漏洞ID]map[编号ID]models.VulCode
	codes map[string]map[string]*models.VulCode
}

var _ VulDao = &VulMemoryDao{}

func NewVulMemoryDao() *VulMemoryDao {
	return &VulMemoryDao{}
}

func (x *VulMemoryDao) Create(ctx context.Context, vul *models.Vul, codes []*models.VulCode) error {
	//TODO implement me
	panic("implement me")
}

func (x *VulMemoryDao) Update(ctx context.Context, vul *models.Vul, codes []*models.VulCode) error {
	//TODO implement me
	panic("implement me")
}

func (x *VulMemoryDao) Upsert(ctx context.Context, vul *models.Vul, codes []*models.VulCode) error {
	//TODO implement me
	panic("implement me")
}

func (x *VulMemoryDao) Delete(ctx context.Context, vulIds ...string) error {
	//TODO implement me
	panic("implement me")
}

func (x *VulMemoryDao) Find(ctx context.Context, vulId string) (*models.Vul, error) {
	//TODO implement me
	panic("implement me")
}

func (x *VulMemoryDao) FindByCve(ctx context.Context, cve string) (*models.Vul, error) {
	//TODO implement me
	panic("implement me")
}

func (x *VulMemoryDao) FindByCode(ctx context.Context, code string, codeType models.CodeType) (*models.Vul, error) {
	//TODO implement me
	panic("implement me")
}

func (x *VulMemoryDao) FindMany(ctx context.Context, vulIds ...string) ([]*models.Vul, error) {
	//TODO implement me
	panic("implement me")
}

func (x *VulMemoryDao) LoadAll(ctx context.Context) ([]*models.Vul, error) {
	//TODO implement me
	panic("implement me")
}

func (x *VulMemoryDao) CreateCodes(ctx context.Context, vulId string, codes []*models.VulCode) error {
	//TODO implement me
	panic("implement me")
}

func (x *VulMemoryDao) ReplaceCodes(ctx context.Context, vulId string, codes []*models.VulCode) error {
	//TODO implement me
	panic("implement me")
}

func (x *VulMemoryDao) UpsertCodes(ctx context.Context, vulId string, codes []*models.VulCode) error {
	//TODO implement me
	panic("implement me")
}

func (x *VulMemoryDao) FindCodes(ctx context.Context, vulIds ...string) ([]*models.VulCode, error) {
	//TODO implement me
	panic("implement me")
}

func (x *VulMemoryDao) DeleteCodeByVulId(ctx context.Context, vulIds ...string) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (x *VulMemoryDao) DeleteCode(ctx context.Context, code string) error {
	//TODO implement me
	panic("implement me")
}

func (x *VulMemoryDao) LoadAllCodes(ctx context.Context) ([]*models.VulCode, error) {
	//TODO implement me
	panic("implement me")
}

func (x *VulMemoryDao) ListCodeByType(ctx context.Context, codeType models.CodeType) ([]*models.VulCode, error) {
	//TODO implement me
	panic("implement me")
}
