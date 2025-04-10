package main

import (
	"math/rand"
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

	rl.ToggleFullscreen()
	rl.SetTargetFPS(90)

	w, h := float32(rl.GetScreenWidth())*scale, float32(rl.GetScreenHeight())*scale

	shouldExit := false

	c := m.MakeCosmos(9e6, m.Space{W: w, H: h})

	for i := 0; i < 200; i++ {
		c.Add(m.NewStaticBody(
			m.RandVec(w, h),
			rl.Remap(rand.Float32(), 0, 1, 1, 10)*scale,
		))
	}

	t := r.NewTelescope(1 / scale)

	for !shouldExit {
		if rl.IsKeyPressed(rl.KeyEscape) || rl.WindowShouldClose() {
			shouldExit = true
		}

		c.Update(rl.GetFrameTime())
		t.Watch(c)
	}

	os.Exit(0)
}
