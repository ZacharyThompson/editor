package main

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
