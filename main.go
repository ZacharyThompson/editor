package main

import (
	// "log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Cursor struct {
	line int
	col  int
}

func NewCursor(line, col int) Cursor {
	return Cursor{line: line, col: col}
}
func (c *Cursor) MoveLeft(count int) {
	c.col = max(0, c.col-count)
}
func (c *Cursor) MoveRight(count, maxCol int) {
	c.col = min(c.col+count, maxCol)
}
func (c *Cursor) MoveUp(count int) {
	c.col = max(0, c.line-count)
}

func main() {
	font := rl.GetFontDefault()
	// font := rl.LoadFontEx("./assets/Iosevka-Regular.ttf", 20, nil, 0)
	// if !rl.IsFontValid(font) {
	// 	log.Fatal("Invalid Font")
	// }
	defer rl.UnloadFont(font)

	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(800, 450, "editor")
	defer rl.CloseWindow()

	rl.SetTargetFPS(120)

	cursor := NewCursor(0, 0)
	textWidth := 20
	textHeight := 20
	textSpacing := 1

	lines := SampleLines

	for !rl.WindowShouldClose() {
		// screenWidth := rl.GetScreenWidth()
		// screenHeight := rl.GetScreenHeight()
		// fmt.Println(screenWidth, screenHeight)

		if rl.IsKeyPressed(rl.KeyL) {
			cursor.MoveRight(1, 10)
		}
		if rl.IsKeyPressed(rl.KeyH) {
			cursor.MoveLeft(1)
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		for i, l := range lines {
			rl.DrawTextEx(font, l, rl.Vector2{X: 0, Y: float32(i * textHeight)}, float32(textHeight), float32(textSpacing), rl.Black)
		}

		cursorPos := rl.Vector2{X: float32(cursor.col * textWidth), Y: float32(cursor.line * textHeight)}
		cursorBottom := rl.Vector2{X: cursorPos.X, Y: cursorPos.Y + float32(textHeight)}
		rl.DrawLineV(cursorPos, cursorBottom, rl.Red)

		rl.EndDrawing()
	}
}
