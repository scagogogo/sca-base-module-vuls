package component_vul_dao

import (
	"github.com/scagogogo/sca-base-module-dao/mysql"
	"testing"

	// 初始化创建表之类的
	_ "github.com/scagogogo/sca-base-module-vuls/pkg/dao"
)

func TestComponentVulMysqlDao(t *testing.T) {
	dao := NewComponentVulMysqlDao(mysql.Gorm)
	ComponentVulDaoTest(t, dao)
}
