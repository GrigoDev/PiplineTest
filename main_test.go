package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestMaxInt(t *testing.T) {
	a, b := 2, 7

	res := MaxInt(a, b)

	if res != b {
		t.Errorf("expected %d, got %d", b, res)
	}
}

func TestMain(m *testing.M) {
	main()
}

func TestRaceInMain(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i) // Здесь будет гонка!
			wg.Done()
		}()
	}

	wg.Wait()
}
