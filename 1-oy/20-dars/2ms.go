package main

import (
 "encoding/json"
 "fmt"
 "os"
)

type Animal struct {
 Name  string json:"name"
 Sex   string json:"sex"
 Class string json:"class"
}

func main() {
 animals := make([]Animal, 0)

 for i := 0; i < 5; i++ {
  var name, sex, class string

  fmt.Print("Name: ")
  fmt.Scan(&name)

  fmt.Print("Sex: ")
  fmt.Scan(&sex)

  fmt.Print("Class: ")
  fmt.Scan(&class)

  animals = append(animals, Animal{Name: name, Sex: sex, Class: class})
 }

 file, err := os.Create("animals.json")
 if err != nil {
  fmt.Println("Fayl yaratilmadi", err)
  return
 }
 defer file.Close()

 encoder := json.NewEncoder(file)
 encoder.SetIndent("", " ")
 err = encoder.Encode(animals)
 if err != nil {
  fmt.Println("Error fayl encoder")
 }

 err = json.NewEncoder(file).Encode(animals)
 if err != nil {
  fmt.Println("Failed to encode JSON:", err)
  return
 }

 fmt.Println("Animals data has been written to animals.json")
}