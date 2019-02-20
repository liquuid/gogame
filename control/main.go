package main

import (
	"bytes"
	"image"
	_ "image/png"
	"log"

	"github.com/liquuid/gogame/resources/images"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Player struct {
	x           float64
	y           float64
	frameOX     int
	frameOY     int
	frameWidth  int
	frameHeight int
	frameNum    int
	sprite      *ebiten.Image
	walkR       *ebiten.Image
	walkL       *ebiten.Image
	idleR       *ebiten.Image
	idleL       *ebiten.Image
	direction   int
}

const (
	screenWidth  = 320
	screenHeight = 240
)

var (
	count = 0
	pl    = Player{}
	p2    = Player{}
)

func (pl *Player) update(screen *ebiten.Image) error {
	count++

	ebitenutil.DebugPrint(screen, "W A S D")

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(pl.frameWidth)/2, -float64(pl.frameHeight)/2)
	op.GeoM.Translate(screenWidth/2, screenHeight/2)
	i := (count / 5) % pl.frameNum
	sx, sy := pl.frameOX, pl.frameOY+i*pl.frameHeight
	r := image.Rect(sx, sy, sx+pl.frameWidth, sy+pl.frameHeight)
	op.SourceRect = &r
	screen.DrawImage(pl.sprite, op)
	return nil
}

func update(screen *ebiten.Image) error {

	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		pl.frameOX = 0
		pl.frameOY = 80
		pl.frameWidth = 90
		pl.frameHeight = 80
		pl.frameNum = 7
		pl.direction = 1
		// Selects preloaded sprite
		pl.sprite = pl.walkL
		// Moves character 3px right
		//charX -= 3
	} else if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		pl.frameOX = 0
		pl.frameOY = 80
		pl.frameWidth = 90
		pl.frameHeight = 80
		pl.frameNum = 7
		pl.direction = 0
		// Selects preloaded sprite
		pl.sprite = pl.walkR
		// Moves character 3px left
		//charX += 3
	} else {
		pl.frameOX = 0
		pl.frameOY = 77
		pl.frameWidth = 90
		pl.frameHeight = 77
		pl.frameNum = 4

		if pl.direction == 1 {
			pl.sprite = pl.idleL
		} else {
			pl.sprite = pl.idleR
		}

	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	pl.update(screen)

	return nil
}

func main() {

	bsl, _, err := image.Decode(bytes.NewReader(images.BossStandIMGL))
	if err != nil {
		log.Fatal(err)
	}
	bwr, _, err := image.Decode(bytes.NewReader(images.BossWalkIMGR))
	if err != nil {
		log.Fatal(err)
	}
	bsr, _, err := image.Decode(bytes.NewReader(images.BossStandIMGR))
	if err != nil {
		log.Fatal(err)
	}
	bwl, _, err := image.Decode(bytes.NewReader(images.BossWalkIMGL))
	if err != nil {
		log.Fatal(err)
	}

	pl.idleL, _ = ebiten.NewImageFromImage(bsl, ebiten.FilterDefault)
	pl.idleR, _ = ebiten.NewImageFromImage(bsr, ebiten.FilterDefault)
	pl.walkL, _ = ebiten.NewImageFromImage(bwl, ebiten.FilterDefault)
	pl.walkR, _ = ebiten.NewImageFromImage(bwr, ebiten.FilterDefault)

	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "Animation two entities"); err != nil {
		log.Fatal(err)
	}
}
