package main

import (
	"errors"
	"fmt"
)

func twoNumb(a, b int) (int, error) {
	if b!= 0 {
		return a / b, nil
	} else {
		return 0, errors.New("error 2-numb 0")
	}
	
}

func main() {
	a := 15
	b := 0
	x, err := twoNumb(a, b)
	if err != nil {
		fmt.Println(err)	
	} else {
		fmt.Println("a/b=", x)
	}

}
