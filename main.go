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

	editor := NewEditor()
	err := editor.OpenFile("hello.txt")
	if err != nil {
		log.Fatal(err)
	}

	// cursor := &(editor.curs)
	textHeight := 20
	textSpacing := 1

	for !rl.WindowShouldClose() {
		// screenWidth := rl.GetScreenWidth()
		// screenHeight := rl.GetScreenHeight()
		// fmt.Println(screenWidth, screenHeight)

		if rl.IsKeyPressed(rl.KeyBackspace) {
			editor.DeleteCharBeforeCursor()
		}
		if rl.IsKeyPressed(rl.KeyS) {
			err := editor.SaveContentsToPath("hello.txt")
			if err != nil {
				log.Fatal(err)
			}
		}
		if rl.IsKeyPressed(rl.KeyH) {
			editor.MoveCursorLeft(1)
		}
		if rl.IsKeyPressed(rl.KeyJ) {
			editor.MoveCursorDown(1)
		}
		if rl.IsKeyPressed(rl.KeyK) {
			editor.MoveCursorUp(1)
		}
		if rl.IsKeyPressed(rl.KeyL) {
			editor.MoveCursorRight(1)
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.DarkGray)
		for i, l := range editor.GetBufferContents() {
			rl.DrawTextEx(rl.GetFontDefault(), l, rl.Vector2{X: 0, Y: float32(i) * float32(textHeight)}, float32(textHeight), float32(textSpacing), rl.RayWhite)
			cursLine, cursCol := editor.GetCursorPos()
			if int(cursLine) == i {
				// stringWidth := rl.MeasureText(l[0:cursCol], int32(textHeight))
				stringSize := rl.MeasureTextEx(rl.GetFontDefault(), l[0:cursCol], float32(textHeight), float32(textSpacing))
				// fmt.Println(stringSize)
				cursorPos := rl.Vector2{X: stringSize.X, Y: float32(int(cursLine) * textHeight)}
				cursorBottom := rl.Vector2{X: cursorPos.X, Y: cursorPos.Y + float32(textHeight)}
				// rl.DrawLineV(cursorPos, cursorBottom, rl.Red)
				rl.DrawLineEx(cursorPos, cursorBottom, 3, rl.Red)
			}
		}

		rl.EndDrawing()
	}
}
