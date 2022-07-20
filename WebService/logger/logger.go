package logger

import (
	"fmt"
	"go/build"
	"log"
	"os"
	"runtime"
)

var NAMELOGFILE = "service.log"
var PATH = build.Default.GOPATH + "src/WebService/logs/"

type Logger struct {
	Filename string
	File     *os.File
}

func NewLogger() *Logger {
	Log := Logger{}
	filename :=  PATH + NAMELOGFILE
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(file)
	Log.File = file
	Log.Filename = filename
	return &Log
}

func (logger *Logger) INFO(s interface{}) {
	_, fn, line, _ := runtime.Caller(1)
	log.New(logger.File, "[INFO] ", log.Ldate|log.Ltime).Println(fmt.Sprintf("[%s:%d] %v", fn, line, s))
}

func (logger *Logger) WARN(s interface{}) {
	_, fn, line, _ := runtime.Caller(1)
	log.New(logger.File, "[WARN] ", log.Ldate|log.Ltime).Println(fmt.Sprintf("[%s:%d] %v", fn, line, s))
}

func (logger *Logger) ERROR(s interface{}) {
	_, fn, line, _ := runtime.Caller(1)
	log.New(logger.File, "[ERROR] ", log.Ldate|log.Ltime).Println(fmt.Sprintf("[%s:%d] %v", fn, line, s))
}
