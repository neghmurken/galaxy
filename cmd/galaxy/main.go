package main

import (
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
	m "github.com/neghmurken/galaxy/pkg/model"
	r "github.com/neghmurken/galaxy/pkg/render"
)

func main() {
	rl.InitWindow(0, 0, "Galaxy")
	rl.SetConfigFlags(rl.FlagVsyncHint)
	defer rl.CloseWindow()

	var scale float32 = 1

	rl.ToggleBorderlessWindowed()
	rl.SetTargetFPS(90)

	w, h := float32(rl.GetScreenWidth())*scale, float32(rl.GetScreenHeight())*scale

	c := m.MakeCosmos(3e7, m.Space{W: w, H: h})

	for i := 0; i < 1000; i++ {
		c.Add(m.NewStaticBody(
			m.RandVec(w, h),
			2*scale,
		))
	}

	t := r.NewTelescope(1 / scale)

loop:
	for {
		switch {
		case rl.IsKeyPressed(rl.KeyEscape) && rl.WindowShouldClose():
			break loop

		case rl.IsKeyPressed(rl.KeyPageDown):
			t.ZoomIn()

		case rl.IsKeyPressed(rl.KeyPageUp):
			t.ZoomOut()

		case rl.IsKeyDown(rl.KeyLeft):
			t.MoveLeft()

		case rl.IsKeyDown(rl.KeyRight):
			t.MoveRight()

		case rl.IsKeyDown(rl.KeyUp):
			t.MoveUp()

		case rl.IsKeyDown(rl.KeyDown):
			t.MoveDown()
		}

		c.Update(rl.GetFrameTime())
		t.Watch(c)
	}

	os.Exit(0)
}
