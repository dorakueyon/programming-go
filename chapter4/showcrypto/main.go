package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
)

func showHash(s string) {
	var c [32]byte
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "256":
			c = sha256.Sum256([]byte(s))
		default:
			c = sha256.Sum256([]byte(s))
		}
	} else {
		c = sha256.Sum256([]byte(s))
	}
	fmt.Printf("%x\n", c)
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		showHash(input.Text())
	}
}
