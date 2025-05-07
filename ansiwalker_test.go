package ansiwalker

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ansiWalkerTestCase struct {
	ansiString  string
	nonAnsiText string
}

func consolidateNonAnsi(input string) string {
	nonAnsiBuf := strings.Builder{}
	idx := 0
	for idx < len(input) {
		r, _, next, _ := ANSIWalk(input, idx)
		nonAnsiBuf.WriteRune(r)
		idx = next
	}
	return nonAnsiBuf.String()
}

func TestCSI(t *testing.T) {
	tests := []ansiWalkerTestCase{
		// testing the parsing of CSI ANSI escape sequences
		{
			ansiString:  "\x1b[31mRed text\x1b[0m and normal",
			nonAnsiText: "Red text and normal",
		},
		{
			ansiString:  "\x1b[?25lHidden\x1b[?25h cursor visible",
			nonAnsiText: "Hidden cursor visible",
		},

		// testing the parsing of OSC ANSI escape sequences
		{
			ansiString:  "\x1b]0;My Title\x07Visible text",
			nonAnsiText: "Visible text",
		},
		{
			ansiString:  "\x1b]52;clipboard;SGVsbG8=\x07After OSC",
			nonAnsiText: "After OSC",
		},

		// testing the parsing of DCS ANSI escape sequences
		{
			ansiString:  "\x1bP1;set;value\x1b\\After DCS",
			nonAnsiText: "After DCS",
		},
		{
			ansiString:  "\x1bPurl=https://ex.com\x1b\\Link text",
			nonAnsiText: "Link text",
		},

		// testing the parsing of SOS ANSI escape sequences
		{
			ansiString:  "\x1bXDATA123\x1b\\End SOS",
			nonAnsiText: "End SOS",
		},
		{
			ansiString:  "\x1bXHelloWorld\x1b\\Done",
			nonAnsiText: "Done",
		},

		// testing the parsing of PM ANSI escape sequences
		{
			ansiString:  "\x1b^SecretMsg\x1b\\OK",
			nonAnsiText: "OK",
		},
		{
			ansiString:  "\x1b^ABC123\x1b\\Finish",
			nonAnsiText: "Finish",
		},

		// testing the parsing of APC ANSI escape sequences
		{
			ansiString:  "\x1b_AppCmd\x1b\\Go",
			nonAnsiText: "Go",
		},
		{
			ansiString:  "\x1b_CustomData\x1b\\Stop",
			nonAnsiText: "Stop",
		},

		// testing the parsing of C1 two-byte ANSI escape sequences
		{
			ansiString:  "\x1b@HelloAfter",
			nonAnsiText: "HelloAfter",
		},
		{
			ansiString:  "\x1bDNextLine",
			nonAnsiText: "NextLine",
		},
	}

	for idx, tt := range tests {
		t.Run(fmt.Sprintf("ANSI Walk Test %d", idx+1), func(t *testing.T) {
			consolidatedNonAnsi := consolidateNonAnsi(tt.ansiString)
			assert.Equal(t, tt.nonAnsiText, consolidatedNonAnsi)
		})
	}
}
