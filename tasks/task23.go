package tasks

import "fmt"

func Task23() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(fastUnordered(arr, 2))
	fmt.Println(remove(arr, 2))
}

func fastUnordered(arr []int, i int) []int {
	arr[i] = arr[len(arr)-1]
	arr[len(arr)-1] = 0
	arr = arr[:len(arr)-1]
	return arr
}

func remove(arr []int, i int) []int {
	return append(arr[:i], arr[i+1:]...)
}
