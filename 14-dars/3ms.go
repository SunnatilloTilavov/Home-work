package main

import (
  "fmt"
  "strconv"
)

func main() {
	coordinates:="c5"
	    for b:=0;b<=3;b++{
	        for i:=1;i<=8;i+=2{
		  arr := [4]string{"a","c","e","g"}
		  s:=arr[b]+strconv.Itoa(i)
		  if coordinates==s{
			  fmt.Print("a")
			  break
		  }
		}
	  }
	  for b:=0;b<=3;b++{
	  for i:=2;i<=8;i+=2{
	  arr := [4]string{"b","d","h","f"}
	  s:=arr[b]+strconv.Itoa(i)
	  if coordinates==s{
		  fmt.Print("a")
		  break
	  }
	}
  }

  for b:=0;b<=3;b++{
	for i:=2;i<=8;i+=2{
  arr := [4]string{"a","c","e","g"}
  s:=arr[b]+strconv.Itoa(i)
  if coordinates==s{
	  fmt.Print("b")
	  break
  }
}
}
for b:=0;b<=3;b++{
	for i:=1;i<=8;i+=2{
	arr := [4]string{"b","d","h","f"}
	s:=arr[b]+strconv.Itoa(i)
	if coordinates==s{
		fmt.Print("b")
		break
	}
  }
}
	}
