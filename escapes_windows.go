// +build windows

package escapes

import (
	"syscall"

	"golang.org/x/sys/windows"
)

// EnableVirtualTerminal enables virtual terminal escapes sequences for a
// console handle. It is only effective when built for Windows. On other OSes,
// it will simply return nil.
func EnableVirtualTerminal(fd uintptr) error {
	var mode uint32
	if err := syscall.GetConsoleMode(syscall.Handle(fd), &mode); err != nil {
		return err
	}
	mode |= windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING

	if err := windows.SetConsoleMode(windows.Handle(fd), mode); err != nil {
		return err
	}
	return nil
}

// DisableVirtualTerminal disables virtual terminal escapes sequences for a
// console handle. It is only effective when built for Windows. On other OSes,
// it will simply return nil.
func DisableVirtualTerminal(fd uintptr) error {
	var mode uint32
	if err := syscall.GetConsoleMode(syscall.Handle(fd), &mode); err != nil {
		return err
	}
	mode ^= windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING

	if err := windows.SetConsoleMode(windows.Handle(fd), mode); err != nil {
		return err
	}
	return nil
}

// GetConsoleSize gets the dimensions of the console.
func GetConsoleSize(fd uintptr) (*ConsoleDim, error) {
	var info windows.ConsoleScreenBufferInfo
	if err := windows.GetConsoleScreenBufferInfo(windows.Handle(fd), &info); err != nil {
		return nil, err
	}

	// Return only the console dimensions
	return &ConsoleDim{
		Rows: int(info.Window.Bottom - info.Window.Top),
		Cols: int(info.Window.Right - info.Window.Left),
	}, nil
}
