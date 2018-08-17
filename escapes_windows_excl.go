// +build !windows

package escapes

// EnableVirtualTerminal enables virtual terminal escapes sequences for a
// console handle. It is only effective when built for Windows. On other OSes,
// it will simply return nil.
func EnableVirtualTerminal(console ConsoleHandle) error {
	return nil
}

// DisableVirtualTerminal disables virtual terminal escapes sequences for a
// console handle. It is only effective when built for Windows. On other OSes,
// it will simply return nil.
func DisableVirtualTerminal(console ConsoleHandle) error {
	return nil
}
