package main

import (
	"fmt"
)

func noempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
		fmt.Println(strings)
	}
	return strings[:i]
}

func main() {
	a := []string{"a", "fea", "", "fe"}
	b := noempty(a)
	fmt.Println(b)
}
