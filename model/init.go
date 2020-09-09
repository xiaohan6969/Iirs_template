package model

import (
	"../server/mysqlServer"
	"./commonStruct"
	"fmt"
)

func MysqlTableInit() {}

func init() {
	db := mysqlServer.JzGorm
	if !db.HasTable(&commonStruct.User{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&commonStruct.User{}).Error; err != nil {
			fmt.Println("unusual===", err)
		}
	}
	if !db.HasTable(&commonStruct.HomePage{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&commonStruct.HomePage{}).Error; err != nil {
			fmt.Println("unusual===", err)
		}
	}
}
