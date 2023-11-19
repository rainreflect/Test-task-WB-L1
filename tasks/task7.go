package tasks

import (
	"fmt"
	"sync"
)

func Task7() {
	m := make(map[int]int, 100)
	mu := sync.Mutex{}

	go func(m map[int]int, mu *sync.Mutex) {
		for i := 0; i < 5; i++ {
			mu.Lock()

			m[i] = i + 1

			mu.Unlock()
		}
	}(m, &mu)
	fmt.Printf("%v", m)
}
