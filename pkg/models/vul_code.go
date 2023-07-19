package models

import (
	"github.com/scagogogo/sca-base-module-vuls/pkg/naming"
	"time"
)

type CodeType int

const (

	// CodeTypeCVE CVE编号
	CodeTypeCVE CodeType = iota

	// CodeTypeGHSA GHSA编号
	CodeTypeGHSA

	// CodeTypeCNVD CNV编号
	CodeTypeCNVD

	// CodeTypeCNNVD CNNVD编号
	CodeTypeCNNVD
)

func (x CodeType) String() string {
	switch x {
	case CodeTypeCVE:
		return "CVE"
	case CodeTypeGHSA:
		return "GHSA"
	case CodeTypeCNVD:
		return "CNVD"
	case CodeTypeCNNVD:
		return "CNNVD"
	default:
		// 没有明确指定的都认为是未知了
		return "Unknown"
	}
}

// VulCode 用于表示漏洞关联到的编号
type VulCode struct {

	// 关联到的漏洞的ID，这个vul_id和code两个字段做一个联合的唯一索引
	VulId string `mapstructure:"vul_id" json:"vul_id,omitempty" gorm:"column:vul_id;index:vul_codes,unique"`

	// 漏洞的编号
	Code string `mapstructure:"code" json:"code,omitempty" gorm:"column:code;index:vul_codes,unique"`

	// 漏洞的编号类型
	CodeType CodeType `mapstructure:"code_type" json:"code_type,omitempty" gorm:"column:code_type"`

	// 几个创建时间更新时间啥的
	CreateTime *time.Time `mapstructure:"create_time" json:"create_time,omitempty" gorm:"column:create_time"`
	UpdateTime *time.Time `mapstructure:"update_time" json:"update_time,omitempty" gorm:"column:update_time"`
	ChangeTime *time.Time `mapstructure:"change_time" json:"change_time,omitempty" gorm:"column:change_time"`
}

func (x *VulCode) TableName() string {
	return naming.TableName("vul_codes")
}
