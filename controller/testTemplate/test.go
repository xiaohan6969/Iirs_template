package testTemplate

import (
	"../../config"
	config2 "../../server/mysqlServer"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
	"strconv"
)

type SqlNature struct{}

var (
	table1 = config.Config.Get("mysql.table1").(string)
)

func (a *SqlNature) BeforeActivation(h mvc.BeforeActivation) {
	h.Handle("GET", "/test/sql", "Test")
	h.Handle("GET", "/index/list", "IndexList")
	h.Handle("POST", "/choice/one/detail", "OneDetail")
}

func (a *SqlNature) IndexList(ctx iris.Context) iris.Map {
	page := ctx.URLParam("page")
	size := ctx.URLParam("size")
	fmt.Println(page)
	page1, err := strconv.Atoi(page)
	size1, err := strconv.Atoi(size)

	db := config2.Mysql
	//查询数据，指定字段名，返回sql.Rows结果集
	sql := "select id,first_name,last_name from " + table1 + " limit " + strconv.Itoa((page1-1)*size1) + "," + size
	type DetailedQuery struct {
		Id        int    `sql:"id"`
		FirstName string `sql:"first_name"`
		LastName  string `sql:"last_name"`
	}
	fmt.Println(sql)
	querySet, err := db.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	res := DetailedQuery{}
	var result []interface{}
	for querySet.Next() {
		err = querySet.Scan(
			&res.Id,        //字段1
			&res.FirstName, //字段1
			&res.LastName,  //字段2
		)
		result = append(result, DetailedQuery{
			res.Id,
			res.FirstName,
			res.LastName,
		})
	}
	defer func() {
		if err := querySet.Close(); err != nil {
			fmt.Println("close fail")
		}
	}()
	return iris.Map{
		"status":  200,
		"data":    result,
		"message": "1111",
	}
}

func (a *SqlNature) OneDetail(ctx iris.Context) iris.Map {
	db := config2.Mysql
	type request struct {
		IndexId int `json:"index_id"`
	}
	values := request{}
	err := ctx.ReadJSON(&values)
	if err != nil {
		log.Println(err)
	}
	//查询数据，指定字段名，返回sql.Rows结果集
	sql := "select id,first_name,last_name from " + table1 + " where id = " + strconv.Itoa(values.IndexId)
	type DetailedQuery struct {
		Id        int    `sql:"id"`
		FirstName string `sql:"first_name"`
		LastName  string `sql:"last_name"`
	}
	querySet, err := db.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	res := DetailedQuery{}
	var result []interface{}
	for querySet.Next() {
		err = querySet.Scan(
			&res.Id,        //字段1
			&res.FirstName, //字段1
			&res.LastName,  //字段2
		)
		result = append(result, DetailedQuery{
			res.Id,
			res.FirstName,
			res.LastName,
		})
	}
	defer func() {
		if err := querySet.Close(); err != nil {
			fmt.Println("close fail")
		}
	}()
	return iris.Map{
		"status":  200,
		"data":    res,
		"message": "1111",
	}
}

func (a *SqlNature) Test() {}
