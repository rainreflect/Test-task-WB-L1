package tasks

import (
	"fmt"
	"os"
	"time"
)

func Task5() {
	var n time.Duration
	fmt.Println("Введите время существования: ")
	fmt.Fscan(os.Stdin, &n)

	ch := make(chan int)

	go func() {
		for {
			ch <- time.Now().Second()
			time.Sleep(time.Second)

		}
	}()
	go func() {
		for {
			fmt.Println(<-ch)
		}
	}()

	<-time.After(time.Second * n)
}
