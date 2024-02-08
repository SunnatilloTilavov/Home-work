package main

import (
	"encoding/json"
	"fmt"
	"os"
	"errors"
	"math/rand"
	"strconv"
)
type People struct{
	Name string `json:"Name"`
	Age int `json:"age"`
	Course int `json:"course"`
	Contract int `json:"contract"`
}

type TotalPeople struct{
	Course int
	Count int
	Total int
}


func main(){
	file,err:=os.Open("people.json")
	if err !=nil{
		fmt.Println("error",err)
	}
	defer file.Close()

	decoder:=json.NewDecoder(file)
	var(
		people []People
	)
	err=decoder.Decode(&people)

	if err !=nil{
		fmt.Println("errorr")
	}
	fmt.Printf("%+v",people)

	for _, name := range people{
		a:=string(name.Name)
		if checkFileExists(a){
			b:=rand.Intn(100)
			a+=strconv.Itoa(b)
		file,err:=os.Create(a)
		if err !=nil{
		fmt.Print("errrrrrr",err)	
		}
		defer file.Close()

		}else{
		file,err:=os.Create(a)
		if err !=nil{
		fmt.Print("errrrrrr",err)	
		}
		defer file.Close()
	}



}

}
func  checkFileExists(filePAth string)bool{
	_,error:=os.Stat(filePAth)
	return !errors.Is(error,os.ErrNotExist)
}


