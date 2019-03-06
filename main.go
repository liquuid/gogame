package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/liquuid/gogame/input"

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
	strokes = make(map[*input.Stroke]struct{})
	cam_x float64
	cam_y float64
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
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		//s := NewStroke(&MouseStrokeSource{})
		//s.SetDraggingObject(g.spriteAt(s.Position()))
		//g.strokes[s] = struct{}{}
		X, Y := ebiten.CursorPosition()
		fmt.Printf("apertou %d %d", X, Y )
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		//s := NewStroke(&MouseStrokeSource{})
		//s.SetDraggingObject(g.spriteAt(s.Position()))
		//g.strokes[s] = struct{}{}
		X, Y := ebiten.CursorPosition()
		fmt.Printf("souto %d %d", X, Y )
	}


	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		return terminated
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		ToggleZoom()
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		cam_x += 50
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		cam_x += -50
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		cam_y += -50
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		cam_y += 50
	}

	window.Update(screen)

	for _ , sprite := range sprites{
		sprite.ScaleTo(zoomScale)
		sprite.Tick(screen, cam_x, cam_y)
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		s := input.NewStroke(&input.MouseStrokeSource{})
		s.SetDraggingObject(spriteAt(s.Position()))
		strokes[s] = struct{}{}
	}
	for _, id := range inpututil.JustPressedTouchIDs() {
		s := input.NewStroke(&input.TouchStrokeSource{id})
		s.SetDraggingObject(spriteAt(s.Position()))
		strokes[s] = struct{}{}
	}

	for s := range strokes {
		updateStroke(s)
		if s.IsReleased() {
			delete(strokes, s)
		}
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	draggingSprites := map[*game.Sprite]struct{}{}

	for s := range strokes {
		if sprite := s.DraggingObject().(*game.Sprite); sprite != nil {
			draggingSprites[sprite] = struct{}{}
		}
	}

	/*for _, s := range sprites {
		if _, ok := draggingSprites[s]; ok {
			continue
		}
		s.Draw(screen, 0, 0, 1)
	}
	for s := range strokes {
		dx, dy := s.PositionDiff()
		if sprite := s.DraggingObject().(*game.Sprite); sprite != nil {
			sprite.Draw(screen, dx, dy, 0.5)
		}
	}*/

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
		sprites[i].Init(x, y, xv, yv, 1)
		sprites[i].LoadAnimations()

	}

}


func spriteAt(x, y int) *game.Sprite {
	// As the sprites are ordered from back to front,
	// search the clicked/touched sprite in reverse order.
	for i := len(sprites) - 1; i >= 0; i-- {
		s := sprites[i]
		if s.In(x, y) {
			return s
		}
	}
	return nil
}

func updateStroke(stroke *input.Stroke) {
	stroke.Update()
	if !stroke.IsReleased() {
		return
	}

	s := stroke.DraggingObject().(*game.Sprite)
	if s == nil {
		return
	}

	s.MoveBy(stroke.PositionDiff())

	index := -1
	for i, ss := range sprites {
		if ss == s {
			index = i
			break
		}
	}

	// Move the dragged sprite to the front.
	sprites = append(sprites[:index], sprites[index+1:]...)
	sprites = append(sprites, s)

	stroke.SetDraggingObject(nil)
}

func main() {

	game.NumSprites = numSprites
	game.InitScreenWidth = initScreenWidth
	game.InitScreenHeight = initScreenHeight


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
