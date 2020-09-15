package userModel

import (
	"../../common/msg"
	"../../common/public"
	"../../middleware/jwt"
	"../../server/mysqlServer"
	"../commonStruct"
	"errors"
	"fmt"
)

func WxProgramLogin(openid string) (commonStruct.User, string, error) {
	var (
		err       error
		token     string
		user_name string
		pass_word = msg.Msg9
		db        = mysqlServer.JzGorm
		res_user  = &commonStruct.User{}
	)
	err = db.Table("users").
		Where("open_id = ?", openid).
		Scan(res_user).
		Error
	if err != nil {
		if err.Error()[0:10] == msg.Msg8 {
			user_name = msg.Msg7 + public.GetCaptcha() + public.RandString(2)
			token, _ = jwt.CreateToken(&jwt.Claims{UserName: user_name})
			user := &commonStruct.User{
				UserName:   user_name,
				PassWord:   pass_word,
				UserAge:    "",
				UserSex:    "",
				CreateTime: public.TimeNowToStr(),
				Token:      token,
				OpenId:     openid,
			}
			err = db.Create(user).Error
			return *user, token, err
		}
		return *res_user, token, err
	} else {
		user_name = res_user.UserName
		token, _ = jwt.CreateToken(&jwt.Claims{UserName: user_name})
		return *res_user, token, nil
	}
}

func RegisterNewUserModel(user_name, pass_word string) error {
	var (
		db = mysqlServer.JzGorm
	)
	user := &commonStruct.User{
		UserName:   user_name,
		PassWord:   pass_word,
		UserAge:    "",
		UserSex:    "",
		CreateTime: public.TimeNowToStr(),
		Token:      "",
	}
	return db.Create(user).Error
}

func LoginModel(user_name, pass_word string) (string, error) {
	var (
		db    = mysqlServer.JzGorm
		token string
		err   error
	)
	token, _ = jwt.CreateToken(&jwt.Claims{UserName: user_name})
	var count int
	err = db.Model(&commonStruct.User{}).
		Where(&commonStruct.User{UserName: user_name, PassWord: pass_word}).
		Count(&count).Error
	if err != nil {
		return "", err
	}
	if count != 1 {
		return "", errors.New(msg.Msg5)
	}
	return token, nil
}
