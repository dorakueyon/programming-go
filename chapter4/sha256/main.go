package main

import (
	"crypto/sha256"
	"fmt"
)

func checkSameDegit(a, b [32]uint8) int {
	var count = 1
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			count++
		}
	}
	return count
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("Xafe"))
	fmt.Println(c1 == c2)
	fmt.Printf("c1:%T c2:%T", c1, c2)
	fmt.Println()
	c := checkSameDegit(c1, c2)
	fmt.Printf("%x\n%x\n%d\n", c1, c2, c)
}
