package tasks

import (
	"fmt"
	"os"
	"time"
)

func Task4() {
	var n int
	ch := make(chan int)

	fmt.Println("Введите количество воркеров:")
	fmt.Fscan(os.Stdin, &n)

	for {
		for i := 0; i < n; i++ {
			go worker(ch, i)
		}
		ch <- time.Now().Second()
		time.Sleep(time.Second)
	}
}

func worker(ch chan int, workerId int) {
	fmt.Println(<-ch, workerId)
}
