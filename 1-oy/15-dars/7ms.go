package main

import "fmt"

type Book interface {
	Author() string
	Title() string
	Count() string
}

func PrintBook(a Book) {
	fmt.Println("Author:",a.Author())
	fmt.Println("Title:",a.Title())

}

type Book1 struct{}

func (d Book1) Author() string {
	return "BookaAuthor1"
}
func (d Book1) Title() string {
	return "BookTitle1"
}

type Book2 struct{}

func (d Book2) Author() string {
	return "BookaAuthor2"
}
func (d Book2) Title() string {
	return "BookTitle2"
}

func main(){ 
a:=Book1{}
b:=Book2{}

PrintBook(a)
PrintBook(b)




}