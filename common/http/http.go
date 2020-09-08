package http

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"strings"
	"time"
)

func GetRequestBytes(url string) ([]byte, error) {
	var (
		err error
		b   []byte
		i   int
		arg = &fasthttp.Args{}
	)
	arg.Set("", "")
	b, i, err = FastHttpRequest(arg, "GET", url, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	if i == 200 {
		return b, nil
	}
	return nil, nil
}

func FastHttpRequest(arg *fasthttp.Args, method string, url string, cookies map[string]interface{}) ([]byte, int, error) {

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	switch method {
	case "GET":
		req.Header.SetMethod(method)
		req.Header.SetContentType("application/json")
		if arg.String() == "" {
			// 拼接url
			var str strings.Builder
			str.Write([]byte(url))
			str.WriteByte('?')
			str.WriteString(arg.String())
			str.String()
		}
	case "POST":
		req.Header.SetMethod(method)
		req.Header.SetContentType("application/json")
		arg.WriteTo(req.BodyWriter())
	}
	if cookies != nil {
		for key, v := range cookies {
			req.Header.SetCookie(key, v.(string))
		}
	}
	req.SetRequestURI(url)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	var err error
	err = fasthttp.DoTimeout(req, resp, time.Second*5)
	return resp.Body(), resp.StatusCode(), err
}
