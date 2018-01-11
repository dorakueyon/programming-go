package main

import (
	"fmt"
)

func f() *int {
	v := 1
	return &v

}

func main() {
	v := f()
	fmt.Println(v)
	v = f()
	fmt.Println(v)
	v = f()
	fmt.Println(v)
	v = f()
	fmt.Println(v)
	fmt.Println(*v)

	fmt.Println()
	p := new(int)
	fmt.Println(p)
	q := new(int)
	fmt.Println(q)
	var pc [8]byte
	fmt.Println(pc)
	for i := range pc {
		v := byte(i & 1)
		fmt.Println(v)

	}

}
