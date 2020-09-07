package indexModel

import (
	"../../config"
	config2 "../../server/mysqlServer"
	"../commonStruct"
	"fmt"
	"reflect"
	"strconv"
)

var (
	table1 = config.Config.Get("mysql.table1").(string)
)

//查询备忘录列表
func IndexListModel(page, size int) ([]interface{}, error) {
	var (
		db     = config2.Mysql
		res    = commonStruct.HomePage{}
		result = []interface{}{}
	)

	//查询数据，指定字段名，返回sql.Rows结果集
	sql := "select id,title,content,create_time,image_list,index_img,rgb from " + table1 + " ORDER BY create_time DESC " +
		"limit " + strconv.Itoa((page-1)*size) + "," + strconv.Itoa(size)
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
		result = append(result, commonStruct.HomePage{
			Id:         res.Id,
			Title:      res.Title,
			Content:    res.Content,
			CreateTime: res.CreateTime,
			ImageList:  res.ImageList,
			IndexImg:   res.IndexImg,
			Rgb:        res.Rgb,
		})
	}
	defer func() {
		if err := querySet.Close(); err != nil {
			fmt.Println("close fail")
		}
	}()
	return result, err
}

//获取单个备忘录详情
func OneDetailModel(index_id int) (commonStruct.HomePage, error) {
	var (
		db  = config2.Mysql
		res = commonStruct.HomePage{}
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

func InsertSqlCommon(table string, title_list []string, values []string) string {
	var (
		title = "("
		value = " VALUES ('"
		l     = len(values)
	)
	for k, v := range title_list {
		if k+1 < l {
			title += v + ","
			value += values[k] + "','"
		} else {
			title += v + ")"
			value += values[k] + "')"
		}
	}
	return "INSERT INTO " + table + title + value
}

// 遍历struct并且自动进行赋值
func structByReflect(data map[string]interface{}, inStructPtr interface{}) {
	rType := reflect.TypeOf(inStructPtr)
	rVal := reflect.ValueOf(inStructPtr)
	if rType.Kind() == reflect.Ptr {
		// 传入的inStructPtr是指针，需要.Elem()取得指针指向的value
		rType = rType.Elem()
		rVal = rVal.Elem()
	} else {
		panic("inStructPtr must be ptr to struct")
	}
	// 遍历结构体
	for i := 0; i < rType.NumField(); i++ {
		t := rType.Field(i)
		f := rVal.Field(i)
		if v, ok := data[t.Name]; ok {
			f.Set(reflect.ValueOf(v))
		} else {
			panic(t.Name + " not found")
		}
	}
}

//新增
func InsertOneContentModel(values commonStruct.HomePage) error {
	var (
		db  = config2.Mysql
		err error
	)
	//查询数据，指定字段名，返回sql.Rows结果集
	_, err = db.Exec(commonStruct.InsertSqlCommonMap(table1, values))
	return err
}

//更新备忘录
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
