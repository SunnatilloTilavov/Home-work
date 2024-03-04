// package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// )

// type Person1 struct {
// 	Name string `json:"name"`
// 	Age  int    `json:"age"`
// }

// func main() {
// 	var name string
// 	var age int
// 	client := http.Client{}

// 	fmt.Println("Insert name:")
// 	fmt.Scan(&name)
// 	fmt.Println("Age:")
// 	fmt.Scan(&age)

// 	person := Person1{
// 		Name: name,
// 		Age:  age,
// 	}

// 	jsonData, err := json.Marshal(person)
// 	if err != nil {
// 		fmt.Println("JSON formatting error:", err)
// 		return
// 	}

// 	url := "https://bec6-94-232-24-122.ngrok-free.app/login"
// 	req, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		fmt.Println("Error creating request:", err)
// 		return
// 	}

// 	req.Header.Set("Content-Type", "application/json")

// 	fmt.Println("Request URL:", req.URL.String())

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("Error making request:", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	fmt.Println("Response Status:", resp.Status)
// }
