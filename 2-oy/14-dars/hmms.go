package main

import (
 "encoding/json"
 "fmt"
 "io/ioutil"
 "net/http"
 "os"
)

func main() {
 url := "https://jsonplaceholder.typicode.com/comments"
 response, err := http.Get(url)
 if err != nil {
  fmt.Println("HTTP so'rov xatosi:", err)
  return
 }
 defer response.Body.Close()

 body, err := ioutil.ReadAll(response.Body)
 if err != nil {
  fmt.Println("Ma'lumotlarni o'qishda xato:", err)
  return
 }

 var comments []map[string]interface{}
 err = json.Unmarshal(body, &comments)
 if err != nil {
  fmt.Println("JSON ma'lumotlarni o'qishda xato:", err)
  return
 }


 outputFile, err := os.Create("1.json")
 if err != nil {
  fmt.Println("Faylni yaratishda xato:", err)
  return
 }
 defer outputFile.Close()

 encodedData, err := json.MarshalIndent(comments, "", "    ")
 if err != nil {
  fmt.Println("JSON ma'lumotlarni kodlashda xato:", err)
  return
 }

 outputFile.Write(encodedData)
 fmt.Println("Ma'lumotlar boshqa faylga yozildi.")
}