package mysqlServer

import (
	"../../config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	JzGorm *gorm.DB
)

func init() {
	user := config.Config.Get("mysql.user").(string)
	password := config.Config.Get("mysql.password").(string)
	database := config.Config.Get("mysql.database").(string)
	host := config.Config.Get("mysql.host").(string)
	port := config.Config.Get("mysql.port").(string)
	db, err := gorm.Open("mysql", user+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	JzGorm = db
}
