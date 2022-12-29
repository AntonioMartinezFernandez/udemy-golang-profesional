package main

import (
	"fmt"

	"github.com/AntonioMartinezFernandez/udemy-golang-profesional/topics/25-packages/modules/customMessages"
	"github.com/AntonioMartinezFernandez/udemy-golang-profesional/topics/25-packages/modules/persona"

	"github.com/donvito/hellomod"
)

func main() {
	fmt.Println("Message from main")

	// Module customMessages
	customMessages.PrintHello()
	customMessages.PrintPrivateMessage()

	// Module persona
	jack := persona.Persona{}
	jack.Constructor("Jack", 35)

	jack.SetAge(36)

	fmt.Println(jack)

	// External module
	hellomod.SayHello()
}
