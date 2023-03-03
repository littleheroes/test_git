package main

import (
	"encoding/json"
	"fmt"
)

type Student_ struct {
	Name  interface{} `json:"name"`
	Age   interface{}
	Sex   interface{}
	Class *Class
}
type Class struct {
	Name  interface{}
	Grade interface{}
}

func main() {
	var stu Student_
	fmt.Println("this is a joke")
	stu.Name = "张三"
	stu.Age = 18
	stu.Sex = "男"
	cla := new(Class)
	cla.Name = "二班"
	cla.Grade = "一年级"
	stu.Class = cla

	jsonStu, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("---===============-")
		fmt.Println(err)
	}

	fmt.Println(string(jsonStu))

}
