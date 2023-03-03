package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string `json:"-"`             //“-”作用是不进行序列化，效果和将结构体字段写成小写一样。
	Age  int    `json:"Age,omitempty"` // “omitempty”作用是在序列化的时候忽略0值或空值。
	Id   int    `json:"idx,string"`    // 序列化时，类型转化为string
	Sex  string `json:"sex"`
}

func main() {
	stu := Student{Name: "Andy", Age: 0, Id: 100}

	ret, err := json.Marshal(stu)
	if err != nil {
		panic(err)
	}
	fmt.Println("json info1 ret=", string(ret))
}
