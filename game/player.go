package sprite

import (
	"bytes"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/liquuid/gogame/resources/images"
	"image"
	"image/color"
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

	numAnimations = 1
)

var (
	NumSprites int
	InitScreenWidth int
	InitScreenHeight int
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

func (s *Sprite) MoveTo() { fmt.Println(("move"))}

func (s *Sprite) ScaleTo(factor float64) {
	s.scale = factor

}

func (s *Sprite) OnClick() {

}

func (s *Sprite) Init(x,y,xv, yv, scale float64 ) {
	s.x = x
	s.y = y
	s.xv = xv
	s.yv = yv
	s.scale = scale
	s.animationsDB = make(map[string]*animation)
}

func (s *Sprite) Walk(screen *ebiten.Image) {

	count++

	//op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(-float64(s.activeAnimation.frameWidth)/2, -float64(s.activeAnimation.frameHeight)/2)
	op.GeoM.Translate(s.x, s.y)

	i := (count / NumSprites) % s.activeAnimation.frameNum
	sx, sy := s.activeAnimation.frameOX, s.activeAnimation.frameOY+i*s.activeAnimation.frameHeight
	r := image.Rect(sx, sy, sx+s.activeAnimation.frameWidth, sy+s.activeAnimation.frameHeight)
	op.SourceRect = &r
	//screen.DrawImage(s.activeAnimation.sequence, op)

	s.x += s.xv
	s.y += s.yv

	w := 512
	if s.x > float64(w) || s.x < 0  {
		s.xv *= -1
		if s.xv > 0{
			s.activeAnimation = s.animationsDB["WalkR"]
		} else {
			s.activeAnimation = s.animationsDB["WalkL"]
		}
	}

	if count > s.activeAnimation.frameNum * 1000{
		count = 0
	}
}


func (s *Sprite) Tick(screen *ebiten.Image)  {
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Scale(s.scale, s.scale )
	s.Walk(screen)
	screen.DrawImage(s.activeAnimation.sequence, op)
}


func (s *Sprite) LoadAnimations(){

	s.animationsDB["WalkR"] = new(animation)
	s.animationsDB["WalkR"].LoadAnimation(0,80,90,80, 7, images.BossWalkIMGR)

	s.animationsDB["WalkL"] = new(animation)
	s.animationsDB["WalkL"].LoadAnimation(0,80,90,80, 7, images.BossWalkIMGL)

	s.animationsDB["StandR"] = new(animation)
	s.animationsDB["StandR"].LoadAnimation(0,77,90,77, 4, images.BossStandIMGR)

	s.animationsDB["StandL"] = new(animation)
	s.animationsDB["StandL"].LoadAnimation(0,77,90,77, 4, images.BossStandIMGL)

	s.activeAnimation = s.animationsDB["WalkR"]
}

// In returns true if (x, y) is in the sprite, and false otherwise.
func (s *Sprite) In(x, y int) bool {
	// Check the actual color (alpha) value at the specified position
	// so that the result of In becomes natural to users.

	return s.activeAnimation.sequence.At(x - int(s.x) , y - int(s.y) ).(color.RGBA).A > 0
}

// MoveBy moves the sprite by (x, y).
func (s *Sprite) MoveBy(x, y int) {
	w, h := s.activeAnimation.sequence.Size()

	s.x += float64(x)
	s.y += float64(y)

	if s.x < 0 {
		s.x = 0
	}
	if s.x > float64(InitScreenWidth - w) {
		s.x = float64(InitScreenWidth - w)
	}
	if s.y < 0 {
		s.y = 0
	}
	if s.y > float64(InitScreenWidth - h) {
		s.y = float64(InitScreenWidth - h)
	}
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

