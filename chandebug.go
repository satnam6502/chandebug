package main

import (
	"fmt"
)

func precharge() chan bool {
	ch := make(chan bool, 20)
	for i := 0; i < 20; i++ {
		ch <- true
	}
	return ch
}

func canAccept(ch chan bool) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}

func main() {
	for i := 0; i < 100000; i++ {
		ch := precharge()
		if !canAccept(ch) {
			fmt.Printf("%d: Can't accept\n")
		}
	}
	fmt.Print("[done]\n")
}
