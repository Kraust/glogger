package glogger

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type GloggerLevel string

const (
	GloggerLevelInfo    = "INFO"
	GloggerLevelError   = "ERROR"
	GloggerLevelWarning = "WARNING"
	GloggerLevelFatal   = "FATAL"
)

type Glogger struct {
	Levels []GloggerLevel
}

// LEVEL: file:line:function():
func (ctx *Glogger) glog(level string, format string, args ...interface{}) {
	pc, fn, ln, _ := runtime.Caller(1)
	name := strings.Split(runtime.FuncForPC(pc).Name(), ".")

	f := fmt.Sprintf("%s: %s:%d:%s(): %s", level, filepath.Base(fn), ln, name[len(name)-1], format)
	fmt.Println(fmt.Sprintf(f, args...))
}

func Printf(format string, args ...interface{}) {
	g := Glogger{}
	g.glog("INFO", format, args...)
}

func Infof(format string, args ...interface{}) {
	g := Glogger{}
	g.glog("INFO", format, args...)
}

func Errorf(format string, args ...interface{}) {
	g := Glogger{}
	g.glog("ERROR", format, args...)
}

func Warningf(format string, args ...interface{}) {
	g := Glogger{}
	g.glog("WARN", format, args...)
}

func Fatalf(format string, args ...interface{}) {
	g := Glogger{}
	g.glog("FATAL", format, args...)
	os.Exit(1)
}
