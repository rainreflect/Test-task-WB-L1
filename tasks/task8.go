package tasks

import "fmt"

func Task8() {
	index := 25
	decision := true
	num := changeBit(index, decision)
	fmt.Printf("%064b\n", num)
	fmt.Println(num)
}

func changeBit(i int, dec bool) int64 {
	switch dec {
	case true:
		var value int64 = 1 << i
		return value
	case false:
		var value int64 = 0 << i
		return value
	}
	return 0
}
