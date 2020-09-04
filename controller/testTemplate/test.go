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

type DetailedQuery struct {
	Id         int    `json:"id" sql:"id"`
	Title      string `json:"title" sql:"title"`
	Content    string `json:"content" sql:"content"`
	CreateTime string `json:"create_time" sql:"create_time"`
}

func (a *SqlNature) IndexList(ctx iris.Context) iris.Map {
	page, err := ctx.URLParamInt("page")
	size, err := ctx.URLParamInt("size")
	fmt.Println(page)
	if page == -1 {
		page = 1
	}
	if size == -1 {
		size = 10
	}
	db := config2.Mysql
	//查询数据，指定字段名，返回sql.Rows结果集
	sql := "select id,title,content,create_time from " + table1 + " limit " + strconv.Itoa((page-1)*size) + "," + strconv.Itoa(size)
	fmt.Println(sql)
	querySet, err := db.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	res := DetailedQuery{}
	var result []interface{}
	for querySet.Next() {
		err = querySet.Scan(
			&res.Id,
			&res.Title,
			&res.Content,
			&res.CreateTime,
		)
		result = append(result, DetailedQuery{
			res.Id,
			res.Title,
			res.Content,
			res.CreateTime,
		})
	}
	defer func() {
		if err := querySet.Close(); err != nil {
			fmt.Println("close fail")
		}
	}()
	if len(result) == 0 {
		return iris.Map{
			"status":  200,
			"data":    []string{},
			"message": "获取成功",
		}
	}
	return iris.Map{
		"status":  200,
		"data":    result,
		"message": "获取成功",
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
	sql := "select id,title,content,create_time from " + table1 + " where id = " + strconv.Itoa(values.IndexId)
	querySet, err := db.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	res := DetailedQuery{}
	for querySet.Next() {
		err = querySet.Scan(
			&res.Id,
			&res.Title,
			&res.Content,
			&res.CreateTime,
		)
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
