package window

import (
	"image"
	"reflect"
	"testing"
)

func TestGetIconImage(t *testing.T) {
	got := GetIconImage()
	if got == nil {
		t.Error("Must not null")
	}
	if !reflect.DeepEqual(reflect.TypeOf(got), reflect.TypeOf(&image.NRGBA{})){
		t.Errorf("Got %s, want %s", reflect.TypeOf(got),reflect.TypeOf(&image.NRGBA{}))
	}
}

