package tasks

import "fmt"

func Task11() {
	arr1 := []int{1, 2, 10, 23, 54}
	arr2 := []int{-3, -10, 1, 2, 10, 233, 542, 23}

	var res []int
	m := make(map[int]bool)
	for _, num := range arr1 {
		m[num] = true
	}
	for _, num := range arr2 {
		if m[num] {
			res = append(res, num)
		}
	}
	fmt.Println(res)
}
