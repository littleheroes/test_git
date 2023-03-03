package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
	//"net/http"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
func Logger() *logrus.Logger {
	now := time.Now()
	now.Format("2012")
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs"
	}
	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		fmt.Println(err.Error())
	}

	//日志文件
	logFileName := now.Format("2006-01-02") + ".log"
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			fmt.Println(err.Error())
		}
	}

	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	check(err)
	//实例化
	logger := logrus.New()

	//设置输出
	logger.Out = src

	//设置日志级别

	logger.SetLevel(logrus.DebugLevel)
	//设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return logger
}

func LoggerToFile() gin.HandlerFunc {
	logger := Logger()
	return func(c *gin.Context) {
		//开始时间
		startTime := time.Now()
		//处理请求
		c.Next()
		//结束时间
		endTime := time.Now()

		//执行时间

		latencyTime := endTime.Sub(startTime)

		//请求方式
		reqMethod := c.Request.Method

		//请求路由
		reqUri := c.Request.RequestURI

		//状态码
		statusCode := c.Writer.Status()
		//请求ip
		clientIP := c.ClientIP()

		//日志格式
		logger.Infof("| %3d | %13v |%15s | %s |  %s|",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}

type userlist struct {
	Id   int    `db:"id" json:"id"`
	Sex  string `db:"sex" json:"sex"`
	Name string `db:"name" json:"name" `
}

func getZhang(the_id int) []map[string]interface{} {
	var split []map[string]interface{}
	db, err := sql.Open("mysql", "root:123456@tcp(10.200.195.142:3306)/zhang")
	if err != nil {
		return split
	}

	defer db.Close()

	rows, err := db.Query("select * from test where id=?", the_id)
	check(err)
	for rows.Next() {
		var s userlist
		rows.Scan(&s.Id, &s.Sex, &s.Name)
		fmt.Println(s)
		var a map[string]interface{}
		a = make(map[string]interface{})
		a["id"] = &s.Id
		a["sex"] = &s.Sex
		a["name"] = &s.Name
		split = append(split, a)
	}
	fmt.Printf("执行成功")
	return split
}

func main() {
	//r := gin.Default()
	//r.GET("/test", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"123": "321",
	//	})
	//})
	//r.Run(":8080")"
	//db, err := sql.Open("mysql", "root:123456@tcp(10.200.195.142:3306)/zhang")
	//if err != nil {
	//	fmt.Printf("链接出现错误1")
	//	return
	//}
	//err = db.Ping() //尝试链接数据库
	//if err != nil {
	//	fmt.Printf("链接出现错误2")
	//	return
	//}
	//fmt.Printf("数据库链接成功")

	//r := gin.Default()
	//r.Use(LoggerToFile())
	//var the_id int
	//fmt.Scan(&the_id)
	//r.GET("/zhang", func(c *gin.Context) {
	//	Logger().WithFields(logrus.Fields{
	//		"name": "zhang",
	//	}).Info("记录一下日志", "info")
	//	c.JSON(200, gin.H{
	//		"data":    getZhang(the_id),
	//		"staus":   200,
	//		"message": "成功读取数据",
	//	})
	//})
	//r.Run(":8080")
	r := setupRouter()
	if err := r.Run(); err != nil {
		fmt.Println("startup service faild,err")
	}
}
