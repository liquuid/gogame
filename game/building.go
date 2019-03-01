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

type Building struct {
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

func (bl *Building) MoveTo() { fmt.Println(("move"))}
func (bl *Building) ScaleTo(factor float64) {
	bl.scale = factor
}

func (bl *Building) Tick(screen *ebiten.Image)  {

	count++
	msg := fmt.Sprintf("%d %d %d",count, bl.frameNum, (count / 5) % bl.frameNum)
	ebitenutil.DebugPrint(screen, msg)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(bl.scale, bl.scale )
	op.GeoM.Translate(-float64(bl.frameWidth)/2, -float64(bl.frameHeight)/2)
	op.GeoM.Translate(bl.x, bl.y)

	i := (count / NumSprites) % bl.frameNum
	sx, sy := bl.frameOX, bl.frameOY+i*bl.frameHeight
	r := image.Rect(sx, sy, sx+bl.frameWidth, sy+bl.frameHeight)
	op.SourceRect = &r
	screen.DrawImage(bl.image, op)

	bl.x += bl.xv
	bl.y += bl.yv

	w, h := 512, 320
	if bl.x > float64(w) || bl.x < 0  {
		bl.xv *= -1
	}

	if bl.y > float64(h) || bl.y < 0 {
		bl.yv *= -1
	}

	if count > bl.frameNum * 1000{
		count = 0
	}
}

func (bl *Building) Init(x,y, scale float64,frameOX, frameOY, frameWidth,frameHeight, frameNum int ) {

	bl.frameOX = 0
	bl.frameOY = 77
	bl.frameWidth = 90
	bl.frameHeight = 77
	bl.frameNum = 4
	bl.x = rand.Float64()*511
	bl.y = rand.Float64()*319
	bl.xv = rand.Float64()*5
	bl.yv = rand.Float64()*5
	bl.scale = scale
	//fmt.Println(rand.Float64())
	img, _, err := image.Decode(bytes.NewReader(images.BossStandIMGR))
	if err != nil {
		log.Fatal(err)
	}

	bl.image, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

}