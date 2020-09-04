package indexModel

import (
	"../../config"
	config2 "../../server/mysqlServer"
	"../commonStruct"
	"encoding/json"
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

//获取单个备忘录详情
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

//sql 拼接，目前只支持全部string  等后续优化
func InsertSqlCommonMap(table string, m interface{}) string {
	var (
		title = "("
		value = " VALUES ('"
		err   error
	)
	var mapResult map[string]interface{}
	s, _ := json.Marshal(m)
	err = json.Unmarshal(s, &mapResult)
	if err != nil {
		fmt.Println("JsonToMapDemo err: ", err)
	}
	l := len(mapResult)
	i := 1
	for k, v := range mapResult {
		title += k
		f := reflect.ValueOf(v)
		fmt.Println(f.Kind())
		switch f.Kind() {
		case reflect.String:
			value += v.(string)
		case reflect.Float64:
			value += strconv.FormatFloat(v.(float64), 'E', -1, 64)
		case reflect.Int:
			value += strconv.Itoa(v.(int))
		}
		if i < l {
			title +=  ","
			value += "','"
		} else {
			title += ")"
			value += "')"
		}
		i++
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
func InsertOneContentModel(values commonStruct.DetailedQuery1) error {
	var (
		db  = config2.Mysql
		err error
	)
	//查询数据，指定字段名，返回sql.Rows结果集
	_, err = db.Exec(InsertSqlCommonMap(table1, values))
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
