package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/liquuid/gogame/window"
	game "github.com/liquuid/gogame/game"
	_ "image/jpeg"
	"log"
	"math/rand"
)

var (
	windowDecorated = flag.Bool("windowdecorated", true, "whether the window is decorated")
	terminated = errors.New("terminated")
	zoomScale = 1.0
	sprites = make([]*game.Sprite, numSprites)

)
func init() {
	
}

const (
	initScreenWidth  = 512
	initScreenHeight = 320
	initScreenScale  = 2
	numSprites = 5
)

func update(screen *ebiten.Image) error {
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		return terminated
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		ToggleZoom()
	}

	window.Update(screen)

	for _ , sprite := range sprites{
		sprite.ScaleTo(zoomScale)
		sprite.Tick(screen)
	}

	return nil
}

func ToggleZoom() {
	if zoomScale == 1.0{
		zoomScale = 0.5
	} else {
		zoomScale = 1
	}
}

func initSprites(){
	for i := 0; i < numSprites ; i++{
	//for _, sprite := range sprites {
		//fmt.Println(sprite)

		//fmt.Println(sprites[i])
		x := rand.Float64() * 511
		y := 150.0
		xv := rand.Float64() * 5
		yv := 0.0

		sprites[i] = new(game.Sprite)
		//sprite = new(game.Sprite)
		sprites[i].Init(x, y, xv, yv, 1, 0, 77, 90, 77, 4)
		//sprite.Init(x, y, xv,yv,1, 0,77,90,77,4)
		sprites[i].LoadAnimations()

	}

}

func main() {
	flag.Parse()

	fmt.Printf("Device scale factor: %0.2f\n", ebiten.DeviceScaleFactor())

	w, h := ebiten.ScreenSizeInFullscreen()
	fmt.Printf("Screen size in fullscreen: %d, %d\n", w, h)

	ebiten.SetWindowDecorated(*windowDecorated)
	initSprites()

	if err := ebiten.Run(update, initScreenWidth, initScreenHeight, initScreenScale, "Edificio Maia"); err != nil && err != terminated {
		log.Fatal(err)
	}

}
