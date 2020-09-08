package commonStruct

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

//sql 拼接，目前只支持全部string  等后续优化
func InsertSqlCommonMap(table string, m interface{}) string {
	var (
		title = "("
		value = " VALUES ('"
		err   error
	)
	fmt.Println("m===", m)
	var mapResult map[string]interface{}
	s, _ := json.Marshal(m)
	err = json.Unmarshal(s, &mapResult)
	if err != nil {
		fmt.Println("JsonToMapDemo unusual: ", err)
	}
	l := len(mapResult)
	i := 1
	for k, v := range mapResult {
		title += k
		f := reflect.ValueOf(v)
		switch f.Kind() {
		case reflect.String:
			value += v.(string)
		case reflect.Float64:
			value += strconv.FormatFloat(v.(float64), 'E', -1, 64)
		case reflect.Int:
			value += strconv.Itoa(v.(int))
		}
		if i < l {
			title += ","
			value += "','"
		} else {
			title += ")"
			value += "')"
		}
		i++
	}
	return "INSERT INTO " + table + title + value
}
