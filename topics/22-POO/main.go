package main

import "fmt"

func main() {
	dependencies := DI{}

	collie := Pet{"Collie", 2, &dependencies}
	fmt.Println(collie)
	collie.sayHello()
	collie.sayAge()

	tobby := Pet{
		name: "Tobby",
		age:  3,
		di:   &dependencies,
	}
	fmt.Println(tobby)
	tobby.sayHello()
	tobby.sayAge()
	fmt.Println(tobby.di.diFunction())

	michi := Pet{
		name: "Michi",
		age:  5,
		di:   &dependencies,
	}

	cat := Cat{
		Pet:   &michi,
		lifes: 7,
	}
	fmt.Println(cat)
	cat.sayHello()
	cat.sayAge()
	fmt.Println("Lifes: ", cat.lifes)
	fmt.Println(cat.di.diFunction())
}

/* Struct (Class) */
type Pet struct {
	// Attributes
	name string
	age  int
	di   *DI
}

// Methods
func (p *Pet) sayHello() {
	fmt.Printf("Hello, I'm %s!\n", p.name)
}

func (p *Pet) sayAge() {
	fmt.Printf("I'm %d years old\n", p.age)
}

/* Composition (one struct contain another struct as attribute, in this case, a pointer of the struct) */
type Cat struct {
	*Pet
	lifes int
}

/* Struct (Dependency Injection Class) */
type DI struct{}

func (di *DI) diFunction() string {
	return "I'm a injected function"
}
