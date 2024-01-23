//5-masala
package main

import "fmt"
func raqamlari(a string) {
for i:=1;i<=len(a);i++{
	fmt.Print(i," ")
}
}

func main() {
 var a string
 fmt.Print("A=")
 fmt.Scan(&a)
 raqamlari(a)

}