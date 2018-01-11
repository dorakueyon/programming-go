package main

import (
	"fmt"
)

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
		fmt.Println(s)
	}
}

func reverse2(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
		fmt.Println(s)
	}
}

func main() {
	a := []int{1, 2, 3, 4, 4, 2, 3, 4, 1, 5, 6}
	//	reverse(a)
	//	fmt.Println(a)
	reverse(a[:2])
	reverse(a[2:])
	reverse(a[:])
}
