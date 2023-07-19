package db

import (
	"github.com/scagogogo/sca-base-module-dao/mysql"
	"github.com/scagogogo/sca-base-module-vuls/pkg/models"
)

func init() {

	if mysql.Gorm == nil {
		return
	}

	err := mysql.Gorm.AutoMigrate(&models.ComponentVul{})
	if err != nil {
		panic(err)
	}

	err = mysql.Gorm.AutoMigrate(&models.Vul{})
	if err != nil {
		panic(err)
	}
	err = mysql.Gorm.AutoMigrate(&models.VulCode{})
	if err != nil {
		panic(err)
	}

}
