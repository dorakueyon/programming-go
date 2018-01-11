package main

import (
	"fmt"
)

func removeDuplicatedInRow(strings []string) []string {
	if len(strings) <= 2 {
		return strings
	}
	i := 0
	for j, s := range strings {
		if j != len(strings)-1 {
			if s != strings[j+1] {
				strings[i] = s
				i++
			}
		} else {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func main() {
	a := []string{"a", "b", "c", "c", "a", "d"}
	b := removeDuplicatedInRow(a)
	fmt.Println(b)
}
