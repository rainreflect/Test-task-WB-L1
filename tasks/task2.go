package tasks

import (
	"fmt"
	"time"
)

const (
	n     = 5
	gN    = 3
	gSize = 5
)

func Task2() {

	ch := make(chan int, gSize)
	arr := [n]int{2, 4, 6, 8, 10}

	for i := 0; i < n; i++ {
		go sq(ch, arr[i:i+1])
	}

	for i := 0; i < n; i++ {
		fmt.Println(<-ch)
	}

	time.Sleep(time.Second)
}

func sq(ch chan int, arr []int) {
	for _, v := range arr {
		ch <- v * v
	}
}
