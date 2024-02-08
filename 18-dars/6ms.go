package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Car struct {
	Brand string `json:"brand"`
	Year  int    `json:"year"`
	Model string `json:"model"`
	Power int    `json:"power"`
}

func main() {
	file, err := os.Open("car.json")
	if err != nil {
		fmt.Println("eroor:", err)
		return
	}
	defer file.Close()

	var cars []Car
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cars); err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%+v",cars)
}
