// beelog project beelog.go
package beelog

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

var (
	level     = LevelTrace
	BeeLogger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
)

const (
	LevelTrace = iota
	LevelDebug
	LevelInfo
	LevelWarning
	LevelError
	LevelCritical
	LevelFatal
)

func SetLevel(l int) {
	level = l
}

func SetLogger(l *log.Logger) {
	BeeLogger = l
}

func logPrint(le int, format string, a ...interface{}) {
	if le < level {
		return
	}
	var s string
	switch le {
	case LevelDebug:
		s = "[D]"
	case LevelTrace:
		s = "[T]"
	case LevelWarning:
		s = "[W]"
	case LevelInfo:
		s = "[I]"
	case LevelError:
		s = "[E]"
	case LevelFatal:
		s = "[F]"
	case LevelCritical:
		s = "[C]"
	}
	// Retrieve the stack infos
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "<unknown>"
		line = -1
	} else {
		file = file[strings.LastIndex(file, "/")+1:]
	}
	BeeLogger.Printf(fmt.Sprintf("%s %s:%d - %s\n", s, file, line, format), a...)
	if le == LevelFatal {
		os.Exit(1)
	}
}

func Trace(v ...interface{}) {
	logPrint(LevelTrace, "%v", v)
}

func Debugf(format string, v ...interface{}) {
	logPrint(LevelDebug, format, v...)
}

func Debug(v ...interface{}) {
	logPrint(LevelDebug, "%v", v)
}

func Info(v ...interface{}) {
	logPrint(LevelInfo, "%v", v)
}

func Warn(v ...interface{}) {
	logPrint(LevelWarning, "%v", v)
}

func Error(v ...interface{}) {
	logPrint(LevelError, "%v", v)
}

func Critical(v ...interface{}) {
	logPrint(LevelCritical, "%v", v)
}
