package indexCon

import (
	"../../common/public"
	"../../common/response"
	"../../model/commonStruct"
	"../../model/indexModel"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
)

type SqlNature struct{}

func (a *SqlNature) BeforeActivation(h mvc.BeforeActivation) {
	h.Handle("POST", "/update/one/content", "UpdateOneContent") //更新备忘录
	h.Handle("POST", "/insert/one/content", "InsertOneContent") //新增
	h.Handle("GET", "/index/list", "FindIndexList")             //查询备忘录列表
	h.Handle("POST", "/choice/one/detail", "FindOneDetail")     //获取单个备忘录详情
}

//新增
func (a *SqlNature) InsertOneContent(ctx iris.Context) iris.Map {
	var (
		err error
	)
	values := commonStruct.HomePage{}
	err = ctx.ReadJSON(&values)
	if err != nil {
		return response.FailResponse(struct{}{}, err)
	}
	values.CreateTime = public.TimeNowToStr()
	err = indexModel.InsertOneContentModel(values)
	if err != nil {
		return response.FailResponse(struct{}{}, err)
	}
	return response.SuccessResponse(struct{}{})
}

//更新备忘录
func (a *SqlNature) UpdateOneContent(ctx iris.Context) iris.Map {
	var (
		err error
	)
	type request struct {
		IndexId int    `json:"index_id"`
		Content string `json:"content"`
	}
	values := request{}
	err = ctx.ReadJSON(&values)
	if err != nil {
		return response.FailResponse(struct{}{}, err)
	}
	err = indexModel.UpdateOneContentModel(values.IndexId, values.Content)
	if err != nil {
		return response.FailResponse(struct{}{}, err)
	}
	return response.SuccessResponse(struct{}{})
}

//查询备忘录列表
func (a *SqlNature) FindIndexList(ctx iris.Context) iris.Map {
	var (
		err        error
		page, size int
		result     = []interface{}{}
	)
	page, err = ctx.URLParamInt("page")
	if err != nil {
		fmt.Println(page, err)
	}
	size, err = ctx.URLParamInt("size")
	if err != nil {
		fmt.Println(size, err)
	}
	if page == -1 {
		page = 1
	}
	if size == -1 {
		size = 10
	}
	result, err = indexModel.IndexListModel(page, size)
	if err != nil {
		return response.FailResponse(result, err)
	}
	if len(result) == 0 {
		return response.SuccessResponse([]string{})
	}
	return response.SuccessResponse(result)
}

//获取单个备忘录详情
func (a *SqlNature) FindOneDetail(ctx iris.Context) iris.Map {
	var (
		err error
		res = commonStruct.HomePage{}
	)
	type request struct {
		IndexId int `json:"index_id"`
	}
	values := request{}
	err = ctx.ReadJSON(&values)
	if err != nil {
		log.Println(err)
	}
	res, err = indexModel.OneDetailModel(values.IndexId)
	if err != nil {
		return response.FailResponse(res, err)
	}
	return response.SuccessResponse(res)
}
