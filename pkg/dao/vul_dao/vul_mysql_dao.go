package vul_dao

import (
	"context"
	"github.com/scagogogo/sca-base-module-vuls/pkg/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type VulMysqlDao struct {
	gorm *gorm.DB
}

var _ VulDao = &VulMysqlDao{}

func NewVulMysqlDao(gorm *gorm.DB) *VulMysqlDao {
	return &VulMysqlDao{
		gorm: gorm,
	}
}

func (x *VulMysqlDao) Create(ctx context.Context, vul *models.Vul, codes []*models.VulCode) error {
	// 在一个事务中更新，尽可能保证外键的一致性
	return x.gorm.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		// 首先保存漏洞信息
		err := tx.Model(&vul).Create(vul).Error
		if err != nil {
			return err
		}

		// 然后再更新漏洞编号，这个漏洞编号必须是全部插入成功的状态，否则此条漏洞保存失败
		for _, code := range codes {
			err = tx.Model(&code).Create(code).Error
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (x *VulMysqlDao) Update(ctx context.Context, vul *models.Vul, codes []*models.VulCode) error {
	// 在一个事务中更新，尽可能保证外键的一致性
	return x.gorm.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		// 保存漏洞信息
		vul.CreateTime = nil
		err := tx.Model(&vul).Where("vul_id = ?", vul.VulId).Updates(vul).Error
		if err != nil {
			return err
		}

		// 然后再更新漏洞编号，漏洞编号可能已经存在，也可能不存在，所以需要upsert
		for _, code := range codes {
			err := tx.Model(&code).Where("code = ?", code.Code).Updates(&code).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (x *VulMysqlDao) Upsert(ctx context.Context, vul *models.Vul, codes []*models.VulCode) error {
	// 在一个事务中更新，尽可能保证外键的一致性
	return x.gorm.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		// 保存漏洞信息，存在则更新，不存在则插入
		err := tx.Model(&vul).Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "vul_id"}},
			DoUpdates: clause.Assignments(map[string]any{
				"cvss_v3":        vul.CVSS3,
				"cwe":            vul.CWE,
				"title":          vul.Title,
				"description":    vul.Description,
				"references":     vul.References,
				"severity":       vul.Severity,
				"published_time": vul.PublishedTime,
				"update_time":    vul.UpdateTime,
			}),
		}).Create(&vul).Error
		if err != nil {
			return err
		}

		// 然后再更新漏洞编号，漏洞编号可能已经存在，也可能不存在，所以需要upsert
		for _, code := range codes {
			err := tx.Model(&code).Clauses(clause.OnConflict{
				Columns: []clause.Column{{Name: "code"}},
				DoUpdates: clause.Assignments(map[string]any{
					"vul_id":      code.VulId,
					"code_type":   code.CodeType,
					"update_time": code.UpdateTime,
					"change_time": code.ChangeTime,
				}),
			}).Create(&code).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (x *VulMysqlDao) Delete(ctx context.Context, vulIds ...string) error {
	return x.gorm.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		// 删除漏洞信息
		db := tx.Where("vul_id in ?", vulIds).Delete(&models.Vul{})
		if db.Error != nil {
			return db.Error
		}

		// 删除漏洞的编号信息
		return tx.Where("vul_id in ?", vulIds).Delete(&models.VulCode{}).Error
	})
}

func (x *VulMysqlDao) Find(ctx context.Context, vulId string) (*models.Vul, error) {
	var r *models.Vul
	err := x.gorm.WithContext(ctx).Model(&r).Where("vul_id = ?", vulId).Scan(&r).Error
	if err != nil {
		return nil, err
	} else {
		return r, nil
	}
}

func (x *VulMysqlDao) FindMany(ctx context.Context, vulIds ...string) ([]*models.Vul, error) {
	var slice []*models.Vul
	err := x.gorm.WithContext(ctx).Model(&models.Vul{}).Where("vul_id in ?", vulIds).Scan(&slice).Error
	if err != nil {
		return nil, err
	} else {
		return slice, nil
	}
}

func (x *VulMysqlDao) LoadAll(ctx context.Context) ([]*models.Vul, error) {
	var slice []*models.Vul
	err := x.gorm.WithContext(ctx).Model(&models.Vul{}).Scan(&slice).Error
	if err != nil {
		return nil, err
	} else {
		return slice, nil
	}
}

func (x *VulMysqlDao) CreateCodes(ctx context.Context, vulId string, codes []*models.VulCode) error {
	return x.gorm.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, code := range codes {
			err := tx.Model(&code).Create(&code).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (x *VulMysqlDao) ReplaceCodes(ctx context.Context, vulId string, codes []*models.VulCode) error {
	return x.gorm.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		// 计算出需要删除的编号
		var existsCodes []*models.VulCode
		err := tx.Model(&models.VulCode{}).Where("vul_id = ?", vulId).Scan(&existsCodes).Error
		if err != nil {
			return err
		}
		needUpsertCodeSet := make(map[string]struct{}, 0)
		for _, code := range codes {
			needUpsertCodeSet[code.Code] = struct{}{}
		}
		needDeleteCodes := make([]string, 0)
		for _, code := range existsCodes {
			_, exists := needUpsertCodeSet[code.Code]
			if !exists {
				needDeleteCodes = append(needDeleteCodes, code.Code)
			}
		}
		// 删除多余的编号
		err = tx.Where("code in ?", needDeleteCodes).Delete(&models.VulCode{}).Error
		if err != nil {
			return err
		}

		// 然后对剩下的编号进行upsert
		for _, code := range codes {
			err := x.gorm.Model(&code).Clauses(clause.OnConflict{
				Columns: []clause.Column{{Name: "code"}},
				DoUpdates: clause.Assignments(map[string]any{
					"vul_id":      code.VulId,
					"code_type":   code.CodeType,
					"update_time": code.UpdateTime,
					"change_time": code.ChangeTime,
				}),
			}).Create(&code).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (x *VulMysqlDao) UpsertCodes(ctx context.Context, vulId string, codes []*models.VulCode) error {
	return x.gorm.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, code := range codes {
			err := x.gorm.Model(&code).Clauses(clause.OnConflict{
				Columns: []clause.Column{{Name: "code"}},
				DoUpdates: clause.Assignments(map[string]any{
					"vul_id":      code.VulId,
					"code_type":   code.CodeType,
					"update_time": code.UpdateTime,
					"change_time": code.ChangeTime,
				}),
			}).Create(&code).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (x *VulMysqlDao) FindCodes(ctx context.Context, vulIds ...string) ([]*models.VulCode, error) {
	var slice []*models.VulCode
	err := x.gorm.WithContext(ctx).Model(&models.VulCode{}).Where("vul_id in ?", vulIds).Scan(&slice).Error
	if err != nil {
		return nil, err
	} else {
		return slice, nil
	}
}

func (x *VulMysqlDao) DeleteCodeByVulId(ctx context.Context, vulIds ...string) (int64, error) {
	tx := x.gorm.WithContext(ctx).Where("vul_id in ?", vulIds).Delete(&models.VulCode{})
	return tx.RowsAffected, tx.Error
}

func (x *VulMysqlDao) DeleteCode(ctx context.Context, code string) error {
	return x.gorm.WithContext(ctx).Where("code = ?", code).Delete(&models.VulCode{}).Error
}

func (x *VulMysqlDao) LoadAllCodes(ctx context.Context) ([]*models.VulCode, error) {
	var slice []*models.VulCode
	err := x.gorm.WithContext(ctx).Model(&models.VulCode{}).Scan(&slice).Error
	if err != nil {
		return nil, err
	} else {
		return slice, nil
	}
}
