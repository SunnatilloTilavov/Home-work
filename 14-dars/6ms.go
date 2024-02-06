package main

import "fmt"

func main() {
	arr := [3]int{7, 1, 0}
	if len(arr) > 1 {
		b := len(arr)
		s := arr[b-1] + 1
		arr[b-1] = s
		fmt.Print(arr)
	}
	if len(arr) == 1 && (arr[1] != 9) {
		s := arr[1] + 1
		fmt.Println(s)
	}
	if len(arr) == 1 && (arr[1] == 9) {
		arr1 := [2]int{1, 0}
		fmt.Println(arr1)

	}

}
