package indexModel

import (
	"../../common/public"
	"../../config"
	"../../server/mysqlServer"
	"../commonStruct"
	"fmt"
	"strconv"
)

var (
	tOne = config.Config.Get("mysql.tOne").(string)
)

//查询备忘录列表
func IndexListModel(page, size int) ([]commonStruct.HomePage, error) {
	var (
		db  = mysqlServer.JzGorm
		res = &[]commonStruct.HomePage{}
	)
	defer func() {
		if err := db.Close();err != nil{
			fmt.Println(err)
		}
	}()
	err := db.Table(tOne).
		Where("status = 1").
		Limit(strconv.Itoa(size)).
		Offset(strconv.Itoa((page - 1) * size)).
		Order("create_time DESC").
		Scan(res).
		Error
	return *res, err
}

//获取单个备忘录详情
func OneDetailModel(index_id int) (commonStruct.HomePage, error) {
	var (
		db  = mysqlServer.JzGorm
		res = &commonStruct.HomePage{}
	)
	defer func() {
		if err := db.Close();err != nil{
			fmt.Println(err)
		}
	}()
	err := db.Table(tOne).
		Where("id = ? AND status = ?", index_id, 1).
		Scan(res).
		Error
	return *res, err
}

//新增
func InsertOneContentModel(values commonStruct.HomePage) error {
	var (
		DB  = mysqlServer.JzGorm
		err error
	)
	defer func() {
		if err := DB.Close();err != nil{
			fmt.Println(err)
		}
	}()
	values.CreateTime = public.TimeNowToStr()
	values.Status = 1
	err = DB.Table(tOne).
		Create(&values).
		Error
	return err
}

//更新备忘录
func UpdateOneContentModel(index_id int, content string) error {
	var (
		DB  = mysqlServer.JzGorm
		err error
	)
	defer func() {
		if err := DB.Close();err != nil{
			fmt.Println(err)
		}
	}()
	err = DB.Table(tOne).
		Where("id = ?", index_id).
		Update("content", content).
		Error
	return err
}
