//go:build solaris
// +build solaris

package platforms

import (
	"syscall"

	"golang.org/x/sys/unix"
)

const TIOCGWINSZ = 21608

func _GetTerminalInfo() error {
	var ws *unix.Winsize
	ws, err := unix.IoctlGetWinsize(syscall.Stdout, TIOCGWINSZ)

	if err != nil {
		SetTerminalInfo(80, 60, 1, 1)

		return err
	}

	SetTerminalInfo(ws.Col, ws.Row, ws.Xpixel, ws.Ypixel)

	return nil
}
