package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])
	for _, s := range os.Args[1:] {
		fmt.Printf("%s ", s)
	}
	fmt.Println()
}
