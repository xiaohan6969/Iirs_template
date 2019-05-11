package mysql

import (
	"config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//var (
//	Mysql *sql.DB
//)
//
//func init() {
//	user := config.Config.Get("mysql.user").(string)
//	password := config.Config.Get("mysql.password").(string)
//	database := config.Config.Get("mysql.database").(string)
//	host := config.Config.Get("mysql.host").(string)
//	port := config.Config.Get("mysql.port").(string)
//	//db, err := sql.Open("mysql", "数据库账号:数据库密码@tcp(数据库Ip:端口)/数据库名称?parseTime=true")
//	db, err := sql.Open("mysql", user+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset=utf8&parseTime=true&loc=Local")
//	if err != nil {
//		fmt.Println(err)
//	}
//	Mysql = db
//}

var (
	db  *gorm.DB
	ERR error
)

func init() {
	user := config.Config.Get("mysql.user").(string)
	pwd := config.Config.Get("mysql.password").(string)
	databases := config.Config.Get("mysql.database").(string)
	addr := config.Config.Get("mysql.host").(string)
	db, ERR = gorm.Open("mysql", user+":"+pwd+"@tcp("+addr+")/"+databases+"?charset=utf8&parseTime=true&loc=Local")
	if ERR != nil {
		panic("--- 数据库连接失败")
	}
	db.DB().SetMaxIdleConns(500)
	db.DB().SetMaxOpenConns(5000)
	db.DB().SetConnMaxLifetime(0)
	db.LogMode(true)
}

func GetDataBase() *gorm.DB {
	return db
}
