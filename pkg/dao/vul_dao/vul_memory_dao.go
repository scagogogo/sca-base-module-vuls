package vul_dao

import (
	"context"
	"github.com/scagogogo/sca-base-module-vuls/pkg/domain"
)

type VulMemoryDao struct {

	// 漏洞
	db map[string]*domain.Vul

	// map[漏洞ID]map[编号ID]domain.VulCode
	codes map[string]map[string]*domain.VulCode
}

var _ VulDao = &VulMemoryDao{}

func NewVulMemoryDao() *VulMemoryDao {
	return &VulMemoryDao{}
}

func (x *VulMemoryDao) Create(ctx context.Context, vul *domain.Vul, codes []*domain.VulCode) error {
	//TODO implement me
	panic("implement me")
}

func (x *VulMemoryDao) Update(ctx context.Context, vul *domain.Vul, codes []*domain.VulCode) error {
	//TODO implement me
	panic("implement me")
}

func (x *VulMemoryDao) Upsert(ctx context.Context, vul *domain.Vul, codes []*domain.VulCode) error {
	//TODO implement me
	panic("implement me")
}

func (x *VulMemoryDao) Delete(ctx context.Context, vulIds ...string) error {
	//TODO implement me
	panic("implement me")
}

func (x *VulMemoryDao) Find(ctx context.Context, vulId string) (*domain.Vul, error) {
	//TODO implement me
	panic("implement me")
}

func (x *VulMemoryDao) FindMany(ctx context.Context, vulIds ...string) ([]*domain.Vul, error) {
	//TODO implement me
	panic("implement me")
}

func (x *VulMemoryDao) LoadAll(ctx context.Context) ([]*domain.Vul, error) {
	//TODO implement me
	panic("implement me")
}

func (x *VulMemoryDao) CreateCodes(ctx context.Context, vulId string, codes []*domain.VulCode) error {
	//TODO implement me
	panic("implement me")
}

func (x *VulMemoryDao) ReplaceCodes(ctx context.Context, vulId string, codes []*domain.VulCode) error {
	//TODO implement me
	panic("implement me")
}

func (x *VulMemoryDao) UpsertCodes(ctx context.Context, vulId string, codes []*domain.VulCode) error {
	//TODO implement me
	panic("implement me")
}

func (x *VulMemoryDao) FindCodes(ctx context.Context, vulIds ...string) ([]*domain.VulCode, error) {
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

func (x *VulMemoryDao) LoadAllCodes(ctx context.Context) ([]*domain.VulCode, error) {
	//TODO implement me
	panic("implement me")
}
