// +build !darwin

package escapes

// We assume that only the Apple Terminal uses 7/8 instead of s/u

// ANSI escape sequences for saving and restoring the cursor position
const (
	CursorSavePosition    = Esc + "s"
	CursorRestorePosition = Esc + "u"
)
