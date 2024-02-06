package main

import "fmt"

func printValue(i interface{}) {
    switch i.(type) {
    case int:
        fmt.Println("Int")
    case string:
        fmt.Println("String")
    case bool:
        fmt.Println("Bool")
    case float32:
        fmt.Println("Float")
    default:
        fmt.Println("Not recognized")
    }
}

func main() {
    var value1 interface{}
    fmt.Print("Enter a value: ")
    fmt.Scan(&value1)
    printValue(value1)
}
