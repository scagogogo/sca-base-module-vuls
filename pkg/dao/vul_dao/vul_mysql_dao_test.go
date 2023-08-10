package vul_dao

import (
	"github.com/scagogogo/sca-base-module-dao/pkg/mysql"
	"testing"

	// 初始化创建表之类的
	_ "github.com/scagogogo/sca-base-module-vuls/pkg/dao"
)

func TestVulMysqlDao(t *testing.T) {
	VulDaoTest(t, NewVulMysqlDao(mysql.Gorm))
}
