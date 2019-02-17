package main

import (
	"bytes"
	"image"
	_ "image/png"
	"log"

	"./images"
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
	image       *ebiten.Image
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

	ebitenutil.DebugPrint(screen, "pota merda")

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(pl.frameWidth)/2, -float64(pl.frameHeight)/2)
	op.GeoM.Translate(pl.x, pl.y)
	i := (count / 5) % pl.frameNum
	sx, sy := pl.frameOX, pl.frameOY+i*pl.frameHeight
	r := image.Rect(sx, sy, sx+pl.frameWidth, sy+pl.frameHeight)
	op.SourceRect = &r
	screen.DrawImage(pl.image, op)
	return nil
}

func update(screen *ebiten.Image) error {

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	pl.update(screen)
	p2.update(screen)

	return nil
}

func main() {
	pl.frameOX = 0
	pl.frameOY = 77
	pl.frameWidth = 90
	pl.frameHeight = 77
	pl.frameNum = 4
	pl.x = 100
	pl.y = 100

	p2.frameOX = 0
	p2.frameOY = 80
	p2.frameWidth = 90
	p2.frameHeight = 80
	p2.frameNum = 7
	p2.x = 200
	p2.y = 100

	img, _, err := image.Decode(bytes.NewReader(images.BossStandIMG))
	if err != nil {
		log.Fatal(err)
	}
	img2, _, err := image.Decode(bytes.NewReader(images.BossWalkIMG))
	if err != nil {
		log.Fatal(err)
	}
	pl.image, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	p2.image, _ = ebiten.NewImageFromImage(img2, ebiten.FilterDefault)

	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "Animation two entities"); err != nil {
		log.Fatal(err)
	}
}
