package common

import "github.com/kataras/iris"

func SuccessResponse(res interface{}) iris.Map {
	return iris.Map{
		"status":  200,
		"data":    res,
		"message": "Success",
	}
}

func FailResponse(res interface{}, err error) iris.Map {
	return iris.Map{
		"status":  404,
		"data":    res,
		"message": err.Error(),
	}
}

func SuccessStruct() iris.Map {
	return iris.Map{
		"status":  200,
		"data": struct {}{},
		"message": "Success",
	}
}

func FailStruct(err error) iris.Map {
	return iris.Map{
		"status":  404,
		"data": struct {}{},
		"message": err.Error(),
	}
}

func SuccessSlice() iris.Map {
	return iris.Map{
		"status":  200,
		"data": []string{},
		"message": "Success",
	}
}