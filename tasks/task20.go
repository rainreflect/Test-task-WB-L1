package tasks

import (
	"fmt"
	"strings"
)

func Task20() {
	s := "snow dog sun" // sun dog snow
	fmt.Println(rotate(s))
}

func rotate(s string) string {
	arr := strings.Split(s, " ")
	left, right := 0, len(arr)-1
	for left < right {
		temp := arr[right]
		arr[right] = arr[left]
		arr[left] = temp
		left++
		right--
	}
	res := strings.Join(arr, " ")
	return res
}
