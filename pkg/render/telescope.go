package render

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
	m "github.com/neghmurken/galaxy/pkg/model"
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

func (this *Telescope) Watch(cosmos *m.Cosmos) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	for _, body := range cosmos.Bodies {
		rl.DrawCircle(
			int32((body.Pos.X*this.Zoom + this.Offset.X)),
			int32((body.Pos.Y*this.Zoom + this.Offset.Y)),
			rl.Remap(body.Mass, 100, 10000, 1, 10)*this.Zoom,
			LerpColor(
				color.RGBA{0x28, 0x1A, 0x66, 0xFF},
				color.RGBA{0xFF, 0x91, 0x00, 0xFF},
				rl.Vector2Length(body.Vel)/1000,
			),
		)
	}

	rl.EndDrawing()
}

func LerpColor(left, right rl.Color, factor float32) rl.Color {
	return rl.NewColor(
		uint8(rl.Lerp(float32(left.R), float32(right.R), factor)),
		uint8(rl.Lerp(float32(left.G), float32(right.G), factor)),
		uint8(rl.Lerp(float32(left.B), float32(right.B), factor)),
		0xFF,
	)
}
