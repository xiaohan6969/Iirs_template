package model

type Test struct {
	Id   int    `sql:"id"`
	Name string `sql:"name"`
	Sex  string `sql:"sex"`
	Age  int    `sql:"age"`
}
