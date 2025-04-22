package aec

import "strconv"

// EraseMode is listed in a variable EraseModes.
type EraseMode uint

var (
	// EraseModes is a list of EraseMode.
	EraseModes = struct {
		// All erase all.
		All EraseMode

		// Head erase to head.
		Head EraseMode

		// Tail erase to tail.
		Tail EraseMode
	}{
		Tail: 0,
		Head: 1,
		All:  2,
	}
)

const (
	// Save saves the cursor position.
	Save ANSI = "\x1b[s"

	// Restore restores the cursor position.
	Restore ANSI = "\x1b[u"

	// Hide hides the cursor.
	Hide ANSI = "\x1b[?25l"

	// Show shows the cursor.
	Show ANSI = "\x1b[?25h"

	// Report reports the cursor position.
	Report ANSI = "\x1b[6n"
)

// itoaUint converts a uint to a string.
func itoaUint(n uint) string {
	return strconv.FormatUint(uint64(n), 10)
}

// Up moves up the cursor.
func Up(n uint) ANSI {
	if n == 0 {
		return empty
	}
	return ANSI(esc + itoaUint(n) + "A")
}

// Down moves down the cursor.
func Down(n uint) ANSI {
	if n == 0 {
		return empty
	}
	return ANSI(esc + itoaUint(n) + "B")
}

// Right moves right the cursor.
func Right(n uint) ANSI {
	if n == 0 {
		return empty
	}
	return ANSI(esc + itoaUint(n) + "C")
}

// Left moves left the cursor.
func Left(n uint) ANSI {
	if n == 0 {
		return empty
	}
	return ANSI(esc + itoaUint(n) + "D")
}

// NextLine moves down the cursor to head of a line.
func NextLine(n uint) ANSI {
	if n == 0 {
		return empty
	}
	return ANSI(esc + itoaUint(n) + "E")
}

// PreviousLine moves up the cursor to head of a line.
func PreviousLine(n uint) ANSI {
	if n == 0 {
		return empty
	}
	return ANSI(esc + itoaUint(n) + "F")
}

// Column set the cursor position to a given column.
func Column(col uint) ANSI {
	return ANSI(esc + itoaUint(col) + "G")
}

// Position set the cursor position to a given absolute position.
func Position(row, col uint) ANSI {
	return ANSI(esc + itoaUint(row) + ";" + itoaUint(col) + "H")
}

// EraseDisplay erases display by given EraseMode.
func EraseDisplay(m EraseMode) ANSI {
	return ANSI(esc + strconv.Itoa(int(m)) + "J")
}

// EraseLine erases lines by given EraseMode.
func EraseLine(m EraseMode) ANSI {
	return ANSI(esc + strconv.Itoa(int(m)) + "K")
}

// ScrollUp scrolls up the page.
func ScrollUp(n int) ANSI {
	if n == 0 {
		return empty
	}
	return ANSI(esc + strconv.Itoa(n) + "S")
}

// ScrollDown scrolls down the page.
func ScrollDown(n int) ANSI {
	if n == 0 {
		return empty
	}
	return ANSI(esc + strconv.Itoa(n) + "T")
}
