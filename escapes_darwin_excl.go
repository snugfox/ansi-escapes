// +build !darwin

package escapes

// ANSI escape sequences for saving and restoring the cursor position
const (
	CursorSavePosition    = Esc + "s"
	CursorRestorePosition = Esc + "u"
)
