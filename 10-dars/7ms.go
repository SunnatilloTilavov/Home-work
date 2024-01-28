package main

import "fmt"

type books struct {
    name string
    author  string
	avaiblecount int
	totalcount int
	b int


}
func(r*books,)Borrow(i int){
	if r.b<r.avaiblecount{
	fmt.Print("you can borrow book")	
	}else{
    fmt.Print("you can't borrow book")
	}
}
func(r*books)Return(){




}
func(r*books)Displayinfo(){
	fmt.Print(%+v)


}

func main(){
	Books:=books{
		name: "absda",
		author: "wadasdas",
		avaiblecount: 105,
		totalcount: 200,

	} 
	Books.Borrow(10)
	Books.Return(5)
	Books.Displayinfo()
}
