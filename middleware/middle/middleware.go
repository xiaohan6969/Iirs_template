package middle

import (
	"../../common/msg"
	"../../common/response"
	"../jwt"
	"errors"
	"fmt"
	"github.com/kataras/iris"
)

func CheckToken(ctx iris.Context) {
	claims, bo := jwt.ValidateToken(ctx.GetHeader("token"))
	fmt.Println(claims, bo, ctx.GetHeader("token"))
	if bo {
		ctx.Values().Set("user_name", claims.UserName)
		ctx.Next()
	} else {
		_, _ = ctx.JSON(response.FailResponse(struct{}{}, errors.New(msg.Msg4)))
		return
	}
}
