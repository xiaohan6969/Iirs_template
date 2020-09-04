package indexModel

import (
	"../../config"
	config2 "../../server/mysqlServer"
	"../commonStruct"
	"fmt"
	"strconv"
)

var (
	table1 = config.Config.Get("mysql.table1").(string)
)

func IndexListModel(page, size int) ([]interface{}, error) {
	var (
		db     = config2.Mysql
		res    = commonStruct.DetailedQuery{}
		result = []interface{}{}
	)

	//查询数据，指定字段名，返回sql.Rows结果集
	sql := "select id,title,content,create_time,image_list,index_img,rgb from " + table1 + " limit " + strconv.Itoa((page-1)*size) + "," + strconv.Itoa(size)
	querySet, err := db.Query(sql)
	if err != nil {
		return result, err
	}
	for querySet.Next() {
		err = querySet.Scan(
			&res.Id,
			&res.Title,
			&res.Content,
			&res.CreateTime,
			&res.ImageList,
			&res.IndexImg,
			&res.Rgb,
		)
		result = append(result, commonStruct.DetailedQuery{
			res.Id,
			res.Title,
			res.Content,
			res.CreateTime,
			res.ImageList,
			res.IndexImg,
			res.Rgb,
		})
	}
	defer func() {
		if err := querySet.Close(); err != nil {
			fmt.Println("close fail")
		}
	}()
	return result, err
}

func OneDetailModel(index_id int) (commonStruct.DetailedQuery, error) {
	var (
		db  = config2.Mysql
		res = commonStruct.DetailedQuery{}
	)
	//查询数据，指定字段名，返回sql.Rows结果集
	sql := "select id,title,content,create_time from " + table1 + " where id = " + strconv.Itoa(index_id)
	querySet, err := db.Query(sql)
	if err != nil {
		return res, err
	}
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
	return res, err
}

func InsertOneContentModel(values commonStruct.DetailedQuery) error {
	var (
		db = config2.Mysql
	)
	//查询数据，指定字段名，返回sql.Rows结果集
	sql := "INSERT INTO "+table1+" (title, content,create_time,image_list,index_img,rgb )" +
		" VALUES ('"+values.Title +"','"+values.Content+"','"+values.CreateTime+"','"+values.ImageList+"','"+values.IndexImg+"','"+values.Rgb+"')"
	fmt.Println(sql)
	_, err := db.Exec(sql)
	return err
}

func UpdateOneContentModel(index_id int, content string) error {
	var (
		db = config2.Mysql
	)
	//查询数据，指定字段名，返回sql.Rows结果集
	sql := "update " + table1 + " set content = '" + content + "' where id = " + strconv.Itoa(index_id)
	fmt.Println(sql)
	_, err := db.Exec(sql)
	return err
}
