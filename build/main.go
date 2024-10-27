package platforms

type TermInfo struct {
	Size      [2]uint16
	CursorPos [2]uint16
}

var _GlobalInfo TermInfo

func GetTerminalInfo() (TermInfo, error) {
	err := _GetTerminalInfo()

	return _GlobalInfo, err
}

func SetTerminalInfo(width uint16, height uint16, cursor_x uint16, cursor_y uint16) {
	_GlobalInfo = TermInfo{
		Size:      [2]uint16{width, height},
		CursorPos: [2]uint16{cursor_x, cursor_y},
	}
}

func GetSize() (uint16, uint16) {
	return _GlobalInfo.Size[0], _GlobalInfo.Size[1]
}

func GetCursorPos() (uint16, uint16) {
	GetTerminalInfo()

	return _GlobalInfo.CursorPos[0], _GlobalInfo.CursorPos[1]
}

func init() {
	GetTerminalInfo()
}
