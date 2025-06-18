package main

import (
	"fmt"
)

type mode int

const (
	normalMode mode = iota + 1
	insertMode
	commandMode
)

type Editor struct {
	buffers []Buffer
	// cmd     string
	mode          mode
	currentBuffer uint
}

func NewEditor() Editor {
	buffers := make([]Buffer, 0)
	empty := NewEmptyBuffer()
	buffers = append(buffers, empty)
	return Editor{buffers, normalMode, 0}
}

func (e *Editor) OpenFile(filePath string) error {
	buf, err := NewBufferFromPath(filePath)
	if err != nil {
		return fmt.Errorf("Editor failed to open file: %v", err)
	}
	e.buffers = append(e.buffers, buf)
	e.currentBuffer = uint(max(0, int(len(e.buffers)-1)))
	return nil
}

func (e *Editor) DeleteCharBeforeCursor() {
	e.buffers[e.currentBuffer].DeleteCharsBeforeCursor(1)
}

func (e *Editor) SaveContents() error {
	err := e.buffers[e.currentBuffer].SaveContents()
	if err != nil {
		return fmt.Errorf("Editor failed to save contents: %v", err)
	}
	return nil
}

func (e *Editor) SaveContentsToPath(filePath string) error {
	err := e.buffers[e.currentBuffer].SaveContentsToPath(filePath)
	if err != nil {
		return fmt.Errorf("Editor failed to save contents: %v", err)
	}
	return nil
}

func (e *Editor) MoveCursorLeft(count uint) {
	buf := &(e.buffers[e.currentBuffer])
	buf.curs.MoveLeft(count)
}

func (e *Editor) MoveCursorRight(count uint) {
	buf := &(e.buffers[e.currentBuffer])
	lineLength := len(buf.contents[buf.curs.line])
	buf.curs.MoveRight(count, uint(lineLength))
}

func (e *Editor) MoveCursorDown(count uint) {
	buf := &(e.buffers[e.currentBuffer])
	lineCount := len(buf.contents)
	buf.curs.MoveDown(count, uint(lineCount))

	lineWidth := uint(len(buf.contents[buf.curs.line]) - 1)
	newCol := min(lineWidth, buf.curs.col)
	buf.curs.MoveTo(buf.curs.line, max(0, newCol))
}

func (e *Editor) MoveCursorUp(count uint) {
	buf := &(e.buffers[e.currentBuffer])
	buf.curs.MoveUp(count)

	lineWidth := uint(len(buf.contents[buf.curs.line]) - 1)
	newCol := min(lineWidth, buf.curs.col)
	buf.curs.MoveTo(buf.curs.line, max(0, newCol))
}

func (e *Editor) GetCursorPos() (line, col uint) {
	buf := &(e.buffers[e.currentBuffer])
	return buf.curs.GetPos()
}

func (e *Editor) GetBufferContents() []string {
	return e.buffers[e.currentBuffer].contents
}
