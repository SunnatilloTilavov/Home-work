package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
var wg sync.WaitGroup
now:=time.Now()

 url := "https://jsonplaceholder.typicode.com/comments"
 response, err := http.Get(url)
 if err != nil {
  fmt.Println("http.Get:", err)
  return
 }
 defer response.Body.Close()

 body, err := ioutil.ReadAll(response.Body)
 if err != nil {
  fmt.Println("ReadAll:", err)
  return
 }

 var comments []map[string]interface{}
 err = json.Unmarshal(body, &comments)
 if err != nil {
  fmt.Println("Unmarshal:", err)
  return
 }
 Data, err := json.MarshalIndent(comments, "", "    ")
 if err != nil {
  fmt.Println("MarshalIndent:", err)
  return
 }
 var a string

 for i:=1;i<=5;i++{
	a=strconv.Itoa(i)+".json"
	Filename, err := os.Create(a)
	if err != nil {
	 fmt.Println("Create:", err)
	 return
	}
	defer Filename.Close()
    go Writefile(Data,Filename,&wg) //
	wg.Add(1)
 }
 wg.Wait()
 fmt.Println("time:",time.Since(now).Seconds())

}

func Writefile(Data []byte,Filename *os.File,wg *sync.WaitGroup){ //,
	defer wg.Done()
	Filename.Write(Data)
	fmt.Println("Write Data .")
}