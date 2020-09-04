package indexCon

import (
	"../../common"
	"../../model/commonStruct"
	"../../model/indexModel"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
)

type SqlNature struct{}

func (a *SqlNature) BeforeActivation(h mvc.BeforeActivation) {
	h.Handle("POST", "/update/one/content", "UpdateOneContent")
	h.Handle("POST", "/insert/one/content", "InsertOneContent")
	h.Handle("GET", "/index/list", "FindIndexList")
	h.Handle("POST", "/choice/one/detail", "FindOneDetail")
}

func (a *SqlNature) InsertOneContent(ctx iris.Context) iris.Map {
	var (
		err error
	)
	values := commonStruct.DetailedQuery{}
	err = ctx.ReadJSON(&values)
	if err != nil {
		return common.FailStruct(err)
	}
	err = indexModel.InsertOneContentModel(values)
	if err != nil {
		return common.FailStruct(err)
	}
	return common.SuccessStruct()
}

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
		return common.FailStruct(err)
	}
	err = indexModel.UpdateOneContentModel(values.IndexId, values.Content)
	if err != nil {
		return common.FailStruct(err)
	}
	return common.SuccessStruct()
}

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
		return common.FailResponse(result, err)
	}
	if len(result) == 0 {
		return common.SuccessResponse([]string{})
	}
	return common.SuccessResponse(result)
}

func (a *SqlNature) FindOneDetail(ctx iris.Context) iris.Map {
	var (
		err error
		res = commonStruct.DetailedQuery{}
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
		return common.FailResponse(res, err)
	}
	return common.SuccessResponse(res)
}
