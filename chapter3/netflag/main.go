package main

import (
	"fmt"
)

type Flags uint

const (
	FlagUp Flags = 1 << iota
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

//func IsCast(v Flags)

func main() {
	var v Flags = FlagMulticast | FlagUp
	fmt.Println(v)

}
