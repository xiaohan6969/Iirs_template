package commonStruct

type HomePage struct {
	Id         int    `json:"id,omitempty" gorm:"primary_key"`
	Title      string `json:"title" gorm:"type:varchar(20);"`
	Content    string `json:"content" gorm:"type:varchar(20);"`
	CreateTime string `json:"create_time" gorm:"type:varchar(20);"`
	ImageList  string `json:"image_list" gorm:"type:varchar(20);"`
	IndexImg   string `json:"index_img" gorm:"type:varchar(20);"`
	Rgb        string `json:"rgb" gorm:"type:varchar(20);"`
	Status     int    `json:"status,omitempty" gorm:"type:int;default:1"`
	UserName   string `json:"user_name" gorm:"type:varchar(20);not null;"` //名称
}

type User struct {
	UserId     int    `json:"user_id" gorm:"primary_key"`                         //用户编号
	UserName   string `json:"user_name" gorm:"type:varchar(20);unique;not null;"` //名称
	PassWord   string `json:"pass_word" gorm:"type:varchar(20);not null;"`        //密码
	UserAge    string `json:"user_age" gorm:"type:varchar(256);not null;"`        //年龄
	UserSex    string `json:"user_sex" gorm:"type:varchar(128);not null;"`        //性别
	CreateTime string `json:"create_time" gorm:"type:varchar(128);not null;"`     //创建时间
	Token      string `json:"token" gorm:"type:varchar(256);default:''"`          //token
	OpenId     string `json:"openid" gorm:"type:varchar(128);unique;"`                   // 小程序Openid
}

type WxApp struct {
	Errcode     int    `json:"errcode"`
	OpenID      string `json:"openid"`
	Session_key string `json:"session_key"`
	Unionid     string `json:"unionid"`
	Errmsg      string `json:"errmsg"`
	Mobile      string `json:"mobile"`
}
