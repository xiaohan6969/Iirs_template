package commonStruct

type HomePage struct {
	Id         int    `json:"id,omitempty" sql:"id"`
	Title      string `json:"title" sql:"title"`
	Content    string `json:"content" sql:"content"`
	CreateTime string `json:"create_time" sql:"create_time"`
	ImageList  string `json:"image_list" sql:"image_list"`
	IndexImg   string `json:"index_img" sql:"index_img"`
	Rgb        string `json:"rgb" sql:"rgb"`
	Status     int    `json:"status,omitempty" sql:"status,omitempty"`
}

type User struct {
	UserId     int    `json:"user_id" gorm:"primary_key"`                     //用户编号
	UserName   string `json:"user_name" gorm:"type:varchar(20);unique;not null;"`    //用户名称
	PassWord   string `json:"pass_word" gorm:"type:varchar(20);not null;"`    //用户密码
	UserAge    string `json:"user_age" gorm:"type:varchar(256);not null;"`    //用户年龄
	UserSex    string `json:"user_sex" gorm:"type:varchar(128);not null;"`    //用户性别
	CreateTime string `json:"create_time" gorm:"type:varchar(128);not null;"` //创建时间
	Token      string `json:"token" gorm:"type:varchar(256);default:''"`
}
