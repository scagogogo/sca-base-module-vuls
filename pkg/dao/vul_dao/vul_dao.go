package vul_dao

import (
	"context"
	"github.com/scagogogo/sca-base-module-vuls/pkg/models"
)

// VulDao 漏洞的dao
type VulDao interface {

	// Create 创建漏洞，需要保证编号与漏洞的原子性
	Create(ctx context.Context, vul *models.Vul, codes []*models.VulCode) error

	// Update 更新漏洞，需要保证编号与漏洞的原子性
	Update(ctx context.Context, vul *models.Vul, codes []*models.VulCode) error

	// Upsert 漏洞存在则更新，不存在则删除，需要保证编号与漏洞的原子性
	Upsert(ctx context.Context, vul *models.Vul, codes []*models.VulCode) error

	// Delete 删除漏洞，支持一次删除多个漏洞
	Delete(ctx context.Context, vulIds ...string) error

	// Find 根据漏洞ID查询详情
	Find(ctx context.Context, vulId string) (*models.Vul, error)

	// FindByCve 根据CVE查询漏洞
	FindByCve(ctx context.Context, cve string) (*models.Vul, error)

	// FindByCode 根据漏洞代码查询漏洞
	FindByCode(ctx context.Context, code string, codeType models.CodeType) (*models.Vul, error)

	// FindMany 一次查询多个漏洞信息
	FindMany(ctx context.Context, vulIds ...string) ([]*models.Vul, error)

	// LoadAll 加载所有的漏洞
	LoadAll(ctx context.Context) ([]*models.Vul, error)

	// CreateCodes 为漏洞创建编号，如果编号已经存在则创建失败
	CreateCodes(ctx context.Context, vulId string, codes []*models.VulCode) error

	// ReplaceCodes 为漏洞更新编号，如果编号已经存在则更新，并且会将多余的编号删除，相当于是覆盖更新
	ReplaceCodes(ctx context.Context, vulId string, codes []*models.VulCode) error

	// UpsertCodes 更新漏洞的编号，如果编号已经存在则更新，否则插入
	UpsertCodes(ctx context.Context, vulId string, codes []*models.VulCode) error

	// FindCodes 查询给定的漏洞的编号，支持一次传入多个漏洞ID
	FindCodes(ctx context.Context, vulIds ...string) ([]*models.VulCode, error)

	// DeleteCodeByVulId 根据给定的漏洞ID删除漏洞编号，支持一次传入多个漏洞ID
	DeleteCodeByVulId(ctx context.Context, vulIds ...string) (int64, error)

	// DeleteCode 删除给定的漏洞编号
	DeleteCode(ctx context.Context, code string) error

	// LoadAllCodes 加载所有的漏洞编号
	LoadAllCodes(ctx context.Context) ([]*models.VulCode, error)
}
