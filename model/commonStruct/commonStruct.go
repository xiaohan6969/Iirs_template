package commonStruct

type DetailedQuery struct {
	Id         int    `json:"id" sql:"id"`
	Title      string `json:"title" sql:"title"`
	Content    string `json:"content" sql:"content"`
	CreateTime string `json:"create_time" sql:"create_time"`
	ImageList  string `json:"image_list" sql:"image_list"`
	IndexImg   string `json:"index_img" sql:"index_img"`
	Rgb        string `json:"rgb" sql:"rgb"`
}
