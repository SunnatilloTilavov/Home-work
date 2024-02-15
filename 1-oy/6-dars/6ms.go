package main

import "fmt"
func raqamlari(a int) {
var t bool
var b string
s:=0
for i:=1;i<=len(a);i++{
	s+=i
}
 if s%2==0 {
  t=true
 } else {
  t=false
 }
 fmt.Print(t)
}

func main() {
 var a int
 fmt.Print("A=")
 fmt.Scan(&a)
 raqamlari(a)

}