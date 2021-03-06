package logger

import (
	"container/list"
)

//LogMessage 基本的日志数据结构
type LogMessage struct {
	levelno    uint
	levelname  string
	pathname   string
	filename   string
	funcName   string
	lineno     uint
	asctime    int64
	thread     int
	threadName string
	process    int
	message    string
}

const (
	debug uint = iota
	info
	warning
	err
	critical
)

//ProcessFunc 日志处理方法结构
type ProcessFunc func(logMessage *LogMessage)

var messageChannels = list.New()

func CreateLogProcessGoroutine(processFunc ProcessFunc) {
	processChannel := make(chan *LogMessage)
	messageChannels.PushBack(processChannel)
	process := func() {
		for {
			logMessage := <-processChannel
			processFunc(logMessage)
		}
	}
	go process()
}
func AddMessage(logMessage *LogMessage) {
	for messageChannel := messageChannels.Front(); messageChannel != nil; messageChannel = messageChannel.Next() {
		channel := messageChannel.Value.(chan *LogMessage)
		channel <- logMessage
	}
}
func init() {
	CreateLogProcessGoroutine(show)
	CreateLogProcessGoroutine(write)
}
