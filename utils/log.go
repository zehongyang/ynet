package utils

import (
	"fmt"
	log2 "log"
	"os"
	"time"
)

var (
	logger *log2.Logger
)

//初始化log
func init()  {
	//将日志记录到runtime/log目录下
	var logDir = "./runtime/log"
	_, err := os.Stat(logDir)
	//如果运用下没日志目录则创建
	if os.IsNotExist(err) {
		err := os.MkdirAll(logDir, 0755)
		if err != nil {
			fmt.Println("mkdir runtime log err:",err)
			return
		}
	}

	filename := fmt.Sprintf("%s.%s",time.Now().Format("2006-01-02"),"log")
	file, err := os.OpenFile("../runtime/log/"+filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("init log err:",err)
		return
	}
	logger = log2.New(file, "logger: ", log2.Lshortfile)
}

//记录信息到日志
func Info(msg ...interface{}) {
	logger.Println(msg)
}
