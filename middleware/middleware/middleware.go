package middleware

import (
	"../../common/msg"
	"../../common/response"
	"../jwt"
	"errors"
	"github.com/kataras/iris"
)

func IsWorker(ctx iris.Context) {
	claims, bo := jwt.ValidateToken(ctx.GetHeader("token"))
	if bo {
		ctx.Values().Set("token_name", claims.UserName)
		ctx.Next()
	} else {
		response.FailResponse(struct{}{}, errors.New(msg.Msg4))
	}
}
