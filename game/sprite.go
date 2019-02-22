package sprite

import "github.com/hajimehoshi/ebiten"

type Sprite interface{
	Tick(screen *ebiten.Image)
	MoveTo()
	ScaleTo(factor float64)
}

type Sprites struct{
	spriteSlice []*Sprite
	Num int
}

func (s *Sprites) Add(sprt *Sprite){
	s.spriteSlice[0] = sprt

}
