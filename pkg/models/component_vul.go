package models

import (
	"fmt"
	"github.com/scagogogo/sca-base-module-vuls/pkg/naming"
	"gorm.io/gorm/schema"
	"time"
)

// ComponentVul 组件漏洞表
type ComponentVul struct {

	// 组件的名字
	Name string `mapstructure:"name" json:"name,omitempty" yaml:"name" db:"name" gorm:"column:name;index:component_vuls,unique"`

	// 组件的版本
	Version string `mapstructure:"version" json:"version,omitempty" yaml:"version" db:"version" gorm:"column:version;index:component_vuls,unique"`

	// 漏洞的ID
	VulId string `mapstructure:"vul_id" json:"vul_id,omitempty" yaml:"vul_id" db:"vul_id" gorm:"column:vul_id;index:component_vuls,unique"`

	// 创建时间
	CreateTime *time.Time `mapstructure:"create_time" json:"create_time,omitempty" yaml:"create_time" db:"create_time" gorm:"create_time"`
	UpdateTime *time.Time `mapstructure:"update_time" json:"update_time,omitempty" yaml:"update_time" db:"update_time" gorm:"update_time"`
	ChangeTime *time.Time `mapstructure:"change_time" json:"change_time,omitempty" yaml:"change_time" db:"change_time" gorm:"change_time"`
}

var _ schema.Tabler = &ComponentVul{}

func (x *ComponentVul) TableName() string {
	return naming.TableName("component_vuls")
}

// BuildComponentId 组件的唯一标识符
func (x *ComponentVul) BuildComponentId() string {
	return fmt.Sprintf("%s:%s", x.Name, x.Version)
}
