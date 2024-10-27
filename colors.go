package utils

import (
	"fmt"
	"math"
)

type Color uint8

const (
	COLOR_BLACK      = Color(0)
	COLOR_RED        = Color(1)
	COLOR_GREEN      = Color(2)
	COLOR_ORANGE     = Color(3)
	COLOR_BLUE       = Color(4)
	COLOR_PURPLE     = Color(5)
	COLOR_CYAN       = Color(6)
	COLOR_LIGHT_GRAY = Color(7)
	COLOR_GRAY       = Color(8)
	COLOR_PINK       = Color(9)
	COLOR_LIME       = Color(10)
	COLOR_YELLOW     = Color(11)
	COLOR_SKY_BLUE   = Color(12)
	COLOR_VIOLET     = Color(13)
	COLOR_AQUA       = Color(14)
	COLOR_WHITE      = Color(15)
)

// Color operations
func GetColorFromRGB(r uint8, g uint8, b uint8) Color {
	if r == g && g == b {
		if r < 8 {
			return 16
		}

		if r > 248 {
			return 231
		}

		return Color(math.Round((float64(r-8)/247)*24) + 232)
	}

	return Color(16 + (36 * math.Round(float64(r)/255*5)) + (6 * math.Round(float64(g)/255*5)) + math.Round(float64(b)/255*5))
}

var _HexColorList map[byte]Color

func GetColorFromHex(hex byte) Color {
	return _HexColorList[hex]
}

// Color management
func (b *Buffer) SetBackgroundColor(color Color) {
	b.BgColor = color

	b.CodeF("48;5;%dm", color)
}

func (b *Buffer) GetBackgroundColor() Color {
	return b.BgColor
}

func (b *Buffer) SetTextColor(color Color) {
	b.TextColor = color

	b.CodeF("38;5;%dm", color)
}

func (b *Buffer) GetTextColor() Color {
	return b.TextColor
}

// Draw with color
func (b *Buffer) Blit(text string, text_color string, bg_color string) {
	if len(text) == len(text_color) && len(text_color) == len(bg_color) {
		str := ""

		for i, c := range text {
			str += fmt.Sprintf("\x1b[48;5;%dm\x1b[38;5;%dm%c", GetColorFromHex(bg_color[i]), GetColorFromHex(text_color[i]), c)
		}

		b.WriteF("%s\x1b[48;5;%dm\x1b[38;5;%dm", str, b.GetBackgroundColor(), b.GetTextColor())
	}
}

func init() {
	hexStr := []byte("0123456789abcdef")

	_HexColorList = map[byte]Color{}

	for i, b := range hexStr {
		_HexColorList[b] = Color(i)
	}
}
