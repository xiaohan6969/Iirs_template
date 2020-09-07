package model

import (
	"../server/mysqlServer"
	"./commonStruct"
	"fmt"
)

func Init() {}
func init() {
	db := mysqlServer.JzGorm
	if !db.HasTable(&commonStruct.User{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&commonStruct.User{}).Error; err != nil {
			fmt.Println("err===", err)
		}
	}
}
