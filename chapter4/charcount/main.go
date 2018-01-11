package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

var counts = make(map[string]map[rune]int)

func addEdge(from string, r rune) {
	edges := counts[from]
	if edges == nil {
		edges = make(map[rune]int)
		counts[from] = edges
	}
	edges[r]++
}
func main() {
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			addEdge("letter", r)
		} else if unicode.IsDigit(r) {
			addEdge("digit", r)
		} else {
			addEdge("other", r)
		}
		utflen[n]++
	}
	fmt.Printf("rune\tcounts\n")
	for c1, n1 := range counts {
		for c2, n2 := range n1 {
			fmt.Printf("%s\t%q\t%d\n", c1, c2, n2)
		}
	}
	fmt.Printf("\nlen\tcounts\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
