# ansi-scapes [![GoDoc](https://godoc.org/github.com/snugfox/ansi-escapes?status.svg)](https://godoc.org/github.com/snugfox/ansi-escapes)
ansi-scapes is a minimal Go library for ANSI escape sequences. It handles the
small differences in the Apple's Terminal app, as well as enable and disable
escape sequences on Windows 10 v1511 and later.

🚧 ansi-escapes is still under development, and is subject to change 🚧

## Features
- Constants for simple escape sequences
- Functions for sequences with parameters
- Correct functionality across terminal emulators
- Ability to enable/disable escape sequence processing on Windows v1511 and
later

## Installation
With the [Go Programming Language](https://golang.org/),
```console
$ go get -u github.com/snugfox/ansi-escapes
```

## Usage
Go's import mechanism does not allow package names to contain hyphens, so import
the package as `escapes`.
```go
import escapes "github.com/snugfox/ansi-escapes"
```

## Example
```go
package main

import (
  "bytes"
  "fmt"
  "os"

  escapes "github.com/snugfox/ansi-escapes"
)

func main() {
  // Enable support on Windows for this application. It is safe to include on
  // OSes other than Windows, as the functions will only return nil; thus
  // compiled out.
  escapes.EnableVirtualTerminal(escapes.Stdout)
  defer escapes.DisableVirtualTerminal(escapes.Stdout)

  // Erase the screen. Remember that fmt.Println would print the newline *after*
  // the escape sequence.
  fmt.Print(escapes.EraseScreen)

  // Move the cursor one column to the right
  fmt.Print(escapes.CursorForward)

  // Move the cursor to (1, 1)
  fmt.Print(escapes.CursorPos(1, 1))

  // Display a super secret image
  var buf bytes.Buffer
  file, _ := os.Open("meow.jpg")
  buf.ReadFrom(file)
  fmt.Print(escapes.Image(buf.Bytes()))
}
```

## License
MIT (c) Snug_Fox
