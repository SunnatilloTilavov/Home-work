//Practice:  Struct, packages
package main
import "fmt"
type Address struct {
    Name    string
    city    string
    Pincode int
}
func main(){
	a1 := Address{"A1", "D1", 3623572}
    fmt.Println("Address1: ", a1)
	a2 := Address{Name: "A2", city: "B2",
	Pincode: 277001}
    fmt.Println("Address2: ", a2)
    a3 := Address{Name: "D1"}
    fmt.Println("Address3: ", a3)
}