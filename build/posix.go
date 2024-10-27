//go:build !windows && !solaris
// +build !windows,!solaris

package platforms

import (
	"os/exec"
	"syscall"
	"unsafe"
)

func _GetTerminalInfo() error {
	ws := struct {
		Row uint16
		Col uint16
		X   uint16
		Y   uint16
	}{}

	_, _, err := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdout),
		uintptr(TIOCGWINSZ),
		uintptr(unsafe.Pointer(&ws)))

	if TIOCGWINSZ == 0 || err != 0 {
		SetTerminalInfo(80, 60, 1, 1)

		return err
	}

	SetTerminalInfo(ws.Col, ws.Row, ws.X, ws.Y)

	return nil
}

func init() {
	exec.Command("stty", "-f", "/dev/tty", "-raw", "echo").Run()
	defer exec.Command("stty", "-f", "/dev/tty", "raw", "-echo").Run()
}
