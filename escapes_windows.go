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

	if err = setConsoleMode(handle, mode); err != nil {
		return errors.New("Failed to set console mode: " + err.Error())
	}

	return nil
}
