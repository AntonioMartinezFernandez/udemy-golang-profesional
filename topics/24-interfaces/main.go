package main

import "fmt"

func main() {
	tobby := Dog{}
	wanda := Fish{}
	canary := Bird{}

	moveAnimal(&tobby)
	moveAnimal(&wanda)
	moveAnimal(&canary)
}

type Animal interface {
	move() string
}

func moveAnimal(animal Animal) {
	fmt.Println(animal.move())
}

type Dog struct{}
type Fish struct{}
type Bird struct{}

func (*Dog) move() string {
	return "I'm a dog and I'm walking"
}

func (*Fish) move() string {
	return "I'm a fish and I'm swimming"
}

func (*Bird) move() string {
	return "I'm a bird and I'm flying"
}
