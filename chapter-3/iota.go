package main

import "fmt"

type Flags uint

const (
	FlagUp = 1 << iota
	FlagBroadCast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

func main() {
	fmt.Printf("FlagUp: %#b\n", FlagUp)
	fmt.Printf("FlagBroadCast: %#b\n", FlagBroadCast)
	fmt.Printf("FlagLoopback: %#b\n", FlagLoopback)
	fmt.Printf("FlagPointToPoint: %#b\n", FlagPointToPoint)
	fmt.Printf("FlagMulticast: %#b\n", FlagMulticast)
}
