//go:build windows
// +build windows

package platforms

import (
	"syscall"
	"unsafe"
)

var (
	kernel32         = syscall.NewLazyDLL("kernel32.dll")
	screenBufferInfo = kernel32.NewProc("GetConsoleScreenBufferInfo")
)

type CONSOLE_SCREEN_BUFFER_INFO struct {
	DwSize              COORD
	DwCursorPosition    COORD
	WAttributes         uint16
	SrWindow            SMALL_RECT
	DwMaximumWindowSize COORD
}

type COORD struct {
	X, Y uint16
}

type SMALL_RECT struct {
	Left, Top, Right, Bottom uint16
}

func _GetTerminalInfo() error {
	var info CONSOLE_SCREEN_BUFFER_INFO
	rc, _, err := screenBufferInfo.Call(
		uintptr(syscall.Stdout),
		uintptr(unsafe.Pointer(&info)))

	if rc == 0 || err != nil {
		SetTerminalInfo(80, 60, 1, 1)

		return err
	}

	SetTerminalInfo(info.DwSize.X, info.DwSize.Y, info.DwCursorPosition.X, info.DwCursorPosition.Y)

	return nil
}
