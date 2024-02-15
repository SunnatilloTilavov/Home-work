package main
import "fmt"
type technical_characteristics struct{
	Motors int
	HorsePower int

}
type Car struct{
	Name string
	Model string
	technical_characteristics
}
func main(){
	P:=Car{
		Name: "Koeniggseg",
		Model:"Agera r",
		technical_characteristics:technical_characteristics{
			Motors: 6,
			HorsePower: 1100,
		},
	}
fmt.Print("P",P.technical_characteristics.Motor)
}