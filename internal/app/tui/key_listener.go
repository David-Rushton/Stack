package tui

import (
	"bufio"
	"errors"
	"io"
	"os"
)

func readKey(a *app) KeyEvent {

	reader := bufio.NewReader(os.Stdin)

	var keysRead = 0
	var buf = [256]byte{}
	for {

		n, err := reader.Read(buf[:])
		if err != nil {
			if !errors.Is(err, io.EOF) {
				a.Log.Errorf("Cannot read rune because %v", err)
				return KeyEvent{KeyEndOfText, 0, 0}
			}

		}

		if n == 0 {
			return KeyEvent{KeyNone, 0, 0}
		}

		// Shutdown
		if n == 1 && buf[0] == 3 {
			a.Log.Information("Ctrl-C pressed")
			a.stop()
			return KeyEvent{KeyEndOfText, 0, 0}
		}

		// Terminal input sequences
		// https://en.wikipedia.org/wiki/ANSI_escape_code#Terminal_input_sequences
		if buf[0] == byte(KeyEscape) && n > 1 {
			// <esc> <char> || Alt keypress or keycode sequence.
			if n == 2 {
				return KeyEvent{KeyEscape, rune(buf[1]), 0}
			}

			// <esc> '['
			if buf[1] == byte('[') {
				// <esc> '[' <nochar> || Alt
				if n == 2 {
					return KeyEvent{KeyNone, 0, ModifierAlt}
				}

				// <esc> '[' (<keycode>) (';'<modifier>) '~' || Vt sequence
				if buf[n-1] == byte('~') {

				}

				// <esc> '[' (<modifier>) <char> || xTerm sequence
			}

		}

		a.Log.Informationln(keysRead, ": ", string(buf[:]), " >> ", buf)

		keysRead++
	}

}
