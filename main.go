package main

import (
	// "fmt"
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
	// "log"
)

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(800, 450, "editor")
	defer rl.CloseWindow()

	rl.SetTargetFPS(120)

	lines := SampleLines
	editor := NewEditor(&lines)
	cursor := &(editor.curs)
	textHeight := 20
	textSpacing := 1

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
		if rl.IsKeyPressed(rl.KeyBackspace) {
			editor.DeleteCharBeforeCursor()
		}
		// if rl.IsKeyPressed(rl.KeyLeftControl) && rl.IsKeyPressed(rl.KeyS) {
		if rl.IsKeyPressed(rl.KeyS) {
			// err := editor.SaveContents()
			err := editor.SaveContentsToPath("hello.txt")
			if err != nil {
				log.Fatal(err)
			}
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.DarkGray)
		for i, l := range lines {
			rl.DrawTextEx(rl.GetFontDefault(), l, rl.Vector2{X: 0, Y: float32(i) * float32(textHeight)}, float32(textHeight), float32(textSpacing), rl.RayWhite)
			if cursor.line == i {
				// stringWidth := rl.MeasureText(l[0:cursor.col], int32(textHeight))
				stringSize := rl.MeasureTextEx(rl.GetFontDefault(), l[0:cursor.col], float32(textHeight), float32(textSpacing))
				// fmt.Println(stringSize)
				cursorPos := rl.Vector2{X: stringSize.X, Y: float32(cursor.line * textHeight)}
				cursorBottom := rl.Vector2{X: cursorPos.X, Y: cursorPos.Y + float32(textHeight)}
				// rl.DrawLineV(cursorPos, cursorBottom, rl.Red)
				rl.DrawLineEx(cursorPos, cursorBottom, 3, rl.Red)
			}
		}

		rl.EndDrawing()
	}
}
