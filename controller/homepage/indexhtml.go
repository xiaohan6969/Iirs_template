package homepage

import (
	"fmt"
	"github.com/kataras/iris"
)

func IndexHtml(ctx iris.Context) {
	_, err := ctx.HTML("<img class='fancybox-image' src='https://backiee.com/static/wpdb/wallpapers/3840x2160/188541.jpg' data-spm-anchor-id='a2c4g.11186623.0.i0.529242e4u4Ia7l'>")
	//_, unusual := ctx.HTML("<img class='fancybox-image' irisTemplate='http://static-aliyun-doc.oss-cn-hangzhou.aliyuncs.com/assets/img/17655/154157532010314_zh-CN.png'")
	if err != nil {
		fmt.Println(err)
	}
}
