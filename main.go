package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	// "log"
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
	c.line = max(0, c.line-count)
}
func (c *Cursor) MoveDown(count int, maxLine int) {
	c.line = min(c.line+count, maxLine)
}

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(800, 450, "editor")
	defer rl.CloseWindow()

	rl.SetTargetFPS(120)

	cursor := NewCursor(0, 0)
	textHeight := 20
	textSpacing := 1

	lines := SampleLines

	for !rl.WindowShouldClose() {
		// screenWidth := rl.GetScreenWidth()
		// screenHeight := rl.GetScreenHeight()
		// fmt.Println(screenWidth, screenHeight)

		if rl.IsKeyPressed(rl.KeyL) {
			cursor.MoveRight(1, len(lines[cursor.line]))
		}
		if rl.IsKeyPressed(rl.KeyH) {
			cursor.MoveLeft(1)
		}
		if rl.IsKeyPressed(rl.KeyJ) {
			cursor.MoveDown(1, len(lines)-1)
			cursor.col = min(cursor.col, len(lines[cursor.line])-1)
		}
		if rl.IsKeyPressed(rl.KeyK) {
			cursor.MoveUp(1)
			cursor.col = min(cursor.col, len(lines[cursor.line])-1)
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		for i, l := range lines {
			rl.DrawTextEx(rl.GetFontDefault(), l, rl.Vector2{X: 0, Y: float32(i) * float32(textHeight)}, float32(textHeight), float32(textSpacing), rl.Black)
			if cursor.line == i {
				// stringWidth := rl.MeasureText(l[0:cursor.col], int32(textHeight))
				stringSize := rl.MeasureTextEx(rl.GetFontDefault(), l[0:cursor.col], float32(textHeight), float32(textSpacing))
				fmt.Println(stringSize)
				cursorPos := rl.Vector2{X: stringSize.X, Y: float32(cursor.line * textHeight)}
				cursorBottom := rl.Vector2{X: cursorPos.X, Y: cursorPos.Y + float32(textHeight)}
				rl.DrawLineV(cursorPos, cursorBottom, rl.Red)
			}
		}

		rl.EndDrawing()
	}
}
