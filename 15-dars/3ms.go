package main

import "fmt"

type Animal interface {
	MakeSaund() string
	Move() string
}

func PrintSount(a Animal) {
	fmt.Printf("Sount: %s \n", a.MakeSaund())
	fmt.Printf("Movement: %s \n", a.Move())
}

type Dog struct{}

func (a Dog) MakeSaund() string {
	return "Woof"
}

func (d Dog) Move() string {
	return "Running"
}

type Bird struct{}

func (b Bird) MakeSaund() string {
	return "Tweet"
}

func (b Bird) Move() string {
	return "flying"
}



func main() {
	dog := Dog{}
	bird := Bird{}

	PrintSount(dog)
	PrintSount(bird)
}
