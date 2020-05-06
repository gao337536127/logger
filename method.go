package logger

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"syscall"
	"time"
)

func generalInfo(skip int, log *LogMessage, message string) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Fprintln(os.Stderr, "未能获得程序运行信息")
	}
	pathname, filename := path.Split(file)
	log.pathname = pathname
	log.filename = filename
	log.lineno = uint(line)
	funcInfo := runtime.FuncForPC(pc)
	if funcInfo == nil {
		fmt.Fprintln(os.Stderr, "未能获得程序运行信息")
	}
	funName := funcInfo.Name()
	log.funcName = funName
	log.asctime = time.Now().UnixNano()
	log.thread = syscall.Getpid()
	log.threadName = os.Args[0]
	log.process = syscall.Getpid()
	log.message = message
}
func Debug(message string) {
	log := new(LogMessage)
	log.levelno = debug
	log.levelname = "Debug"
	skip := 2
	generalInfo(skip, log, message)
	AddMessage(log)
}

func Info(message string) {
	log := new(LogMessage)
	log.levelno = info
	log.levelname = "Info"
	skip := 2
	generalInfo(skip, log, message)
	AddMessage(log)
}
func Warning(message string) {
	log := new(LogMessage)
	log.levelno = warning
	log.levelname = "Warning"
	skip := 2
	generalInfo(skip, log, message)
	AddMessage(log)
}
func Err(message string) {
	log := new(LogMessage)
	log.levelno = err
	log.levelname = "Error"
	skip := 2
	generalInfo(skip, log, message)
	AddMessage(log)
}
func Critical(message string) {
	log := new(LogMessage)
	log.levelno = critical
	log.levelname = "Critical"
	skip := 2
	generalInfo(skip, log, message)
	AddMessage(log)
}
