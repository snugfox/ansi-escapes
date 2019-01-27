// +build !windows

package escapes

// We make the assumption that non-Windows OSes support escape sequences by
// default. Thus, EnableVirtualTerminal and DisableVirtualTerminal are defined
// as no-ops.

// EnableVirtualTerminal enables virtual terminal escapes sequences for a
// console handle. It is only effective when built for Windows. On other OSes,
// it will simply return nil.
func EnableVirtualTerminal(fd uintptr) error {
	return nil
}

// DisableVirtualTerminal disables virtual terminal escapes sequences for a
// console handle. It is only effective when built for Windows. On other OSes,
// it will simply return nil.
func DisableVirtualTerminal(fd uintptr) error {
	return nil
}
