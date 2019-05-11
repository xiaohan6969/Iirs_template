package sql

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"model"
	"server/databases/mysql"
)

type Query struct{}

func (a *Query) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/query", "Query")
	b.Handle("GET", "/test", "Test")
}

func (a *Query) Query(ctx iris.Context) iris.Map {
	db := mysql.Mysql
	//关闭数据库，db会被多个goroutine共享，可以不调用
	//defer db.Close()
	result := model.Test{}
	//查询数据，指定字段名，返回sql.Rows结果集
	sql := "select * from test"
	res, err := db.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	var res1 []model.Test
	for res.Next() {
		err = res.Scan(&result.Id, &result.Name, &result.Sex, &result.Age)
		res1 = append(res1, result)
	}

	//for result.Next() {
	//
	//	if err := result.Scan(&v,&b,&c,&d); err != nil {
	//		panic(err)
	//	}
	//	fmt.Println("sadsa==",a)
	//}
	return iris.Map{
		"result": res1,
	}
}

func (a *Query) Test(ctx iris.Context) {

}

func (a *Query) Insert(ctx iris.Context) {
	//create table test(id int primary key auto_increment, Name varchar(18),  sex varchar(2),  age int);
	//insert into test(Name,sex,age) values("2","3",12);
}
