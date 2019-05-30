package testTemplate

import (
	"../../config"
	"fmt"
	"github.com/kataras/iris/mvc"
)

type SqlNature struct{}

var (
	table1 = config.Config.Get("mysql.table1").(string)
)

func (a *SqlNature) BeforeActivation(h mvc.BeforeActivation) {
	h.Handle("GET", "/test/sql", "Test")
}

func (a *SqlNature) Test() {
	fmt.Println("table1", table1)
}
