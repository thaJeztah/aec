package aec

import (
	"strings"
)

const esc = "\x1b["

// Reset resets SGR effect.
const Reset = "\x1b[0m"

// ANSI represents an ANSI escape code.
type ANSI string

var empty = ANSI("")

// With returns a new ANSI sequence composed of this and the provided ANSI codes.
func (a ANSI) With(codes ...ANSI) ANSI {
	return concat(append([]ANSI{a}, codes...))
}

// Apply wraps the given string with the ANSI sequence and a reset code.
func (a ANSI) Apply(s string) string {
	return string(a) + s + Reset
}

// String returns the ANSI escape code as a string.
func (a ANSI) String() string {
	return string(a)
}

// Apply wraps the given string with all provided ANSI sequences.
func Apply(s string, codes ...ANSI) string {
	if len(codes) == 0 {
		return s
	}
	return concat(codes).Apply(s)
}

// concat combines multiple ANSI codes into a single ANSI sequence.
func concat(codes []ANSI) ANSI {
	if len(codes) == 1 {
		return codes[0]
	}
	var b strings.Builder
	for _, c := range codes {
		b.WriteString(string(c))
	}
	return ANSI(b.String())
}
