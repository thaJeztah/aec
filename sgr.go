package aec

import "strconv"

// RGB3Bit is a 3bit RGB color.
type RGB3Bit uint8

// RGB8Bit is a 8bit RGB color.
type RGB8Bit uint8

// NewRGB3Bit create a RGB3Bit from given RGB.
func NewRGB3Bit(r, g, b uint8) RGB3Bit {
	return RGB3Bit((r >> 7) | ((g >> 6) & 0x2) | ((b >> 5) & 0x4))
}

// NewRGB8Bit create a RGB8Bit from given RGB.
func NewRGB8Bit(r, g, b uint8) RGB8Bit {
	return RGB8Bit(16 + 36*(r/43) + 6*(g/43) + b/43)
}

// Color3BitF set the foreground color of text.
func Color3BitF(c RGB3Bit) ANSI {
	return ANSI(esc + itoaUint(uint(c+30)) + "m")
}

// Color3BitB set the background color of text.
func Color3BitB(c RGB3Bit) ANSI {
	return ANSI(esc + itoaUint(uint(c+40)) + "m")
}

// Color8BitF set the foreground color of text.
func Color8BitF(c RGB8Bit) ANSI {
	return ANSI(esc + "38;5;" + itoaUint(uint(c)) + "m")
}

// Color8BitB set the background color of text.
func Color8BitB(c RGB8Bit) ANSI {
	return ANSI(esc + "48;5;" + itoaUint(uint(c)) + "m")
}

// FullColorF set the foreground color of text.
func FullColorF(r, g, b uint8) ANSI {
	return ANSI(esc + "38;2;" +
		strconv.Itoa(int(r)) + ";" +
		strconv.Itoa(int(g)) + ";" +
		strconv.Itoa(int(b)) + "m")
}

// FullColorB set the foreground color of text.
func FullColorB(r, g, b uint8) ANSI {
	return ANSI(esc + "48;2;" +
		strconv.Itoa(int(r)) + ";" +
		strconv.Itoa(int(g)) + ";" +
		strconv.Itoa(int(b)) + "m")
}

// Style
const (
	// Bold set the text style to bold or increased intensity.
	Bold ANSI = "\x1b[1m"

	// Faint set the text style to faint.
	Faint ANSI = "\x1b[2m"

	// Italic set the text style to italic.
	Italic ANSI = "\x1b[3m"

	// Underline set the text style to underline.
	Underline ANSI = "\x1b[4m"

	// BlinkSlow set the text style to slow blink.
	BlinkSlow ANSI = "\x1b[5m"

	// BlinkRapid set the text style to rapid blink.
	BlinkRapid ANSI = "\x1b[6m"

	// Inverse swap the foreground color and background color.
	Inverse ANSI = "\x1b[7m"

	// Conceal set the text style to conceal.
	Conceal ANSI = "\x1b[8m"

	// CrossOut set the text style to crossed out.
	CrossOut ANSI = "\x1b[9m"

	// Frame set the text style to framed.
	Frame ANSI = "\x1b[51m"

	// Encircle set the text style to encircled.
	Encircle ANSI = "\x1b[52m"

	// Overline set the text style to overlined.
	Overline ANSI = "\x1b[53m"
)

// Foreground color of text.
const (
	// DefaultF is the default color of foreground.
	DefaultF ANSI = "\x1b[39m"

	// Normal color

	BlackF   ANSI = "\x1b[30m"
	RedF     ANSI = "\x1b[31m"
	GreenF   ANSI = "\x1b[32m"
	YellowF  ANSI = "\x1b[33m"
	BlueF    ANSI = "\x1b[34m"
	MagentaF ANSI = "\x1b[35m"
	CyanF    ANSI = "\x1b[36m"
	WhiteF   ANSI = "\x1b[37m"

	// Light color

	LightBlackF   ANSI = "\x1b[90m"
	LightRedF     ANSI = "\x1b[91m"
	LightGreenF   ANSI = "\x1b[92m"
	LightYellowF  ANSI = "\x1b[93m"
	LightBlueF    ANSI = "\x1b[94m"
	LightMagentaF ANSI = "\x1b[95m"
	LightCyanF    ANSI = "\x1b[96m"
	LightWhiteF   ANSI = "\x1b[97m"
)

// Background color of text.
const (
	// DefaultB is the default color of background.
	DefaultB ANSI = "\x1b[49m"

	// Normal color

	BlackB   ANSI = "\x1b[40m"
	RedB     ANSI = "\x1b[41m"
	GreenB   ANSI = "\x1b[42m"
	YellowB  ANSI = "\x1b[43m"
	BlueB    ANSI = "\x1b[44m"
	MagentaB ANSI = "\x1b[45m"
	CyanB    ANSI = "\x1b[46m"
	WhiteB   ANSI = "\x1b[47m"

	// Light color

	LightBlackB   ANSI = "\x1b[100m"
	LightRedB     ANSI = "\x1b[101m"
	LightGreenB   ANSI = "\x1b[102m"
	LightYellowB  ANSI = "\x1b[103m"
	LightBlueB    ANSI = "\x1b[104m"
	LightMagentaB ANSI = "\x1b[105m"
	LightCyanB    ANSI = "\x1b[106m"
	LightWhiteB   ANSI = "\x1b[107m"
)
