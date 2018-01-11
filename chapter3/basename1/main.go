package main

import (
	"fmt"
)

func basename(s string) string {
	for i := len(s); i > 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			fmt.Printf("%d %s\n", i, s)
			break
		}
	}
	return s
}

func main() {
	w := basename("/feafea/fea/fea/fea.go")
	fmt.Println(w)
}
