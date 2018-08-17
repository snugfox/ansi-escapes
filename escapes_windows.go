// +build windows

package escapes

import (
	"errors"
	"syscall"
)

const enableVTProcessing uint32 = 0x0004

var (
	kernel32           = syscall.NewLazyDLL("kernel32.dll")
	procSetConsoleMode = kernel32.NewProc("SetConsoleMode")
)

// Enable virtual terminal processing using the Windows SetConsoleMode syscall
// https://docs.microsoft.com/en-us/windows/console/setconsolemode
func setConsoleMode(console syscall.Handle, mode uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procSetConsoleMode.Addr(), 2, uintptr(console), uintptr(mode), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = e1
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

// EnableVirtualTerminal enables virtual terminal escapes sequences for a
// console handle. It is only effective when built for Windows. On other OSes,
// it will simply return nil.
func EnableVirtualTerminal(console ConsoleHandle) error {
	handle, err := syscall.GetStdHandle(int(console))
	if err != nil {
		return errors.New("Failed to get console handle: " + err.Error())
	}

	var mode uint32
	if err = syscall.GetConsoleMode(handle, &mode); err != nil {
		return errors.New("Failed to get console mode: " + err.Error())
	}
	mode |= enableVTProcessing

	if err = setConsoleMode(handle, mode); err != nil {
		return errors.New("Failed to set console mode: " + err.Error())
	}

	return nil
}

// DisableVirtualTerminal disables virtual terminal escapes sequences for a
// console handle. It is only effective when built for Windows. On other OSes,
// it will simply return nil.
func DisableVirtualTerminal(console ConsoleHandle) error {
	handle, err := syscall.GetStdHandle(int(console))
	if err != nil {
		return errors.New("Failed to get console handle: " + err.Error())
	}

	var mode uint32
	if err = syscall.GetConsoleMode(handle, &mode); err != nil {
		return errors.New("Failed to get console mode: " + err.Error())
	}
	mode ^= enableVTProcessing

	if err = setConsoleMode(handle, mode); err != nil {
		return errors.New("Failed to set console mode: " + err.Error())
	}

	return nil
}
