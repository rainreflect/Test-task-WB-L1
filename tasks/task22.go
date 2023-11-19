package tasks

import (
	"fmt"
	"math"
)

func Task22() {
	var a, b uint64

	a = uint64(math.Pow(2, 20)) + 1
	b = uint64(math.Pow(2, 20)) + 2
	fmt.Println(calculate(a, b, '*'))
}

func calculate(a, b uint64, operation byte) uint64 {

	switch operation {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	case '/':
		return a / b
	}
	return 0
}
