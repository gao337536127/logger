package logger

import (
	"fmt"
	"os"
	"time"
)

func write(logMessage *LogMessage) {
	tm := time.Unix(0, logMessage.asctime)
	timeStr := tm.Format("2006-01-02 15:04:05.999")
	pathname := logMessage.pathname
	filename := logMessage.filename
	funcName := logMessage.funcName
	levelname := logMessage.levelname
	message := logMessage.message

	printStr := fmt.Sprintf("%s %s %s %s [%s]: %s", timeStr, pathname, filename, funcName, levelname, message)

	dateStr := time.Now().Format("2006-01-02")
	logFileName := fmt.Sprintf("./out.%s.log", dateStr)

	file, err := os.OpenFile(logFileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "open log file %s error", logFileName)
	}
	defer file.Close()
	fmt.Fprintln(file, printStr)
}
