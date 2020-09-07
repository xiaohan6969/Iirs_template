package log

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"os"
	"strings"
	"time"
	"../config"
)

var (
	deleteFileOnExit = config.Config.Get("log.deleteFileOnExit").(bool)
	osName = config.Config.Get("log.osName").(string)
)

//根据日期获取文件名，文件日志以最常用的方式工作
//但这些只是好的命名方式。
func todayFilename() string {
	today := time.Now().Format("Jan 02 2006")
	return osName + today + ".txt"
}
func newLogFile() *os.File {
	filename := todayFilename()
	//打开一个输出文件，如果重新启动服务器，它将追加到今天的文件中
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return f
}
var excludeExtensions = [...]string{
	".js",
	".css",
	".jpg",
	".png",
	".ico",
	".svg",
}
func ILogger() (h iris.Handler, close func() error) {
	close = func() error { return nil }
	c := logger.Config{
		Status:  true,
		IP:      true,
		Method:  true,
		Path:    true,
		Columns: true,
	}
	logFile := newLogFile()
	close = func() error {
		err := logFile.Close()
		if deleteFileOnExit {
			err = os.Remove(logFile.Name())
		}
		return err
	}
	c.LogFunc = func(now time.Time, latency time.Duration, status, ip, method, path string, message interface{}, headerMessage interface{}) {
		output := logger.Columnize(now.Format("2006/01/02 - 15:04:05"), latency, status, ip, method, path, message, headerMessage)
		logFile.Write([]byte(output))
	}
	//我们不想使用记录器，一些静态请求等
	c.AddSkipper(func(ctx iris.Context) bool {
		path := ctx.Path()
		for _, ext := range excludeExtensions {
			if strings.HasSuffix(path, ext) {
				return true
			}
		}
		return false
	})
	h = logger.New(c)
	return
}
