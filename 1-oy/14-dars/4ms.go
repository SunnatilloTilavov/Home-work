package main

import (
	"fmt"
	"strings"
)

func main() {
	var d string
	fmt.Print(":::")
	fmt.Scan(&d)
	s := strings.ReplaceAll(d, "()", "o")
	fmt.Println(strings.ReplaceAll(s, "(al)", "al"))
}
