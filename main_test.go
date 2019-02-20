package main

import (
	"errors"
	"testing"
)

func TestScreenGeometry(t *testing.T){
	if initScreenWidth != 512{
		t.Errorf("Got %d, want %d", initScreenWidth, 512)
	}
	if initScreenHeight != 320{
		t.Errorf("Got %d, want %d", initScreenHeight, 320)
	}
}

func TestTerminated(t *testing.T){
	err := errors.New("terminated")

	if (terminated.Error() != err.Error()){
		t.Errorf("Got %s, want %s", terminated.Error(), err.Error())
	}
}