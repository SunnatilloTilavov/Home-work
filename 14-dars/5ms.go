package main

import "fmt"

func main() {
	g := 0
	words := [6]string{"abc", "car", "ada", "racecar", "cool", "asa"}
	for i := 0; i < len(words); i++ {
		b := words[i]
		s := ""

		for a := (len(b) - 1); a >= 0; a-- {
			s = s + string(b[a])
		}
		if s == words[i] && g == 0 {
			fmt.Printf("The first string that is palindromic is ", words[i])
			g += 1

		}
		if s == words[i] && g == 1 {
			fmt.Printf("Note that %v is also palindromic, but it is not the first.\n ", words[i])
			g += 1

		}
		if s == words[i] && g == 2 {
			fmt.Printf("Note that %v is also palindromic, but it is not the first.\n", words[i])
			g += 1

		}
	}
	//fmt.Println("There are no palindromic strings, so the empty string is returned.")
}
