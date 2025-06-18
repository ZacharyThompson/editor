package main

type Cursor struct {
	line uint
	col  uint
}

func NewCursor(line, col uint) Cursor {
	return Cursor{line: line, col: col}
}
func (c *Cursor) MoveLeft(count uint) {
	c.col = max(0, c.col-count)
}
func (c *Cursor) MoveRight(count, maxCol uint) {
	c.col = min(c.col+count, maxCol)
}
func (c *Cursor) MoveUp(count uint) {
	c.line = max(0, c.line-count)
}
func (c *Cursor) MoveDown(count uint, maxLine uint) {
	c.line = min(c.line+count, maxLine)
}
func (c *Cursor) GetPos() (line, col uint) {
	return c.line, c.col
}
func (c *Cursor) MoveTo(line, col uint) {
	c.line = line
	c.col = col
}
