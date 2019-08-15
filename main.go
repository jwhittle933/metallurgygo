package main

import (
	"sync"
)

func main() {
	// get flags
	args := ParseFlags()
	NewFiles(args).
		Decode().
		Encode().
		Write()
}

// Await ...
func Await(files Files) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
	}()
	wg.Wait()
}
