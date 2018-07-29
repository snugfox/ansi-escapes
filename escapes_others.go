// +build !darwin

package escapes

const (
	CursorSavePosition    = ESC + "s"
	CursorRestorePosition = ESC + "u"
)
