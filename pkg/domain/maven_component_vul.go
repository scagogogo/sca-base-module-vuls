package domain

import (
	"github.com/scagogogo/sca-base-module-vuls/pkg/naming"
	"gorm.io/gorm/schema"
	"time"
)

// MavenComponentVul TODO 准备废弃
type MavenComponentVul struct {
	GroupId    string     `mapstructure:"group_id" json:"group_id,omitempty" yaml:"group_id" db:"group_id" gorm:"column:group_id"`
	ArtifactId string     `mapstructure:"artifact_id" json:"artifact_id,omitempty" yaml:"artifact_id" db:"artifact_id" gorm:"column:artifact_id"`
	Version    string     `mapstructure:"version" json:"version,omitempty" yaml:"version" db:"version" gorm:"column:version"`
	VulId      string     `mapstructure:"vul_id" json:"vul_id,omitempty" yaml:"vul_id" db:"vul_id" gorm:"column:vul_id"`
	CreateTime *time.Time `mapstructure:"create_time" json:"create_time,omitempty" yaml:"create_time" db:"create_time" gorm:"create_time"`
	UpdateTime *time.Time `mapstructure:"update_time" json:"update_time,omitempty" yaml:"update_time" db:"update_time" gorm:"update_time"`
}

var _ schema.Tabler = &MavenComponentVul{}

func (x *MavenComponentVul) TableName() string {
	return naming.TableName("maven_component_vuls")
}
