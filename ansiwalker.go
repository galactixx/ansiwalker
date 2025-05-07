package ansiwalker

import (
	"unicode/utf8"
)

// returnRune decodes and returns the next visible rune from
// the string starting at index i, along with its byte size, the
// index of the next position, and a boolean indicating if the
// end of the string has been reached.
func returnRune(s string, i int) (r rune, size int, next int, ok bool) {
	if i >= len(s) {
		return rune(0), 0, -1, false
	}
	r, rSize := utf8.DecodeRuneInString(s[i:])
	nextI := i + rSize
	notEOF := nextI < len(s)
	return r, rSize, nextI, notEOF
}

// ANSIWalk scans the input string s from index i, skipping any
// ANSI/VT100 escape sequences, and returns the next visible
// rune, its byte size, the index to resume from, and a boolean
// indicating if the end of the string has been reached.
func ANSIWalk(s string, i int) (r rune, size int, next int, ok bool) {
	// If it’s not ESC (0x1B), decode & emit the rune
	if s[i] != 0x1B {
		return returnRune(s, i)
	}

	// We saw ESC—now decide which ANSI family it is
	if i+1 >= len(s) {
		return returnRune(s, i+1)
	}

	switch s[i+1] {
	case '[':
		// ───── CSI (Control Sequence Introducer) ─────
		// Skip from ESC [ up to and including the final byte
		// 0x40–0x7E
		j := i + 2
		for j < len(s) && (s[j] < 0x40 || s[j] > 0x7E) {
			j++
		}
		return returnRune(s, j+1)

	case ']':
		// ───── OSC (Operating System Command) ─────
		// Skip until BEL (0x07) or ESC '\' terminator
		j := i + 2
		for j < len(s) {
			if s[j] == 0x07 {
				return returnRune(s, j+1)
			}
			if s[j] == 0x1B && j+1 < len(s) && s[j+1] == '\\' {
				return returnRune(s, j+2)
			}
			j++
		}
		return returnRune(s, j)

	case 'P', '_', '^', 'X':
		// ─── DCS / APC / PM / SOS ───
		// These all end on the two‑byte ESC '\' sequence
		j := i + 2
		for j+1 < len(s) {
			if s[j] == 0x1B && s[j+1] == '\\' {
				return returnRune(s, j+2)
			}
			j++
		}
		return returnRune(s, j)

	default:
		// ─── Other two‑byte controls (C1) ───
		// e.g. ESC c, ESC E, etc.: just skip 2 bytes
		return returnRune(s, i+2)
	}
}
