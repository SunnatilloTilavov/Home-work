package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	var name,a string
	s:=".txt"
	fmt.Print("name:")
	fmt.Scan(&name)
	a=name+s

	// _, err := os.Stat(a)
	// b:=!errors.Is(err, os.ErrNotExist)

	if checkFileExists(a) {
	file, err := os.Create(a)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() 
}else{
	b:=name+strconv.Itoa(rand.Intn(10))+s
	if checkFileExists1(b){
	file, err := os.Create(b)
	if err != nil {
		fmt.Println(err)
		return}
		defer file.Close()
	}else{
		d:=name+strconv.Itoa(rand.Intn(10))+s
		file, err := os.Create(d)
	if err != nil {
		fmt.Println(err)
		return
		
	}
	defer file.Close() }
}
}


func checkFileExists(a string) bool {
	_, err := os.Stat(a)
	return errors.Is(err, os.ErrNotExist)
  }


func checkFileExists1(b string) bool {
	_, err := os.Stat(b)
	return errors.Is(err, os.ErrNotExist)
  }

