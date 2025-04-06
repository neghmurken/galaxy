package main

import (
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(0, 0, "Lenia")
	rl.SetConfigFlags(rl.FlagVsyncHint)
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	shouldExit := false

	for !shouldExit {
		if rl.IsKeyPressed(rl.KeyEscape) || rl.WindowShouldClose() {
			shouldExit = true
		}

		// draw here
	}

	os.Exit(0)
}
