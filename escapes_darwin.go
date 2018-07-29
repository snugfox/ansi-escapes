package escapes

import "os"

var (
	CursorSavePosition, CursorRestorePosition string
)

func init() {
	if os.Getenv("TERM_PROGRAM") == "Apple_Terminal" {
		CursorSavePosition = ESC + "7"
		CursorRestorePosition = ESC + "8"
	}
}
