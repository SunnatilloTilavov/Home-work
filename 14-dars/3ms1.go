package main

import (
	"fmt"
	"strconv"
)

func main() {
	var coordinates string
	fmt.Println("Yacheykani kiritin:")
	fmt.Scan(&coordinates)
	for b := 0; b <= 3; b++ {
		for i := 1; i <= 8; i += 2 {
			arr := [4]string{"a", "c", "e", "g"}
			arr1 := [4]string{"b", "d", "h", "f"}
			a := arr[b] + strconv.Itoa(i)
			d := arr1[b] + strconv.Itoa(i)
			if coordinates == a {
				fmt.Print("qora")
				break
			}
			if coordinates == d {
				fmt.Print("oq")
				break
			}
		}

		for i := 2; i <= 8; i += 2 {
			arr := [4]string{"b", "d", "h", "f"}
			arr1 := [4]string{"a", "c", "e", "g"}
			s := arr[b] + strconv.Itoa(i)
			c := arr1[b] + strconv.Itoa(i)
			if coordinates == s {
				fmt.Print("qora")
				break
			}
			if coordinates == c {
				fmt.Print("oq")
				break
			}
		}
	}
}
