// beelog project beelog.go
package beelog

import "log"
import "os"

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
)

func SetLevel(l int) {
	level = l
}

func SetLogger(l *log.Logger) {
	BeeLogger = l
}

func Trace(v ...interface{}) {
	if level <= LevelTrace {
		BeeLogger.Printf("[T] %v\n", v)
	}
}

func Debug(v ...interface{}) {
	if level <= LevelDebug {
		BeeLogger.Printf("[D] %v\n", v)
	}
}

func Info(v ...interface{}) {
	if level <= LevelInfo {
		BeeLogger.Printf("[I] %v\n", v)
	}
}

func Warn(v ...interface{}) {
	if level <= LevelWarning {
		BeeLogger.Printf("[W] %v\n", v)
	}
}

func Error(v ...interface{}) {
	if level <= LevelError {
		BeeLogger.Printf("[E] %v\n", v)
	}
}

func Critical(v ...interface{}) {
	if level <= LevelCritical {
		BeeLogger.Printf("[C] %v\n", v)
	}
}
