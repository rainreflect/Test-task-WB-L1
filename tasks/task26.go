package tasks

import (
	"fmt"
	"strings"
)

func Task26() {
	s := "aAbBc"
	s1 := "aabcd"
	fmt.Println(isUniq(s))
	fmt.Println(isUniq(s1))
}

func isUniq(s string) bool {
	m := make(map[string]bool)
	slice := strings.Split(s, "")
	for _, char := range slice {
		if !m[char] {
			m[char] = true
		} else {
			return false
		}
	}
	return true
}
