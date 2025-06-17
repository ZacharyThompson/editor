package main

type Editor struct {
	curs Cursor
	buf  []string
}

func NewEditor(buf *[]string) Editor {
	return Editor{NewCursor(0, 0), *buf}
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
