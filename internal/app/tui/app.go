package tui

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/term"
)

func main() {

	// TODO: Move SigInt - etc - to os.Signal calls.
	// syscall.sig

	restoreState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatalf("Cannot enter raw terminal mode because %v", err)
	}
	defer term.Restore(int(os.Stdin.Fd()), restoreState)

	var buf = [256]byte{}
	for {
		n, err := os.Stdin.Read(buf[:])
		if err != nil {
			log.Fatalf("Cannot read from standard-in because %v", err)
		}

		if n == 0 {
			continue
		}

		var keyCode, modifiers = readBuffer(buf[0:n])

		fmt.Println(keyCode, " > ", string(keyCode.GetValue()), " > ", modifiers, " > ", buf[0:n], string(buf[0:n]))
	}
}

func readBuffer(b []byte) (KeyCode, Modifiers) {
	var bufTxt = string(b)
	var isInputSequence = strings.HasPrefix(bufTxt, "\x1b[")
	var isVtSequence = strings.HasSuffix(bufTxt, "~")
	var keyTxt string
	var keyCode KeyCode
	var modifierTxt = "0"

	// Terminal input sequences.
	// See also: https://en.wikipedia.org/wiki/ANSI_escape_code#Terminal_input_sequences
	if isInputSequence {
		// Vt sequence.
		if isVtSequence {
			var elements = strings.Split(bufTxt[2:len(bufTxt)-1], ";")

			switch elements[0] {
			case "1":
				keyCode = KeyHome
			case "2":
				keyCode = KeyInsert
			case "3":
				keyCode = KeyDelete
			case "4":
				keyCode = KeyEnd
			case "5":
				keyCode = KeyPageUp
			case "6":
				keyCode = KeyPageDown
			case "7":
				keyCode = KeyHome
			case "8":
				keyCode = KeyEnd
			case "11":
				keyCode = KeyF1
			case "12":
				keyCode = KeyF2
			case "13":
				keyCode = KeyF3
			case "14":
				keyCode = KeyF4
			case "15":
				keyCode = KeyF5
			case "17":
				keyCode = KeyF6
			case "18":
				keyCode = KeyF7
			case "19":
				keyCode = KeyF8
			case "20":
				keyCode = KeyF9
			case "21":
				keyCode = KeyF10
			case "23":
				keyCode = KeyF11
			case "24":
				keyCode = KeyF12
			case "25":
				keyCode = KeyF13
			case "26":
				keyCode = KeyF14
			case "28":
				keyCode = KeyF15
			case "29":
				keyCode = KeyF16
			case "31":
				keyCode = KeyF17
			case "32":
				keyCode = KeyF18
			case "33":
				keyCode = KeyF19
			case "34":
				keyCode = KeyF20
			default:
				log.Fatalf("Cannot decode VT sequence %v", elements[0])
			}

			if len(elements) == 2 {
				modifierTxt = elements[1]
			}
		}

		// xTerm sequence.
		if !isVtSequence {
			var elements = strings.Split(bufTxt[2:], ";")
			var last = elements[len(elements)-1]

			switch len(last) {
			case 1:
				keyTxt = last
			case 2:
				keyTxt = string(last[1])
				modifierTxt = string(last[0])
			}

			switch keyTxt {
			case "A":
				keyCode = KeyUpArrow
			case "B":
				keyCode = KeyDownArrow
			case "C":
				keyCode = KeyRightArrow
			case "D":
				keyCode = KeyLeftArrow
			case "F":
				keyCode = KeyEnd
			case "G":
				keyCode = KeyNumPad5
			case "H":
				keyCode = KeyHome
			case "P":
				keyCode = KeyF1
			case "Q":
				keyCode = KeyF2
			case "R":
				keyCode = KeyF3
			case "S":
				keyCode = KeyF4
			default:
				log.Fatalf("Cannot decode VT sequence %v", elements)
			}
		}

		num, err := strconv.ParseInt(modifierTxt, 10, 64)
		if err != nil {
			log.Fatalf("Cannot read modifier key(s) due to unexpected value of %v", err)
		}
		var modifiers = Modifiers(num)

		return keyCode, modifiers
	}

	if !isInputSequence {
		keyTxt = bufTxt
	}

	// Special cases
	switch bufTxt {
	case "\x1bOP":
		return KeyF1, 0
	case "\x1bOQ":
		return KeyF2, 0
	case "\x1bOR":
		return KeyF3, 0
	case "\x1bOS":
		return KeyF4, 0
	}

	if slices.Equal(b, []byte{194, 163}) {
		return KeyPoundSterlingSymbol, 0
	}

	if len(keyTxt) != 1 {
		log.Fatalf("Unexpected key value of %v %v", keyTxt, b)
	}

	// ascii, err := strconv.ParseInt(string(keyTxt[0]), 10, 64)
	// if err != nil {
	// 	log.Fatalf("Cannot read key code because %v for value %v", err, b)
	// }

	var keyCode2, found = KeyCodes[int(keyTxt[0])]
	if !found {
		// We don't support all control characters (0 -> 32).
		if len(b) == 1 && b[0] <= 31 {
			return KeyNone, 0
		}

		log.Fatalf("Unexpected key value of %v %v", keyTxt, b)
	}

	return keyCode2, 0
}
