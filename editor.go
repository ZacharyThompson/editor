package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

type Mode int

const (
	NormalMode Mode = iota + 1
	InsertMode
	CommandMode
)

type Editor struct {
	curs            Cursor
	buf             []string
	cmd             string
	mode            Mode
	currentFilePath string
}

func NewEditor(buf *[]string) Editor {
	return Editor{NewCursor(0, 0), *buf, "", NormalMode, ""}
}

func (e *Editor) DeleteCharBeforeCursor() {
	if e.curs.col == 0 {
		return
	}
	line := e.buf[e.curs.line]
	line = line[0:e.curs.col-1] + line[e.curs.col:]
	e.buf[e.curs.line] = line
	e.curs.MoveLeft(1)
}

func (e *Editor) SaveContents() error {
	if e.currentFilePath == "" {
		return errors.New("Cannot save contents: No filepath given")
	}
	// TODO: maybe check fs.ValidPath?
	file, err := os.Create(e.currentFilePath)
	if err != nil {
		return fmt.Errorf("Cannot save contents: Cannot create file: %v", err)
	}
	file.WriteString(strings.Join(e.buf, "\n") + "\n")
	return nil
}

func (e *Editor) SaveContentsToPath(filePath string) error {
	if filePath == "" {
		return errors.New("Cannot save contents to path: Empty filepath")
	}
	if !fs.ValidPath(filePath) {
		return errors.New("Cannot save contents to path: Invalid filepath")
	}
	e.currentFilePath = filePath
	e.SaveContents()
	return nil
}
