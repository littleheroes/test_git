package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func openTxt(txt string) map[string]int {
	ip_map := make(map[string]int)
	filePath := txt
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("文件打开失败 = ", err)
		return ip_map
	}
	defer file.Close()              // 关闭文本流
	reader := bufio.NewReader(file) // 读取文本数据

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		//var buffer bytes.Buffer
		str = strings.Replace(str, "\n", "", -1)
		ip_map[str] = 1
		//buffer.WriteString("http://")
		//buffer.WriteString(str)
		//buffer.WriteString(":8092/metrics")

		//str = buffer.String()
		//fmt.Println(str)
		//fmt.Print(str)
	}

	fmt.Println("文件读取结束")
	return ip_map
}

type the_body struct {
	Job   string `json:"job"`
	Token string `json:"token"`
}

type Result struct {
	Code    string     `json:"code"`
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    []The_data `json:"data"`
}
type The_data struct {
	Id              int     `json:"id"`
	Url             string  `json:"url"`
	Bearer_token    string  `json:"bearer_token"`
	Cluster         string  `json:"cluster"`
	Metrics_limit   int     `json:"metrics_limit"`
	Scrape_interval int     `json:"scrape_interval"`
	Tag             The_tag `json:"tag"`
	Update_at       string  `json:"update_at"`
	Create_at       string  `json:"create_at"`
}
type The_tag struct {
	Application   string `json:"application"`
	Business_line string `json:"business_line"`
	Cost_account  string `json:"cost_account"`
	Environment   string `json:"environment"`
	Exsi_host     string `json:"exsi_host"`
	Hostname      string `json:"hostname"`
	Idc           string `json:"idc"`
	Monitor_type  string `json:"monitor_type"`
	Target        string `json:"target"`
	Url           string `json:"url"`
}

func main() {
	url := "http://api.monitor.ke.com/kernel/task/list"
	contentType := "application/json"
	body_data := the_body{
		Job:   "lft_monitor",
		Token: "v4jl04x1-9nl6-pba2-wc9q-ei42gwcn1oea",
	}
	jsonstr, _ := json.Marshal(body_data)
	resp, err := http.Post(url, contentType, bytes.NewReader(jsonstr))
	//resp, err := http.Post(url, contentType, bytes.NewReader(jsonstr))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	var r Result
	result, _ := io.ReadAll(resp.Body)
	//fmt.Println(string(result))

	//err = json.Unmarshal([]byte(string(result)), &r)
	err = json.Unmarshal(result, &r)

	ip_map := openTxt("/Users/zph/Downloads/ip.txt")
	//fmt.Println(ip_map)
	for k1 := range ip_map {
		for key := range r.Data {
			//str := r.Data[key].URL
			str := r.Data[key].Url
			if strings.Contains(str, k1) {
				fmt.Printf("此IP号为：%d", r.Data[key].Id)
				fmt.Printf("  ")
				fmt.Printf("此IP地址为：%s\n", k1)
			}
		}
	}

}
