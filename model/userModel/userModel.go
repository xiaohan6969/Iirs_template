package userModel

import (
	"../../common/msg"
	"../../common/public"
	"../../middleware/jwt"
	"../../server/mysqlServer"
	"../commonStruct"
	"errors"
	"github.com/jinzhu/gorm"
)

func WxProgramLogin(openid string) (string, error) {
	var (
		err       error
		token     string
		user_name string
		pass_word = ""
		db        = mysqlServer.JzGorm
		res_user  = &commonStruct.User{}
	)
	err = db.Table("users").
		Where("openid = ?", openid).
		Scan(res_user).
		Error
	user_name = res_user.UserName
	if err == gorm.ErrRecordNotFound {
		user_name = msg.Msg7 + public.RandString(10)
		user := &commonStruct.User{
			UserName:   user_name,
			PassWord:   pass_word,
			UserAge:    "",
			UserSex:    "",
			CreateTime: public.TimeNowToStr(),
			Token:      "",
			OpenId:     openid,
		}
		err = db.Create(user).Error
		if err != nil {
			return token, err
		}
	}
	token, _ = jwt.CreateToken(&jwt.Claims{UserName: user_name})
	return token, nil
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
