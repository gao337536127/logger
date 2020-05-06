package logger

import (
	"fmt"
	"time"
)

func show(logMessage *LogMessage) {
	tm := time.Unix(0, logMessage.asctime)
	timeStr := tm.Format("2006-01-02 15:04:05.999")
	pathname := logMessage.pathname
	filename := logMessage.filename
	funcName := logMessage.funcName
	levelname := logMessage.levelname
	message := logMessage.message

	printStr := fmt.Sprintf("%s %s %s %s [%s]: %s", timeStr, pathname, filename, funcName, levelname, message)
	fmt.Println(printStr)
}
