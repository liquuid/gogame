package sprite

import (
	"bytes"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/liquuid/gogame/resources/images"
	"image"
	_ "image/png"
	"log"
)

type Sprite struct {
	x           float64
	y           float64
	xv 			float64
	yv			float64
	scale 		float64

	activeAnimation *animation
	animationsDB map[string]*animation
}

const (
	numSprites = 5
	numAnimations = 1
)

var (
	count = 0
	animations = make([]*animation, numAnimations)
	op = &ebiten.DrawImageOptions{}
)

type animation struct{
	frameOX     int
	frameOY     int
	frameWidth  int
	frameHeight int
	frameNum    int

	sequence *ebiten.Image
}

func (pl *Sprite) MoveTo() { fmt.Println(("move"))}

func (pl *Sprite) ScaleTo(factor float64) {
	pl.scale = factor

}

func (pl *Sprite) OnClick() {

}

func (pl *Sprite) Init(x,y,xv, yv, scale float64 ) {
	pl.x = x
	pl.y = y
	pl.xv = xv
	pl.yv = yv
	pl.scale = scale
	pl.animationsDB = make(map[string]*animation)
}

func (pl *Sprite) Walk(screen *ebiten.Image) {

	count++

	//op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(-float64(pl.activeAnimation.frameWidth)/2, -float64(pl.activeAnimation.frameHeight)/2)
	op.GeoM.Translate(pl.x, pl.y)

	i := (count / numSprites) % pl.activeAnimation.frameNum
	sx, sy := pl.activeAnimation.frameOX, pl.activeAnimation.frameOY+i*pl.activeAnimation.frameHeight
	r := image.Rect(sx, sy, sx+pl.activeAnimation.frameWidth, sy+pl.activeAnimation.frameHeight)
	op.SourceRect = &r
	//screen.DrawImage(pl.activeAnimation.sequence, op)

	pl.x += pl.xv
	pl.y += pl.yv

	w := 512
	if pl.x > float64(w) || pl.x < 0  {
		pl.xv *= -1
		if pl.xv > 0{
			pl.activeAnimation = pl.animationsDB["WalkR"]
		} else {
			pl.activeAnimation = pl.animationsDB["WalkL"]
		}
	}

	if count > pl.activeAnimation.frameNum * 1000{
		count = 0
	}
}


func (pl *Sprite) Tick(screen *ebiten.Image)  {
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Scale(pl.scale, pl.scale )
	pl.Walk(screen)
	screen.DrawImage(pl.activeAnimation.sequence, op)
}


func (pl *Sprite) LoadAnimations(){

	pl.animationsDB["WalkR"] = new(animation)
	pl.animationsDB["WalkR"].LoadAnimation(0,80,90,80, 7, images.BossWalkIMGR)

	pl.animationsDB["WalkL"] = new(animation)
	pl.animationsDB["WalkL"].LoadAnimation(0,80,90,80, 7, images.BossWalkIMGL)

	pl.animationsDB["StandR"] = new(animation)
	pl.animationsDB["StandR"].LoadAnimation(0,77,90,77, 4, images.BossStandIMGR)

	pl.animationsDB["StandL"] = new(animation)
	pl.animationsDB["StandL"].LoadAnimation(0,77,90,77, 4, images.BossStandIMGL)

	pl.activeAnimation = pl.animationsDB["WalkR"]
}

func (a *animation) LoadAnimation(frameOX, frameOY, frameWidth, frameHeight, frameNum int, b []byte){
	a.frameOX = frameOX
	a.frameOY = frameOY
	a.frameWidth = frameWidth
	a.frameHeight = frameHeight
	a.frameNum = frameNum

	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}

	a.sequence, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
}

