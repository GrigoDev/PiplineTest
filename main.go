package main

import (
	"fmt"
	"sync"
)

var counter int

func main() {
	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func() {
			counter++
			fmt.Println(counter)
			wg.Done()
		}()
	}

	wg.Wait()
}

func MaxInt(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
