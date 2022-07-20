package receiver

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

type Logger struct {
	Filename string
	File     *os.File
}

func NewLogger(filepath string) *Logger {
	filename := filepath
	logger := Logger{}
	
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	logger.File = file
	logger.Filename = filename
	return &logger
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
