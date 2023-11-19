package tasks

import (
	"fmt"
	"math"
)

func Task10() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	m := make(map[int][]float64)

	for _, v := range arr {
		key := int(math.Floor(v/10) * 10)
		m[key] = append(m[key], v)

	}
	fmt.Printf("%v", m)

}
