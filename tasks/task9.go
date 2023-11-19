package tasks

import (
	"fmt"
)

const (
	N = 10
)

func Task9() {
	arr := [N]int{1, 2, 3, 4}

	sender := generatingChan(arr)
	receiver := squaringChan(sender)
	for values := range receiver {
		fmt.Println(values)
	}
}

func generatingChan(arr [N]int) <-chan int {
	sender := make(chan int)
	go func() {
		for _, num := range arr {
			sender <- num
			fmt.Println("Заполнение значений канала")
		}
		close(sender)
	}()
	return sender
}

func squaringChan(receivingValue <-chan int) <-chan int {
	receiver := make(chan int)
	go func() {
		for num := range receivingValue {
			fmt.Println("Заполнение второго канала квардратами выгруженных чисел")
			receiver <- num * num
		}
		close(receiver)
	}()
	return receiver
}
