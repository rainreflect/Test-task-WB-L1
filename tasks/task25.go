package tasks

import (
	"fmt"
	"time"
)

func Task25() {
	fmt.Println("time start")
	leep(3)
}

func leep(x int) {
	<-time.After(time.Second * time.Duration(x))
}
