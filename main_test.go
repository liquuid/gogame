package main

import (
	"errors"
	"github.com/liquuid/gogame/game"
	"reflect"
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

func TestAllocatedSprites(t *testing.T){
	t.Run("Test the number of sprites", func(t *testing.T){
		got := len(sprites)
		want := 200

		if got != want {
			t.Errorf("got %d, want %d", got , want)
		}
	})
	t.Run("Test the content of sprites", func(t *testing.T){
		initSprites()
		for _, sprt := range sprites {
			if sprt == nil  {
				t.Errorf("got %v, want %v", sprt,sprite.Sprite{} )
			}
		}
	})
	t.Run("Test the initial content of sprites", func(t *testing.T){
		initSprites()
		for _, sprt := range sprites {
			if reflect.TypeOf(*sprt) != reflect.TypeOf(sprite.Sprite{}) {
				t.Errorf("got %v, want %v", reflect.TypeOf(*sprt),reflect.TypeOf(sprite.Sprite{}) )
			}
		}
	})

}

func TestToggleZoom(t *testing.T) {
	t.Run("Test if togglezoom is working", func(t *testing.T){
		initial := zoomScale
		ToggleZoom()
		if initial == zoomScale{
			t.Errorf("got %v, want %v", zoomScale, 0.5)
		}
		ToggleZoom()
		if initial != zoomScale{
			t.Errorf("got %v, want %v", zoomScale, 1)
		}
	})
}