package main

import (
	"time"
	"math/rand"
	"os"
	"fmt"
)

func main() {
	var maxMs int
	var bufSz int
	fmt.Sscanf(os.Args[1], "%d", &maxMs)
	fmt.Sscanf(os.Args[2], "%d", &bufSz)
	b := make([]byte, bufSz)
	for {
		n, err := os.Stdin.Read(b)
		if err != nil {
			break
		}
		os.Stdout.Write(b[:n])
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}