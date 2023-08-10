package models

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	osv_schema "github.com/scagogogo/osv-schema"
	"github.com/scagogogo/sca-base-module-vuls/pkg/naming"
	"golang.org/x/text/language"
	"time"
)

// Vul 表示一条漏洞信息
type Vul struct {

	// 漏洞ID
	VulId string `mapstructure:"vul_id" json:"vul_id,omitempty" yaml:"vul_id" db:"vul_id" gorm:"column:vul_id;primaryKey"`

	// 漏洞的CVSS3评分
	CVSS3 string `mapstructure:"cvss_v3" json:"cvss_v3,omitempty" yaml:"cvss_v3" db:"cvss_v3" gorm:"column:cvss_v3"`

	// 漏洞的cwe，json的string array
	CWE osv_schema.Aliases `mapstructure:"cwe" json:"cwe,omitempty" yaml:"cwe" db:"cwe" gorm:"column:cwe;serializer:json"`

	// 漏洞的标题
	Title I18nText `mapstructure:"title" json:"title,omitempty" yaml:"title" db:"title" gorm:"column:title;serializer:json"`

	// 漏洞的描述
	Description I18nText `mapstructure:"description" json:"description,omitempty" yaml:"description" db:"description" gorm:"column:description;serializer:json"`

	// 漏洞相关的引用文章
	References References `mapstructure:"references" json:"references,omitempty" yaml:"references" db:"references" gorm:"column:references;serializer:json"`

	Severity string `mapstructure:"severity" json:"severity,omitempty" yaml:"severity" db:"severity" gorm:"column:severity"`

	// 漏洞的发布时间
	PublishedTime *time.Time `mapstructure:"published_time" json:"published_time,omitempty" yaml:"published_time" db:"published_time" gorm:"column:published_time"`

	//Codes []*VulCode `mapstructure:"codes" json:"codes,omitempty" yaml:"codes" db:"codes" gorm:"column:codes"`

	// 几个时间
	CreateTime *time.Time `mapstructure:"create_time" json:"create_time,omitempty" yaml:"create_time" db:"create_time" gorm:"column:create_time"`
	UpdateTime *time.Time `mapstructure:"update_time" json:"update_time,omitempty" yaml:"update_time" db:"update_time" gorm:"column:update_time"`
	ChangeTime *time.Time `mapstructure:"change_time" json:"change_time,omitempty" yaml:"change_time" db:"change_time" gorm:"column:change_time"`
}

// MergeTitle 把给定的语言合并到title中，如果已经存在的话则忽略
func (x *Vul) MergeTitle(language language.Tag, title string) *Vul {
	x.Title = x.Title.Merge(language, title)
	return x
}

func (x *Vul) MergeTitleAll(title I18nText) *Vul {
	x.Title = x.Title.MergeAll(title)
	return x
}

// UpdateTitle 把给定的语言合并到title中，如果已经存在的话则覆盖
func (x *Vul) UpdateTitle(language language.Tag, title string) *Vul {
	x.Title = x.Title.Set(language, title)
	return x
}

func (x *Vul) UpdateTitleAll(title I18nText) *Vul {
	x.Title = x.Title.SetAll(title)
	return x
}

// MergeDescription 把给定的语言合并到描述中，如果已经存在的话则忽略
func (x *Vul) MergeDescription(language language.Tag, description string) *Vul {
	x.Description = x.Description.Merge(language, description)
	return x
}

func (x *Vul) MergeDescriptionAll(description I18nText) *Vul {
	x.Description = x.Description.MergeAll(description)
	return x
}

// UpdateDescription 把给定的语言合并到描述中，如果已经存在的话则覆盖
func (x *Vul) UpdateDescription(language language.Tag, description string) *Vul {
	x.Description = x.Description.Set(language, description)
	return x
}

func (x *Vul) UpdateDescriptionAll(description I18nText) *Vul {
	x.Description = x.Description.SetAll(description)
	return x
}

// MergeReferences 把给定的语言合并到引用中，如果已经存在的话则忽略
func (x *Vul) MergeReferences(language language.Tag, reference *Reference) *Vul {

	// 先去重，如果已经存在的话则忽略
	for _, reference := range x.References {
		if reference.URL == reference.URL {
			return x
		}
	}

	x.References = append(x.References, reference)
	return x
}

// MergeReferencesAll 把给定的语言合并到引用中，如果已经存在的话则忽略
func (x *Vul) MergeReferencesAll(references References) *Vul {

	// 把已经存在的做个set先
	existsReferenceUrlSet := make(map[string]struct{}, 0)
	for _, reference := range x.References {
		existsReferenceUrlSet[reference.URL] = struct{}{}
	}

	// 然后把新增的加进来
	for _, reference := range references {
		if _, exists := existsReferenceUrlSet[reference.URL]; exists {
			continue
		}
		x.References = append(x.References, reference)
	}

	return x
}

// UpdateReferences 把给定的语言合并到引用中，如果已经存在的话则覆盖
func (x *Vul) UpdateReferences(language language.Tag, reference *Reference) *Vul {

	// 如果已经存在的话则更新其值
	for index, reference := range x.References {
		if reference.URL == reference.URL {
			x.References[index] = reference
			return x
		}
	}

	// 如果不存在的话则新增
	x.References = append(x.References, reference)
	return x
}

// UpdateReferencesAll 更新引用，如果引用已经存在了则覆盖更新，否则新增，在修改引用列表的时候会尽可能的保持其原有顺序
func (x *Vul) UpdateReferencesAll(references References) *Vul {

	// 把新增的做个map先
	needUpdateReferenceUrlMap := make(map[string]*Reference, 0)
	for _, reference := range references {
		needUpdateReferenceUrlMap[reference.URL] = reference
	}

	// 然后把需要更新的先更新了，更新的同时要删除被替换掉的reference，防止后面被重复处理
	for index, reference := range x.References {
		if newReference, exists := needUpdateReferenceUrlMap[reference.URL]; exists {
			x.References[index] = newReference
			delete(needUpdateReferenceUrlMap, reference.URL)
		}
	}

	// 然后再把新增的加进来，保持其原有顺序
	for _, reference := range references {
		// 只处理没被处理过的，因为可能会有一部分已经存在的url在上一步被处理掉了
		if _, exists := needUpdateReferenceUrlMap[reference.URL]; !exists {
			continue
		}
		x.References = append(x.References, reference)
	}

	return x
}

func (x *Vul) TableName() string {
	return naming.TableName("vuls")
}

type CWE []string

var _ sql.Scanner = &CWE{}
var _ driver.Valuer = &CWE{}

func (x *CWE) Scan(src any) error {
	if src == nil {
		return nil
	}
	bytes, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("scan error")
	}
	if len(bytes) == 0 {
		return nil
	}
	return json.Unmarshal(bytes, &x)
}

func (x CWE) Value() (driver.Value, error) {
	if len(x) == 0 {
		return nil, nil
	}
	marshal, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	if len(marshal) == 0 {
		return nil, nil
	}
	return string(marshal), nil
}
