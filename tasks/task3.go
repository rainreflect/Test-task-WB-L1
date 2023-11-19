package tasks

import "fmt"

func Task3() {
	ch := make(chan int)
	arr := []int{2, 4, 6, 8, 10}
	var sum int

	go sqSum(arr[len(arr)/2:], ch)
	go sqSum(arr[:len(arr)/2], ch)

	sum += <-ch
	sum += <-ch

	fmt.Println(sum)
}

func sqSum(arr []int, ch chan int) {
	var sum int
	for _, v := range arr {
		sum += v * v
	}
	ch <- sum
}
