package main

import (
	"fmt"
	//"strings"
	"time"
)
func main(){
	n:=100
//sentense:="adafrgshb,shvjkbg"
//joing :=[]string{" i"," have ","a" ,"macbook"}
//subsstring:="macbook"
//com:=strings.Contains(sentense ,subsstring)  borligini ekshiradi
//comm:=strings.Count(sentense,subsstring)  necha marta qatnashganini
//com:=strings.joing(sentense ,subsstring)
//fmt.Println(com,comm)
now:=time.Now()//.Add(time.Hour*10)
fmt.Println(now)

//for i:=0;i<=n;i++{
//time.Sleep(time.Microsecond)
//fmt.Println("i:",i)
time.Sleep(time.Millisecond==1)
fmt.Println(time.Since(now).Seconds())




}




}