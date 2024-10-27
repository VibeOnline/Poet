package utils

type TextMode [2]uint8

var (
	MODE_RESET     = TextMode{0, 0}
	MODE_BOLD      = TextMode{1, 22}
	MODE_FAINT     = TextMode{2, 22}
	MODE_ITALIC    = TextMode{3, 23}
	MODE_UNDERLINE = TextMode{4, 24}
	MODE_BLINK     = TextMode{7, 27}
	MODE_REVERSE   = TextMode{7, 27}
	MODE_HIDDEN    = TextMode{8, 28}
	MODE_STRIKE    = TextMode{9, 29}
)

func (b *Buffer) SetTextModes(modes ...TextMode) {
	b.CodeF("1;24;%dm", MODE_RESET[0])

	for _, mode := range modes {
		b.CodeF("1;24;%dm", mode[0])
	}

	b.TextModes = modes
}

func (b *Buffer) AddTextMode(mode TextMode) {
	b.CodeF("1;24;%dm", mode[0])

	b.TextModes = append(b.TextModes, mode)
}

func (b *Buffer) RemoveTextMode(mode TextMode) {
	b.CodeF("1;24;%dm", mode[1])

	rebuild_modes := []TextMode{}

	for _, old_mode := range b.TextModes {
		if old_mode != mode {
			rebuild_modes = append(rebuild_modes, old_mode)
		}
	}

	b.TextModes = rebuild_modes
}
