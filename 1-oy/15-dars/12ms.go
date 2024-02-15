package main

import (
	"fmt"
	"strconv"
)

func printValue(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Int")
	case string:
		fmt.Println("String")
	case bool:
		fmt.Println("Bool")
	case float64:
		fmt.Println("Float")
	default:
		fmt.Printf("Unknown type: %v\n", v)
	}
}

func main() {
	var input string
	fmt.Print("Enter a value: ")
	fmt.Scan(&input)
	if intValue, err := strconv.Atoi(input); err == nil {
		printValue(intValue)
	} else if floatValue, err := strconv.ParseFloat(input, 64); err == nil {
		printValue(floatValue)
	} else if boolValue, err := strconv.ParseBool(input); err == nil {
		printValue(boolValue)
	} else {
		printValue(input)
	}
}
