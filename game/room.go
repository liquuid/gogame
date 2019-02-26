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

func (r *Room) MoveTo() { fmt.Println(("move"))}
func (r *Room) ScaleTo(factor float64) {
	r.scale = factor
}

func (r *Room) Tick(screen *ebiten.Image)  {

	count++
	msg := fmt.Sprintf("%d %d %d",count, r.frameNum, (count / 5) % r.frameNum)
	ebitenutil.DebugPrint(screen, msg)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(r.scale, r.scale )
	op.GeoM.Translate(-float64(r.frameWidth)/2, -float64(r.frameHeight)/2)
	op.GeoM.Translate(r.x, r.y)

	i := (count / numSprites) % r.frameNum
	sx, sy := r.frameOX, r.frameOY+i*r.frameHeight
	r := image.Rect(sx, sy, sx+r.frameWidth, sy+r.frameHeight)
	op.SourceRect = &r
	screen.DrawImage(r.image, op)

	r.x += r.xv
	r.y += r.yv

	w, h := 512, 320
	if r.x > float64(w) || r.x < 0  {
		r.xv *= -1
	}

	if r.y > float64(h) || r.y < 0 {
		r.yv *= -1
	}

	if count > r.frameNum * 1000{
		count = 0
	}
}

func (r *Room) Init(x,y, scale float64,frameOX, frameOY, frameWidth,frameHeight, frameNum int ) {

	r.frameOX = 0
	r.frameOY = 77
	r.frameWidth = 90
	r.frameHeight = 77
	r.frameNum = 4
	r.x = rand.Float64()*511
	r.y = rand.Float64()*319
	r.xv = rand.Float64()*5
	r.yv = rand.Float64()*5
	r.scale = scale
	//fmt.Println(rand.Float64())
	img, _, err := image.Decode(bytes.NewReader(images.BossStandIMGR))
	if err != nil {
		log.Fatal(err)
	}

	r.image, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

}