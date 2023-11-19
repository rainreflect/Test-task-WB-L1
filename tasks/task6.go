package tasks

import (
	"context"
	"fmt"
	"time"
)

func Task6() {
	//остановка горутины закрытием канала
	ch1 := make(chan string)
	go func() {
		for {
			for v := range ch1 {
				fmt.Println(v)
			}
		}
	}()

	ch1 <- "hello"
	ch1 <- "world"

	close(ch1)
	fmt.Println("Chan closed")
	time.Sleep(time.Second)

	//периодический пулинг каналов
	ch2 := make(chan string)
	done := make(chan struct{})
	go func() {
		for {
			select {
			case ch2 <- "foo":
			case <-done:
				close(ch2)
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		time.Sleep(3 * time.Second)
		done <- struct{}{}
	}()

	for i := range ch2 {
		fmt.Println("Received: ", i)
	}

	fmt.Println("Finish")

	//использование контекста
	ch3 := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				ch3 <- struct{}{}
				return
			default:
				fmt.Println("foo...")
			}

			time.Sleep(500 * time.Millisecond)
		}
	}(ctx)

	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()

	<-ch3
	fmt.Println("Finish")
}
