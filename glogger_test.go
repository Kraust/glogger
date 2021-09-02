package glogger

import (
	"testing"
)

func TestGlog(t *testing.T) {
	Printf("this is Printf")
	Infof("this is Infof")
	Errorf("this is Errorf")
	Warningf("this is Warningf")
	Fatalf("this is Fatalf")
}
