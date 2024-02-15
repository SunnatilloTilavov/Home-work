package main 
import "fmt"
func teskari(str string) (result string) { 
  for _, v := range str { 
    result = string(v) + result 
  } 
  return
} 

func main() { 

  str := "12345"
  strRev := teskari(str) 
  fmt.Println(strRev) 
}