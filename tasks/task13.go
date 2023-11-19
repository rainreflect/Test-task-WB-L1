package tasks

import "fmt"

func Task13() {
	a, b := 5, 10
	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)

}
