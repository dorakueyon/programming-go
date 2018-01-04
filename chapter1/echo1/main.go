package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	for _, s := range args {
		fmt.Printf("%s ", s)
	}

}
