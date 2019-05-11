package sql

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"server/databases/mysql"
)

type Query struct{}

func (a *Query) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/test", "Test")
}

func (a *Query) Test(ctx iris.Context) {
	db := mysql.Mysql
	//关闭数据库，db会被多个goroutine共享，可以不调用
	defer db.Close()
	//查询数据，指定字段名，返回sql.Rows结果集
	rows2, _ := db.Query("select * from 宁波港口指数")
	fmt.Println("===", rows2)
}
func (a *Query) Insert(ctx iris.Context) {
	//create table test(id int primary key auto_increment, Name varchar(18),  sex varchar(2),  age int);
	//insert into test(Name,sex,age) values("2","3",12);
}

func (a *Query) Find(ctx iris.Context) {
	//db := mysql.Mysql
}

func (a *Query) Query(ctx iris.Context) {
}
