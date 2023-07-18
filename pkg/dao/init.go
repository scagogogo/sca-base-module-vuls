package db

import (
	"github.com/scagogogo/sca-base-module-dao/mysql"
	"github.com/scagogogo/sca-base-module-vuls/pkg/domain"
)

func init() {

	if mysql.Gorm == nil {
		return
	}

	err := mysql.Gorm.AutoMigrate(&domain.ComponentVul{})
	if err != nil {
		panic(err)
	}

	err = mysql.Gorm.AutoMigrate(&domain.Vul{})
	if err != nil {
		panic(err)
	}
	err = mysql.Gorm.AutoMigrate(&domain.VulCode{})
	if err != nil {
		panic(err)
	}

}
