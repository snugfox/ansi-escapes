package escapes

import (
	"encoding/base64"
	"strconv"
)

type ConsoleHandle int

const (
	Stdin  ConsoleHandle = -10
	Stdout ConsoleHandle = -11
	Stderr ConsoleHandle = -12
)

const (
	ESC = "\u001B["
	OSC = "\u001B]"
	BEL = "\u0007"

	CursorUp       = ESC + "A"
	CursorDown     = ESC + "B"
	CursorForward  = ESC + "C"
	CursorBackward = ESC + "D"
	CursorNextLine = ESC + "E"
	CursorPrevLine = ESC + "F"
	CursorLeft     = ESC + "G"
	CursorTop      = ESC + "d"
	CursorTopLeft  = ESC + "H"

	CursorBlinkEnable  = ESC + "?12h"
	CursorBlinkDisable = ESC + "?12I"
	CursorShow         = ESC + "?25h"
	CursorHide         = ESC + "?25l"

	ScrollUp   = ESC + "S"
	ScrollDown = ESC + "T"

	TextInsertChar = ESC + "@"
	TextDeleteChar = ESC + "P"
	TextEraseChar  = ESC + "X"
	TextInsertLine = ESC + "L"
	TextDeleteLine = ESC + "M"

	EraseRight  = ESC + "K"
	EraseLeft   = ESC + "1K"
	EraseLine   = ESC + "2K"
	EraseDown   = ESC + "J"
	EraseUp     = ESC + "1J"
	EraseScreen = ESC + "2J"

	ClearScreen = "\u001Bc"
)

func CursorPosX(x int) string {
	return ESC + strconv.Itoa(x+1) + "G"
}

func CursorPosY(y int) string {
	return ESC + strconv.Itoa(y+1) + "d"
}

func CursorPos(x, y int) string {
	return ESC + strconv.Itoa(y+1) + ";" + strconv.Itoa(x+1) + "H"
}

func CursorMove(x, y int) string {
	var s string
	if x < 0 {
		s += ESC + strconv.Itoa(-x) + "D"
	} else if x > 0 {
		s += ESC + strconv.Itoa(x) + "C"
	}
	if y < 0 {
		s += ESC + strconv.Itoa(-y) + "A"
	} else if y > 0 {
		s += ESC + strconv.Itoa(y) + "B"
	}
	return s
}

func Scoll(n int) string {
	if n > 0 {
		return ESC + strconv.Itoa(n) + "S"
	} else if n < 0 {
		return ESC + strconv.Itoa(-n) + "T"
	} else {
		return ""
	}
}

func TextInsertChars(n int) string {
	return ESC + strconv.Itoa(n) + "@"
}

func TextDeleteChars(n int) string {
	return ESC + strconv.Itoa(n) + "P"
}

func TextEraseChars(n int) string {
	return ESC + strconv.Itoa(n) + "X"
}

func TextInsertLines(n int) string {
	return ESC + strconv.Itoa(n) + "L"
}

func TextDeleteLines(n int) string {
	return ESC + strconv.Itoa(n) + "M"
}

func Link(url, text string) string {
	return OSC + "8;;" + url + BEL + text + OSC + "8;;" + BEL
}

func Image(img []byte) string {
	return ImageWidthHeight(img, 0, 0, true)
}

func ImageWidthHeight(img []byte, height, width int, preserveAspectRatio bool) string {
	s := OSC + "1337;File=inline=1"
	if height > 0 {
		s += ";width=" + strconv.Itoa(height)
	}
	if width > 0 {
		s += ";height=" + strconv.Itoa(width)
	}
	if !preserveAspectRatio {
		s += ";preserveAspectRatio=0"
	}

	return s + ":" + base64.StdEncoding.EncodeToString(img) + BEL
}

func SetCwd(dir string) string {
	return OSC + "50;CurrentDir=" + dir + BEL
}
