package main 

import ( 
	"fmt"
) 

type Address struct { 
	City string
	State string
} 

type Person struct { 
	Name string
	Address 
} 

func main() { 
	p := Person{ 
		Name: "John Doe", 
		Address: Address{ 
			City: "New York", 
			State: "NY", 
		}, 
	} 

	fmt.Println("Name:", p.Name) 
	fmt.Println("City:", p.City) 
	fmt.Println("State:", p.State) 
	fmt.Println("%v",p)
} 
