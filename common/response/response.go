package response

import "github.com/kataras/iris"

//成功
func SuccessResponse(res interface{}) iris.Map {
	return iris.Map{
		"status":  200,
		"data":    res,
		"message": "SUCCESS",
		"token":   "",
	}
}

//失败带错误信息
func FailResponse(res interface{}, err error) iris.Map {
	return iris.Map{
		"status":  404,
		"data":    res,
		"message": err.Error(),
		"token":   "",
	}
}

//成功返回token
func SuccessAndToken(res interface{}, message, token string) iris.Map {
	return iris.Map{
		"status":  200,
		"data":    res,
		"message": message,
		"token":   token,
	}
}
