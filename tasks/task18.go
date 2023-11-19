package tasks

import (
	"fmt"
	"sync"
	"time"
)

type incr struct {
	i int
}

func Task18() {
	inc := incr{i: 0}
	mu := sync.Mutex{}

	for i := 0; i < 10; i++ {
		go func(inc *incr, mu *sync.Mutex) {
			mu.Lock()
			inc.i++
			mu.Unlock()
			time.Sleep(time.Second)
		}(&inc, &mu)

	}

	time.Sleep(time.Second)
	fmt.Println(inc.i)
}
