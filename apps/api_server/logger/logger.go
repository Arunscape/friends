package logger

import (
	"fmt"
	"log"
	"os"
)

const (
	ocd = iota
	debug
	info
	warn
	err
)

var levelMap = map[string]int{
	"OCD":   ocd,
	"DEBUG": debug,
	"INFO":  info,
	"WARN":  warn,
	"ERROR": err,
	"":      info, // default if env isn't set
}

var level = levelMap[os.Getenv("LOG_LEVEL")]

func jLog(level string, v ...interface{}) {
	log.Print("[", level, "] ", fmt.Sprint(v...))
}

func OCD(v ...interface{}) {
	if level <= ocd {
		jLog("OCD", v...)
	}
}

func Debug(v ...interface{}) {
	if level <= debug {
		jLog("DEBUG", v...)
	}
}

func Info(v ...interface{}) {
	if level <= info {
		jLog("INFO", v...)
	}
}

func Warn(v ...interface{}) {
	if level <= warn {
		jLog("WARN", v...)
	}
}

func Error(v ...interface{}) {
	if level <= err {
		jLog("ERROR", v...)
	}
}
