package tasks

import "fmt"

//Using Percent T with Printf
//Using reflect.TypeOf Function
//Using reflect.ValueOf.Kind() Function

type typeDefiner interface {
}

func Type(tD typeDefiner) string {
	return fmt.Sprintf("var = %T\n", tD)
}

func Task14() {
	var tD typeDefiner
	tD = 0
	fmt.Println(Type(tD))
	tD = "string"
	fmt.Println(Type(tD))
	tD = true
	fmt.Println(Type(tD))
	tD = make(chan any)
	fmt.Println(Type(tD))

}
