package window

import (
	"bytes"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/liquuid/gogame/resources/images"
	"image"
	"image/color"
	"log"
)

func Update(screen *ebiten.Image) error{
	ebiten.SetWindowIcon([]image.Image{GetIconImage()})


	screenScale := ebiten.ScreenScale()
	screenWidth, screenHeight := screen.Size()
	fullscreen := ebiten.IsFullscreen()
	tps := 6 // ebiten.MaxTPS()

	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		switch screenScale {
		case 0.75:
			screenScale = 1
		case 1:
			screenScale = 1.5
		case 1.5:
			screenScale = 2
		case 2:
			screenScale = 0.75
		default:
			panic("not reached")
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		fullscreen = !fullscreen
	}

	ebiten.SetScreenSize(screenWidth, screenHeight)
	ebiten.SetScreenScale(screenScale)
	ebiten.SetFullscreen(fullscreen)
	ebiten.SetRunnableInBackground(false)
	ebiten.SetCursorVisible(true)
	ebiten.SetVsyncEnabled(true)
	ebiten.SetMaxTPS(tps)

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	screen.Fill(color.RGBA{0x80, 0x80, 0xc0, 0xff})

	x, y := ebiten.CursorPosition()

	msg := fmt.Sprintf(`Press S key to change the window scale
Press F key to switch the fullscreen state
Press Q key to quit
Cursor: (%d, %d)
FPS: %0.2f
Device Scale Factor: %0.2f
ScreenSize %d %d`, x, y, ebiten.CurrentFPS(), ebiten.DeviceScaleFactor(),screenWidth, screenHeight )
	ebitenutil.DebugPrint(screen, msg)
	return nil
}

func GetIconImage() image.Image {
	icon, _, err := image.Decode(bytes.NewReader(images.Icon))
	if err != nil {
		log.Fatal(err)
	}
	return icon
}
