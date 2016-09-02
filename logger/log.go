package logger

import (
	"github.com/astaxie/beego/logs"
	"os"
)

var ErrLogs *logs.BeeLogger
var DBLogs *logs.BeeLogger

func init() {
	ErrorDiary()
	//DBDiary()
}

func ErrorDiary() {
	ErrLogs = logs.NewLogger(10000)
	ErrLogs.Async()
	ErrLogs.EnableFuncCallDepth(true)
	//创建err_diary.log文件
	os.Mkdir("diary", os.ModePerm)
	_, err := os.OpenFile("diary/err_diary.log", os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	ErrLogs.SetLogger("file", `{"filename":"./diary/err_diary.log"}`)
}

func DBDiary() {
	DBLogs = logs.NewLogger(10000)
	DBLogs.Async()
	DBLogs.EnableFuncCallDepth(true)
	//创建err_diary.log文件
	os.Mkdir("diary", os.ModePerm)
	_, err := os.OpenFile("diary/db_diary.log", os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	DBLogs.SetLogger("file", `{"filename":"./diary/db_diary.log"}`)
}
