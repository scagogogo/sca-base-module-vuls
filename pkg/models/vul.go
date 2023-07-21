package models

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	osv_schema "github.com/scagogogo/osv-schema"
	"github.com/scagogogo/sca-base-module-vuls/pkg/naming"
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
