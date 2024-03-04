package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Count int    `json:"count"`
}

func main() {
	var name1, name2, name3 string
	fmt.Println("Insetr name ")
	fmt.Scan(&name1, &name2, &name3)

	client := http.Client{}
	url := fmt.Sprintf("https://api.agify.io?name[]=%s&name[]=%s&name[]=%s", name1, name2, name3)

	request, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer([]byte{}))
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
	

	p := []Person{}
	b, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error", err.Error())
		return
	}

	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println("err", err.Error())
		return
	}
	for _, v := range p {
		fmt.Printf("name:%v age:%v count:%v \n", v.Name, v.Age, v.Count)
	}

}
