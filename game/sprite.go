package sprite

import "fmt"

type Sprites struct{
	spriteSlice []*Sprite
	Num int
}

func (s *Sprites) Add(sprt *Sprite){
	s.spriteSlice[0] = sprt

}

func (s *Sprites) Update() {
	for _ , sprite := range s.spriteSlice{
		fmt.Println(sprite)
	}
}