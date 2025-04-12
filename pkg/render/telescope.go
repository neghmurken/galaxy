package render

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
	m "github.com/neghmurken/galaxy/pkg/model"
)

var (
	MOVE_STEP float32 = 6
)

type Telescope struct {
	Offset m.Vec
	Zoom   float32
}

func NewTelescope(zoom float32) *Telescope {
	return &Telescope{
		Offset: m.Vec{X: float32(rl.GetScreenWidth()) / 2, Y: float32(rl.GetScreenHeight()) / 2},
		Zoom:   zoom,
	}
}

func (this *Telescope) ZoomIn() {
	this.Zoom *= .5
}

func (this *Telescope) ZoomOut() {
	this.Zoom /= .5
}

func (this *Telescope) MoveLeft() {
	this.Offset.X += MOVE_STEP / this.Zoom
}

func (this *Telescope) MoveRight() {
	this.Offset.X -= MOVE_STEP / this.Zoom
}

func (this *Telescope) MoveUp() {
	this.Offset.Y += MOVE_STEP / this.Zoom
}

func (this *Telescope) MoveDown() {
	this.Offset.Y -= MOVE_STEP / this.Zoom
}

func (this *Telescope) Watch(cosmos *m.Cosmos) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.NewColor(0x10, 0x10, 0x10, 0xff))

	for _, body := range cosmos.Bodies {
		rl.DrawCircle(
			int32(body.Pos.X*this.Zoom+this.Offset.X),
			int32(body.Pos.Y*this.Zoom+this.Offset.Y),
			max(body.Size*this.Zoom, 1),
			rl.ColorLerp(
				color.RGBA{0x38, 0x2A, 0x76, 0xFF},
				color.RGBA{0xFF, 0x91, 0x00, 0xFF},
				body.SizeGrowth/(body.SizeGrowth+body.Size)*2,
			),
		)

		rl.DrawLine(
			int32(body.Pos.X*this.Zoom+this.Offset.X),
			int32(body.Pos.Y*this.Zoom+this.Offset.Y),
			int32((body.Pos.X+body.Vel.X)*this.Zoom+this.Offset.X),
			int32((body.Pos.Y+body.Vel.Y)*this.Zoom+this.Offset.Y),
			rl.NewColor(0xff, 0xff, 0x00, 0x45),
		)
	}

	rl.EndDrawing()
}
