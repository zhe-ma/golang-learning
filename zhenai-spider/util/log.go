package util

import (
	"io"
	"log"
	"os"
)

var (
	TraceLog *log.Logger
	InfoLog  *log.Logger
	WarnLog  *log.Logger
	ErrorLog *log.Logger
)

// 每个包可以包含任意多个 init 函数，这些函数都会在程序执行开始的时候被调用。
// 所有被编译器发现的 init 函数都会安排在 main 函数之前执行。
// init 函数用在设置包、初始化变量或其他要在程序运行前优先完成的引导工作。
// 这里就说的很明确了，只要能被编译器发现，都会在 main 函数之前执行，也就是不会在每次引用改包的时候都执行。
func InitLogging() {
	exist, err := PathExists("./log")
	if err != nil {
		log.Fatalln("Failed to check log dir:", err)
	}

	if !exist {
		err = os.Mkdir("./log", os.ModePerm)
		if err != nil {
			log.Fatalln("Failed to create log dir:", err)
		}
	}

	file, err := os.OpenFile("./log/spider.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatalln("Failed to create log file:", err)
	}

	TraceLog = log.New(io.MultiWriter(file, os.Stdout), "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLog = log.New(io.MultiWriter(file, os.Stdout), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarnLog = log.New(io.MultiWriter(file, os.Stdout), "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLog = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
