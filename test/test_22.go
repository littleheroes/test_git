package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "\"{\\\"aas23\\\":\\\"aw12s\\\",\\\"bds12\\\":\\\"b32asd\\\"}\""
	fmt.Println(str)
	str = strings.Replace(str, "\\\"", "", -1)
	str = strings.Trim(str, "\"{")
	str = strings.Trim(str, "}\"")
	fmt.Println(str)
	s := strings.Split(str, ",")
	the_map := make(map[string]string)
	for i := 0; i < len(s); i++ {
		temp := strings.Split(s[i], ":")
		the_map[temp[0]] = temp[1]
	}
	for k, v := range the_map {
		fmt.Println(k)
		fmt.Println(v)
		fmt.Println("-------")
	}
}
