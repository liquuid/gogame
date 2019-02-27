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

type Room struct {
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

func (rm *Room) MoveTo() { fmt.Println(("move"))}
func (rm *Room) ScaleTo(factor float64) {
	rm.scale = factor
}

func (rm *Room) Tick(screen *ebiten.Image)  {

	count++
	msg := fmt.Sprintf("%d %d %d",count, rm.frameNum, (count / 5) % rm.frameNum)
	ebitenutil.DebugPrint(screen, msg)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(rm.scale, rm.scale )
	op.GeoM.Translate(-float64(rm.frameWidth)/2, -float64(rm.frameHeight)/2)
	op.GeoM.Translate(rm.x, rm.y)

	i := (count / numSprites) % rm.frameNum
	sx, sy := rm.frameOX, rm.frameOY+i*rm.frameHeight
	r := image.Rect(sx, sy, sx+rm.frameWidth, sy+rm.frameHeight)
	op.SourceRect = &r
	screen.DrawImage(rm.image, op)

	rm.x += rm.xv
	rm.y += rm.yv

	w, h := 512, 320
	if rm.x > float64(w) || rm.x < 0  {
		rm.xv *= -1
	}

	if rm.y > float64(h) || rm.y < 0 {
		rm.yv *= -1
	}

	if count > rm.frameNum * 1000{
		count = 0
	}
}

func (rm *Room) Init(x,y, scale float64,frameOX, frameOY, frameWidth,frameHeight, frameNum int ) {

	rm.frameOX = 0
	rm.frameOY = 77
	rm.frameWidth = 90
	rm.frameHeight = 77
	rm.frameNum = 4
	rm.x = rand.Float64()*511
	rm.y = rand.Float64()*319
	rm.xv = rand.Float64()*5
	rm.yv = rand.Float64()*5
	rm.scale = scale
	//fmt.Println(rand.Float64())
	img, _, err := image.Decode(bytes.NewReader(images.BossStandIMGR))
	if err != nil {
		log.Fatal(err)
	}

	rm.image, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

}