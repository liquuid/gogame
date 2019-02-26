package sprite

import (
	"bytes"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/liquuid/gogame/resources/images"
	"image"
	_ "image/png"
	"log"
	"math/rand"
)

type Sprite struct {
	x           float64
	y           float64
	xv 			float64
	yv			float64

	scale float64
	frameOX     int
	frameOY     int
	frameWidth  int
	frameHeight int
	frameNum    int
	image       *ebiten.Image

}
const (
	numSprites = 20
)
var (
	count = 0
)

func (pl *Sprite) MoveTo() { fmt.Println(("move"))}
func (pl *Sprite) ScaleTo(factor float64) {
	pl.scale = factor
}

func (pl *Sprite) Tick(screen *ebiten.Image)  {

	count++
	msg := fmt.Sprintf("%d %d %d",count, pl.frameNum, (count / 5) % pl.frameNum)
	ebitenutil.DebugPrint(screen, msg)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(pl.scale, pl.scale )
	op.GeoM.Translate(-float64(pl.frameWidth)/2, -float64(pl.frameHeight)/2)
	op.GeoM.Translate(pl.x, pl.y)

	i := (count / numSprites) % pl.frameNum
	sx, sy := pl.frameOX, pl.frameOY+i*pl.frameHeight
	r := image.Rect(sx, sy, sx+pl.frameWidth, sy+pl.frameHeight)
	op.SourceRect = &r
	screen.DrawImage(pl.image, op)

	pl.x += pl.xv
	pl.y += pl.yv

	w, h := 512, 320
	if pl.x > float64(w) || pl.x < 0  {
		pl.xv *= -1
	}

	if pl.y > float64(h) || pl.y < 0 {
		pl.yv *= -1
	}

	if count > pl.frameNum * 1000{
		count = 0
	}
}

func (pl *Sprite) Init(x,y, scale float64,frameOX, frameOY, frameWidth,frameHeight, frameNum int ) {

	pl.frameOX = 0
	pl.frameOY = 77
	pl.frameWidth = 90
	pl.frameHeight = 77
	pl.frameNum = 4
	pl.x = rand.Float64()*511
	pl.y = rand.Float64()*319
	pl.xv = rand.Float64()*5
	pl.yv = rand.Float64()*5
	pl.scale = scale
	//fmt.Println(rand.Float64())
	img, _, err := image.Decode(bytes.NewReader(images.BossStandIMGR))
	if err != nil {
		log.Fatal(err)
	}

	pl.image, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

}