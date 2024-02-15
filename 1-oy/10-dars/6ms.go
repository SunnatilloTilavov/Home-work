package main

import "fmt"

type car struct{
	color string
	year string
	maxspeed int
	brand string
	mechanic bool
}

func main(){
	abs:=car{
		color:"red",
		year:"2023",

	}
	fmt.Println(car.color)

}