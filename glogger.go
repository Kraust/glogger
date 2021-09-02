package glogger

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type GloggerLevel string

const (
	LevelInfo    = "INFO"
	LevelError   = "ERROR"
	LevelWarning = "WARNING"
	LevelFatal   = "FATAL"
)

type Glogger struct {
	Levels    []GloggerLevel
	UseSyslog bool
}

func (g *Glogger) glog(level GloggerLevel, format string, args ...interface{}) {
	pc, fn, ln, _ := runtime.Caller(1)
	name := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	appname := filepath.Base(os.Args[0])

	for _, l := range g.Levels {
		if l == level {
			f := fmt.Sprintf("[%s] %s:%d:%s(): %s", level, filepath.Base(fn),
				ln, name[len(name)-1], format)
			if g.UseSyslog {
				lw, err := syslog.New(syslog.LOG_INFO, appname)
				if err == nil {
					log.SetOutput(lw)
					log.Print(fmt.Sprintf(f, args...))
				}
			} else {
				log.Println(fmt.Sprintf(f, args...))
			}
		}
	}
}

func (g *Glogger) Printf(format string, args ...interface{}) {
	g.glog(LevelInfo, format, args...)
}

func (g *Glogger) Infof(format string, args ...interface{}) {
	g.glog(LevelInfo, format, args...)
}

func (g *Glogger) Errorf(format string, args ...interface{}) {
	g.glog(LevelError, format, args...)
}

func (g *Glogger) Warningf(format string, args ...interface{}) {
	g.glog(LevelWarning, format, args...)
}

func (g *Glogger) Fatalf(format string, args ...interface{}) {
	g.glog(LevelFatal, format, args...)
	os.Exit(1)
}

/* These do not need a context */
func Printf(format string, args ...interface{}) {
	g := Glogger{
		Levels: []GloggerLevel{LevelInfo},
	}
	g.glog(LevelInfo, format, args...)
}

func Infof(format string, args ...interface{}) {
	g := Glogger{
		Levels: []GloggerLevel{LevelInfo},
	}
	g.glog(LevelInfo, format, args...)
}

func Errorf(format string, args ...interface{}) {
	g := Glogger{
		Levels: []GloggerLevel{LevelError},
	}
	g.glog(LevelError, format, args...)
}

func Warningf(format string, args ...interface{}) {
	g := Glogger{
		Levels: []GloggerLevel{LevelWarning},
	}
	g.glog(LevelWarning, format, args...)
}

func Fatalf(format string, args ...interface{}) {
	g := Glogger{
		Levels: []GloggerLevel{LevelFatal},
	}
	g.glog(LevelFatal, format, args...)
	os.Exit(1)
}
