package main
import "fmt"
func main(){
	s:=0
	c:=0
a := [3]int{5, 6, 7 }
b := [3]int{3, 6, 10 }
for i:=0;i<=2;i++{
	if a[i]>b[i]{
		s+=1
	}else if a[i]==b[i]{
	s+=0; c+=0
}else{
	c+=1
}
}
fmt.Println(s)
fmt.Println(c)
}

