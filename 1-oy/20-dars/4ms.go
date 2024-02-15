package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var son int
	var name, gmail, jb1 string
	file, err := os.Create("namejob.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("sonni:")
	fmt.Scan(&son)
	for i := 1; i <= son; i++ {
		fmt.Printf("%v-person \n",i)
		fmt.Print("Name:")
		fmt.Scan(&name)
		fmt.Print("job")
		fmt.Scan(&jb1)
		fmt.Print("Email:")
		fmt.Scan(&gmail)
		if strings.Contains(gmail, "@gmail.com") || strings.Contains(gmail,"@mail.com ") {
			_, err := file.WriteString(fmt.Sprintf("Name: %s Job: %s Gmail: %s\n", name, jb1, gmail))
			if err != nil {
				fmt.Println(err)
				return
			}

			}
	
	}
}	


