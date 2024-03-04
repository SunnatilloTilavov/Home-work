package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Person1 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// var name string
	// var age int
	client := http.Client{}
	// fmt.Println("Insetr name ")
	// fmt.Scan(&name)
	// fmt.Println("Age")
	// fmt.Scan(&age)
	person := Person1{
		Name: "Sunnatillo",
		Age:  123,
	}

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("JSON formatiga o'tkazishda xato:", err)
	}

	url := "https://e62d-94-232-24-122.ngrok-free.app/login"

	request, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("err", err.Error())
		return

	}

	fmt.Println("rec url", request.URL.String())

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("errr", err)
		return

	}

	fmt.Println("status code:", response.Status)

	b, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error", err.Error())
		return
	}
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		fmt.Println(string(b))
	} else {
		fmt.Println("Status code 200<x>300")
	}
}
