package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	var keysRead = 0
	var buf = [256]byte{}
	for {

		_, err := reader.Read(buf[:])
		if err != nil {
			log.Fatalf("Cannot read rune because %v", err)
		}

		fmt.Println(keysRead, ": ", string(buf[:]), " >> ", buf)

		keysRead++
	}

}
