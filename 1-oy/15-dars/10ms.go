package main
import "fmt"
func printValue(i interface{}){

    switch i.(type){ 
       case int: 
       fmt.Println("Int") 
       case string: 
       fmt.Println("string") 
       case bool: 
       fmt.Println("Bool") 
       case float32: 
	   fmt.Println("float") 
	default:
	   fmt.Println("not") 
	}
}

func main(){
	var value1 interface{}
fmt.Scan(&value1)
	printValue(value1)

}