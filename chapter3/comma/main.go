package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + comma(s[n-3:])

}

//func comma(s string) string {
//	var buf bytes.Buffer
//	for i := len(s); i > 0; i-- {
//		if (i+1)%3 == 0 {
//			buf.WriteString(",")
//
//		}
//
//	}
//}

func main() {
	i := comma("1000877r8478433430")
	fmt.Println(i)
}
