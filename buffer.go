package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"slices"
	"strings"
)

type Buffer struct {
	currentFilePath string
	contents        []string
	curs            Cursor
}

func NewEmptyBuffer() Buffer {
	return Buffer{"", []string{""}, NewCursor(0, 0)}
}

func NewBufferFromPath(filePath string) (Buffer, error) {
	if filePath == "" {
		return NewEmptyBuffer(), errors.New("Cannot create new Buffer: Empty filepath")
	}
	if !fs.ValidPath(filePath) {
		return NewEmptyBuffer(), errors.New("Cannot create new Buffer: Invalid filepath")
	}
	data, err := os.ReadFile(filePath)
	if err != nil {
		return NewEmptyBuffer(), fmt.Errorf("Cannot create new Buffer: %v", err)
	}
	lines := strings.Split(string(data), "\n")
	return Buffer{"", lines, NewCursor(0, 0)}, nil
}

func (b *Buffer) DeleteCharsBeforeCursor(count uint) {
	if b.curs.line == 0 && b.curs.col == 0 {
		return
	} else if b.curs.col == 0 {
		// delete newline
		old := int(b.curs.line)
		new := old - 1
		b.contents[old-1] += b.contents[old]
		b.curs.MoveTo(uint(old-1), uint(len(b.contents[new])))
		b.contents = slices.Delete(b.contents, old, old+1)
	} else {
		line := b.contents[b.curs.line]
		line = line[0:max(0, b.curs.col-count)] + line[b.curs.col:]
		b.contents[b.curs.line] = line
		b.curs.MoveLeft(1)
	}
}

func (b *Buffer) SaveContents() error {
	if b.currentFilePath == "" {
		return errors.New("Cannot save contents: No filepath given")
	}
	// TODO: maybe check fs.ValidPath?
	file, err := os.Create(b.currentFilePath)
	if err != nil {
		return fmt.Errorf("Cannot save contents: Cannot create file: %v", err)
	}
	file.WriteString(strings.Join(b.contents, "\n") + "\n")
	return nil
}

func (b *Buffer) SaveContentsToPath(filePath string) error {
	if filePath == "" {
		return errors.New("Cannot save contents to path: Empty filepath")
	}
	if !fs.ValidPath(filePath) {
		return errors.New("Cannot save contents to path: Invalid filepath")
	}
	b.currentFilePath = filePath
	b.SaveContents()
	return nil
}
