package tasks

import "fmt"

func Task19() {

	s := "главрыба"

	fmt.Println(reverse(s))
}

func reverse(s string) string {
	var res []rune
	runeS := []rune(s)
	for i := len(runeS) - 1; i >= 0; i-- {
		res = append(res, runeS[i])
	}
	return string(res)
}
