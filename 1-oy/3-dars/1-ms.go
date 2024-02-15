package main

import "fmt"

func main() {
	var (name string
	     age int
	     work string)
	fmt.Print("Name=")
	fmt.Scan(&name)
	fmt.Print("Age=")
	_,err:=fmt.Scan(&age)
	fmt.Println(err)
	fmt.Print("Work=")
	fmt.Scan(&work)
	//fmt.Printf("My name is= %v age %v work %v\n", name,age,work)
	Sprintf:=fmt.Sprintf("My name is= %v age %v work %v\n", name,age,work)
	fmt.Print(Sprintf)
}
