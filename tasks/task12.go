package tasks

import "fmt"

func Task12() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	s := set(arr)
	fmt.Println(s)
}

func set(arr []string) []string {
	var res []string
	m := make(map[string]bool)
	for _, value := range arr {
		m[value] = true
	}

	for key, value := range m {
		if value {
			res = append(res, key)
		}
	}

	return res
}
