package main

import (
	"bufio"
	"fmt"
	"github.com/dorakueyon/programming-go/chapter2/conv"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			convFafrenheit(arg)
			convCelsius(arg)

		}
	} else {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			convFafrenheit(input.Text())
			convCelsius(input.Text())
		}
	}
}

func convFafrenheit(s string) {
	t, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	f := conv.Fahrenheit(t)
	fmt.Println(f)
}
func convCelsius(s string) {
	t, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	c := conv.Celsius(t)
	fmt.Println(c)
}
