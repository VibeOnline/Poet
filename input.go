package utils

import (
	"bytes"
	"os"
)

var _InputFocus *[]byte

// Capturing text input
func ReadCapture(focus *[]byte) {
	_InputFocus = focus
}

func ReadRelease() {
	_InputFocus = nil
}

func Read(output *[]byte) {
	if _InputFocus == nil {
		_InputFocus = output
	}
}

func ReadUntil(output *[]byte, end []byte) bool {
	Read(output)

	found := bytes.HasSuffix(*output, end)
	if found && _InputFocus == output {
		ReadRelease()
	}

	return found
}

func init() {
	go func() {
		for {
			b := make([]byte, 1)

			os.Stdin.Read(b)

			if b[0] == byte(3) {
				os.Exit(0)
			}

			if _InputFocus != nil {
				switch b[0] {
				case byte(127):
					if len(*_InputFocus) > 0 {
						*_InputFocus = (*_InputFocus)[0 : len(*_InputFocus)-1]
					}
				case byte(13):
					*_InputFocus = append(*_InputFocus, '\n')
				default:
					*_InputFocus = append(*_InputFocus, b[0])
				}
			}
		}
	}()
}
