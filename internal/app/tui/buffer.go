package tui

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/term"
)

type buffer struct {
	top      int
	left     int
	width    int
	height   int
	position int
	values   []rune
}

func (b *buffer) isEndOfLine() bool {
	return b.position%b.width == 0
}

func (b *buffer) isTabStop() bool {
	// FIX: This won't work for sub buffers that start at offsets other than 4.
	return b.position%4 == 0
}

func (b *buffer) write(char rune) {

	// TODO: Rune literals here?  Or constant values?  Or both?  Not comments.
	switch char {
	// backspace
	case 8:
		b.values[b.position] = '0'
		if b.position > 0 {
			b.position--
		}
	// tab
	case 9:
		if b.isEndOfLine() {
			b.position++
		} else {
			for !b.isTabStop() {
				b.values[b.position] = '0'
				b.position++
			}
		}

	// carriage return
	case 10:
		// no-op
	// line feed
	case 13:
		for !b.isEndOfLine() {
			b.values[b.position] = '0'
			b.position++
		}
	default:
		b.values[b.position] = char
		b.position++
	}

}

func (b *buffer) writeLn(s string) {
	for _, char := range s {
		b.write(char)

		if b.isEndOfLine() {
			return
		}
	}

	for {
		if b.isEndOfLine() {
			return
		}

		b.write('0')
	}
}

// app
//   windows
//    screens
//      components
//        render/set

func BufferTest() {

	// app.

	var width, height, err = term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		log.Fatalf("Cannot get terminal size because %v", err)
	}
	fmt.Printf("The terminal is %v by %v\n", width, height)

	for range 116 {
		fmt.Print("x")
	}

	var x = []rune{'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x'}
	var buf = buffer{
		top:      0,
		left:     0,
		width:    0,
		height:   0,
		position: 0,
		values:   x[:],
	}

	var subBuf001 = GetSubBuf(buf.values[0:1])
	var subBuf002 = GetSubBuf(buf.values[12:20])

	renderBuf(subBuf002)
	renderBuf(subBuf001)
	renderBuf(buf.values[:])

}

func GetSubBuf(buf []rune) []rune {
	for i := range buf {
		buf[i] = 'y'
	}

	return buf
}

func renderBuf(buf []rune) {
	fmt.Println(string(buf))
}
