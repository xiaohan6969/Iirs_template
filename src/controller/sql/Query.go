package sql

import (
	"config"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"model"
	"server/databases/mysql"
)

type Query struct{}

var (
	database = config.Config.Get("mysql.database").(string)
)

func (a *Query) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/query", "Query")
	b.Handle("GET", "/test", "Test")
}

func (a *Query) Query(ctx iris.Context) interface{} {
	db := mysql.Mysql
	//关闭数据库，db会被多个goroutine共享，可以不调用
	//defer db.Close()
	//查询数据，指定字段名，返回sql.Rows结果集
	//sql := "select * from " + database
	sql := "select id,name,sex,age from " + database
	querySet, err := db.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	res := model.Test{}
	var result []model.Test
	for querySet.Next() {
		err = querySet.Scan(
			&res.Id,
			&res.Name,
			&res.Sex,
			&res.Age)
		result = append(result, res)
	}
	err = querySet.Close()
	return result
}

func (a *Query) Test(ctx iris.Context) {

}

func (a *Query) Insert(ctx iris.Context) {
	//create table test(id int primary key auto_increment, Name varchar(18),  sex varchar(2),  age int);
	//insert into test(Name,sex,age) values("2","3",12);
}
