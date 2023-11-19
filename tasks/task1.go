package tasks

import "fmt"

type Human struct {
	name    string
	surname string
}

type Action struct {
	Human
	DO string
}

func (h Human) Dream() string {
	return fmt.Sprintf("I'm %s %s", h.name, h.surname)
}

func Task1() {
	a := Action{Human{name: "Roma", surname: "Empire"}, "world lead"}
	fmt.Printf("%s, and my dream is - %s", a.Dream(), a.DO)

}
