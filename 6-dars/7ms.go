//tushunmadm
package main

import (
	"fmt"
	"strconv"
)

func yigindisiJuftBoolsa(a int) bool {
	s := strconv.Itoa(a)
	juftYigindi := 0

	for _, rakamRune := range s {
		rakam, err := strconv.Atoi(string(rakamRune))
		if err != nil {
			fmt.Println("Xato")
			return false
		}

		juftYigindi += rakam
	}

	return juftYigindi%2 == 0
}

func main() {
	// Test qilish
	fmt.Println(yigindisiJuftBoolsa(1234)) // true
	fmt.Println(yigindisiJuftBoolsa(456))  // false
}
