package models

import (
	"github.com/scagogogo/sca-base-module-vuls/pkg/naming"
	"gorm.io/gorm/schema"
	"time"
)

// MavenComponentVul TODO 准备废弃
type MavenComponentVul struct {

	// Maven组件的唯一标识
	GroupId    string `mapstructure:"group_id" json:"group_id,omitempty" yaml:"group_id" db:"group_id" gorm:"column:group_id;index:maven_component_vuls,unique"`
	ArtifactId string `mapstructure:"artifact_id" json:"artifact_id,omitempty" yaml:"artifact_id" db:"artifact_id" gorm:"column:artifact_id;index:maven_component_vuls,unique"`
	Version    string `mapstructure:"version" json:"version,omitempty" yaml:"version" db:"version" gorm:"column:version;index:maven_component_vuls,unique"`

	// 关联到的漏洞的ID
	VulId string `mapstructure:"vul_id" json:"vul_id,omitempty" yaml:"vul_id" db:"vul_id" gorm:"column:vul_id;index:maven_component_vuls,unique"`

	// 创建时间
	CreateTime *time.Time `mapstructure:"create_time" json:"create_time,omitempty" yaml:"create_time" db:"create_time" gorm:"create_time"`
	UpdateTime *time.Time `mapstructure:"update_time" json:"update_time,omitempty" yaml:"update_time" db:"update_time" gorm:"update_time"`
}

var _ schema.Tabler = &MavenComponentVul{}

func (x *MavenComponentVul) TableName() string {
	return naming.TableName("maven_component_vuls")
}
