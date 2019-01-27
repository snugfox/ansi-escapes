// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris

package escapes

import (
	"golang.org/x/sys/unix"
)

func GetConsoleSize(fd uintptr) (*ConsoleDim, error) {
	ws, err := unix.IoctlGetWinsize(int(fd), unix.TIOCGWINSZ)
	if err != nil {
		return nil, err
	}

	// Unpack the row and column dimensions from the C struct
	return &ConsoleDim{
		Rows: int(ws.Row),
		Cols: int(ws.Col),
	}, nil
}
