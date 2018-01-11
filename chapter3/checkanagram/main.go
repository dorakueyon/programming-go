package main

import (
	"fmt"
	"sort"
	"strings"
)

func checkanagram(f, l string) bool {
	sortedf := strings.Split(f, "")
	sort.Strings(sortedf)
	sortedl := strings.Split(l, "")
	sort.Strings(sortedl)
	if a, b := strings.Join(sortedf, ""), strings.Join(sortedl, ""); a == b {
		return true
	} else {
		return false
	}
}

func main() {
	ok := checkanagram("test", "estt")
	fmt.Println(ok)
	fmt.Println(checkanagram("te", "fea"))
	fmt.Println(checkanagram("e", "e"))
}
