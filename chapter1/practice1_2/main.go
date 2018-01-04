package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])
	for i, s := range os.Args[1:] {
		fmt.Printf("%d %s\n", i+1, s)
	}
}
