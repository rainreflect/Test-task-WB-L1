package tasks

import "fmt"

func Task17() {
	arr := []int{1, 2, 3, 4, 5, 6, 10, 23, 34, 46}
	fmt.Println(binarySearch(arr, 10))
}

func binarySearch(arr []int, target int) int {
	var res int
	mid := len(arr) / 2

	switch {
	case len(arr) == 0:
		return -1
	case arr[mid] > target:
		res = binarySearch(arr[:mid], target)
	case arr[mid] < target:
		res = binarySearch(arr[mid+1:], target)
		res += mid + 1
	default:
		res = mid
	}
	return res
}
