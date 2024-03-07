package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
now:=time.Now()
 response, err := http.Get("https://jsonplaceholder.typicode.com/comments")
 if err != nil {
  fmt.Println("Get error", err)
  return
 }
 defer response.Body.Close()

 body, err := ioutil.ReadAll(response.Body)
 if err != nil {
  fmt.Println("ReadAll error", err)
  return
 }

 var comments []map[string]interface{}
 err = json.Unmarshal(body, &comments)
 if err != nil {
  fmt.Println("JSON Unmaarshal error", err)
  return
 }
 Data, err := json.MarshalIndent(comments, "", "    ")
 if err != nil {
  fmt.Println("JSON MArshal error", err)
  return
 }
 var a string

 for i:=1;i<=5;i++{
	a=strconv.Itoa(i)+".json"
	Filename, err := os.Create(a)
	if err != nil {
	 fmt.Println("Crate error", err)
	 return
	}
	defer Filename.Close()
    go Writefile1(Data,Filename) //
 }

 fmt.Println("time:",time.Since(now).Seconds())

}

func Writefile1(Data []byte,Filename *os.File){ //,

	Filename.Write(Data)
	fmt.Println("Write data.")
}