package logrecoder

import (
	"os"
	"time"
	"strings"
	"log"
	"utils"
)

var (
	FileHandler    *os.File
	consoleTrace   *log.Logger //控制台记录所有日志
	consoleInfo    *log.Logger //控制台记录重要信息
	consoleWarning *log.Logger //控制台记录警告信息
	consoleError   *log.Logger //记录错误信息

	fileTrace   *log.Logger //控制台记录所有日志
	fileInfo    *log.Logger //控制台记录重要信息
	fileWarning *log.Logger //控制台记录警告信息
	fileError   *log.Logger //记录错误信息

	flag = log.Ldate | log.Ltime | log.Lshortfile

	goEnv = ""
)

func init() {
	date := strings.Split(time.Now().String(), " ")[0]
	dirPath := utils.RootDir + "/log/"
	createDir(dirPath)
	filepath := dirPath + date + ".log"
	FileHandler, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		log.Fatal("log:init:os.OpenFile:==", err)
	}

	consoleTrace = log.New(os.Stdout, "TRACE", flag)
	consoleInfo = log.New(os.Stdout, "INFO", flag)
	consoleWarning = log.New(os.Stdout, "WARNING", flag)
	consoleError = log.New(os.Stdout, "ERROR", flag)

	fileTrace = log.New(FileHandler, "TRACE", flag)
	fileInfo = log.New(FileHandler, "INFO", flag)
	fileWarning = log.New(FileHandler, "WARNING", flag)
	fileError = log.New(FileHandler, "ERROR", flag)

	sysEnv := os.Getenv("GOENV")
	if len(sysEnv) > 0 && (sysEnv == "production" || sysEnv == "testing" || sysEnv == "development") {
		goEnv = sysEnv
	} else {
		goEnv = "development"
	}

}

func createDir(path string) {
	_, err := os.Stat(path)
	if err == nil {
		return
	}
	if os.IsNotExist(err) {
		//创建目录
		err := os.MkdirAll(path, 0777)
		if err != nil {
			log.Fatal("log:init:createDir:os.MkdirAll:==", err)
		}
	}
	return
}

func TRACE(trace string) {
	if goEnv == "production" {
		fileTrace.Println(trace)
	} else {
		consoleTrace.Println(trace)
	}
}

func INFO(info string) {
	if goEnv == "production" {
		fileInfo.Println(info)
	} else {
		consoleInfo.Println(info)
	}
}

func WARNING(warning string) {
	if goEnv == "production" {
		fileWarning.Println(warning)
	} else {
		consoleWarning.Println(warning)
	}
}

func ERROR(error string) {
	if goEnv == "production" {
		fileError.Println(error)
	} else {
		consoleError.Println(error)
	}
}
