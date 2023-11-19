package tasks

import "fmt"

//Классы cat и dog чужие, в чужих библиотеках, мы не можем изменять их.

type Cat struct {
}

func (c *Cat) MeowMeow() {
	fmt.Println("MEEEEEEEEEEEEEEEEEEEEEEEEEEEEEOW")
}

type Dog struct {
}

func (d *Dog) WoofWoof() {
	fmt.Println("WOOF, wof wof")
}

type AnimalAdapter interface {
	Noise()
}

type CatAdapter struct {
	*Cat
}

type DogAdapter struct {
	*Dog
}

func (c *CatAdapter) Noise() {
	c.MeowMeow()
}
func (d *DogAdapter) Noise() {
	d.WoofWoof()
}

func NewCatAdapter(c *Cat) AnimalAdapter {
	return &CatAdapter{c}
}

func NewDogAdapter(d *Dog) AnimalAdapter {
	return &DogAdapter{d}
}

func Task21() {
	myAnimals := [2]AnimalAdapter{NewCatAdapter(&Cat{}), NewDogAdapter(&Dog{})}

	for _, animal := range myAnimals {
		animal.Noise()
	}
}
