package glogger

import (
	"fmt"
	"testing"
)

func TestGlog(t *testing.T) {
	Printf("this is Printf")
	Infof("this is Infof")
	Errorf("this is Errorf")
	Warningf("this is Warningf")
	// Fatalf("this is Fatalf")

	fmt.Println("====================")
	g := Glogger{
		Levels: []GloggerLevel{
			LevelInfo,
			LevelError,
			LevelWarning,
			LevelFatal,
		},
	}
	g.Printf("this is Printf with ctx")
	g.Infof("this is Infof with ctx")
	g.Errorf("this is Errorf with ctx")
	g.Warningf("this is Warningf with ctx")
	// g.Fatalf("this is Fatalf with ctx")

	g.UseSyslog = true
	g.Printf("this is Printf with ctx")
	g.Infof("this is Infof with ctx")
	g.Errorf("this is Errorf with ctx")
	g.Warningf("this is Warningf with ctx")
	// g.Fatalf("this is Fatalf with ctx")

}
