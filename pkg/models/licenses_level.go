package models

import (
	"github.com/scagogogo/sca-base-module-vuls/pkg/naming"
	"time"
)

// TODO 数据结构待调整

type License struct {
	Identifier string `mapstructure:"identifier" json:"identifier,omitempty" yaml:"identifier" db:"identifier" gorm:"column:identifier;primaryKey"`
	LevelId    int    `mapstructure:"level_id" json:"level_id,omitempty" yaml:"level_id" db:"level_id" gorm:"column:level_id"`

	// 默认值: 允许商业集成
	LevelDesc string    `mapstructure:"level_desc" json:"level_desc,omitempty" yaml:"level_desc" db:"level_desc" gorm:"column:level_desc"`
	CreateAt  time.Time `mapstructure:"create_at" json:"create_at,omitempty" yaml:"create_at" db:"create_at" gorm:"column:create_at"`
	UpdateAt  time.Time `mapstructure:"update_at" json:"update_at,omitempty" yaml:"update_at" db:"update_at" gorm:"column:update_at"`
}

func (x *License) TableName() string {
	return naming.TableName("licenses")
}

func NewLicensesLevel() *License {
	return &License{
		LevelId:   0,
		LevelDesc: "允许商业集成",
	}
}
