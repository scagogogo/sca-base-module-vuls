package maven_component_vul_dao

import (
	"github.com/scagogogo/sca-base-module-dao/pkg/mysql"
	"github.com/scagogogo/sca-base-module-vuls/pkg/models"
)

// Maven的结构可以暂时先这样不用动，暂时不需要加入到他统一架构中

// UpsertMavenComponentVul 有组件漏洞则更新，否则插入
func UpsertMavenComponentVul(v *models.MavenComponentVul) error {
	vul, err := FindMavenComponentVul(v.GroupId, v.ArtifactId, v.Version, v.VulId)
	if err != nil {
		return err
	}
	if vul == nil {
		// 之前不存在
		return mysql.Gorm.Model(&models.MavenComponentVul{}).Create(v).Error
	} else {
		// 之前就已经存在，则保留之前的创建时间，其它字段更新
		v.CreateTime = vul.CreateTime
		return mysql.Gorm.Model(&models.MavenComponentVul{}).
			Where("group_id = ? AND artifact_id = ? AND version = ? AND vul_id = ?", v.GroupId, v.ArtifactId, v.Version, v.VulId).
			Save(v).
			Error
	}
}

// FindMavenComponentVul 根据GAV和漏洞查询之前存储的漏洞信息
func FindMavenComponentVul(groupId, artifactId, version, vulId string) (*models.MavenComponentVul, error) {
	var r *models.MavenComponentVul
	err := mysql.Gorm.Model(&models.MavenComponentVul{}).Where("group_id = ? AND artifact_id = ? AND version = ? AND vul_id = ?", groupId, artifactId, version, vulId).Scan(&r).Error
	if err != nil {
		return nil, err
	}
	return r, nil
}
