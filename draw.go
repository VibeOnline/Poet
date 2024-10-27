package utils

import (
	"fmt"

	platforms "utils/build"
)

type Buffer struct {
	Data      []byte
	Cursor    [2]uint16
	Size      [2]uint16
	Offset    [2]uint16
	TextColor Color
	BgColor   Color
	TextModes []TextMode
	Buffering bool
}

// Buffer writing
func (b *Buffer) Draw() {
	fmt.Print(string(b.Data))

	b.Data = []byte{}
}

func (b *Buffer) Write(text string) {
	(*b).Data = append((*b).Data, []byte(text)...)
}

func (b *Buffer) WriteBytes(text []byte) {
	(*b).Data = append((*b).Data, text...)
}

func (b *Buffer) WriteF(text string, data ...any) {
	b.Write(fmt.Sprintf(text, data...))
}

func (b *Buffer) Code(code string) {
	b.WriteF("\x1b[%s", code)
}

func (b *Buffer) CodeF(code string, data ...any) {
	b.Code(fmt.Sprintf(code, data...))
}

// Cursor management
func (b *Buffer) SetCursorPos(x uint16, y uint16) {
	b.CodeF("%d;%dH", y, x)
}

func (b *Buffer) GetCursorPos() (x uint16, y uint16) {
	platforms.GetCursorPos()

	return b.Cursor[0], b.Cursor[1]
}

func (b *Buffer) Scroll() {
	b.Code("1E")
}

func (b *Buffer) ScrollBy(y int16) {
	if y != 0 {
		if y < 0 {
			b.CodeF("%dF", -y)
		} else {
			b.CodeF("%dE", y)
		}
	}
}

func (b *Buffer) MoveCursorBy(x int16, y int16) {
	if y != 0 {
		if y < 0 {
			b.CodeF("%dA", -y)
		} else {
			b.CodeF("%dB", y)
		}
	}

	if x != 0 {
		if x < 0 {
			b.CodeF("%dD", -x)
		} else {
			b.CodeF("%dC", x)
		}
	}
}

func (b *Buffer) HideCursor() {
	b.Code("?25l")
}

func (b *Buffer) ShowCursor() {
	b.Code("?25h")
}

func (b *Buffer) SaveCursor() {
	b.Code("s")
}

func (b *Buffer) LoadCursor() {
	b.Code("u")
}

func (b *Buffer) ClearFromCursor() {
	b.Code("0J")
}

func (b *Buffer) ClearToCursor() {
	b.Code("1J")
}

func (b *Buffer) ClearLineFromCursor() {
	b.Code("0K")
}

func (b *Buffer) ClearLineToCursor() {
	b.Code("1K")
}

// Window management
func (b *Buffer) GetSize() (uint16, uint16) {
	b.Size[0], b.Size[1] = platforms.GetSize()

	return b.Size[0] - b.Offset[0], b.Size[1] - b.Offset[1]
}

func (b *Buffer) SetSize(width uint16, height uint16) {
	b.Size[0] = width
	b.Size[1] = height
}

func (b *Buffer) GetOffset() (uint16, uint16) {
	return b.Offset[0], b.Offset[1]
}

func (b *Buffer) SetOffset(x uint16, y uint16) {
	b.Offset[0] = x
	b.Offset[1] = y
}

func (b *Buffer) Clear() {
	b.Code("2J")
}

func (b *Buffer) ClearLine() {
	b.Code("2K")
}

func (b *Buffer) ClearHistory() {
	b.Code("3J")
}

func (b *Buffer) Save() {
	b.Code("?47h")
}

func (b *Buffer) Load() {
	b.Code("?47h")
}

func (b *Buffer) SetBuffer(enabled bool) {
	b.Buffering = enabled

	if enabled {
		b.Code("?1049h")
	} else {
		b.Code("?1049l")
	}
}
