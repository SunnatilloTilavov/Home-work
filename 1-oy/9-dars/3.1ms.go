package main
import "fmt"
func main(){
	var target int
	fmt.Print("target=")
	fmt.Scan(&target)
nums := [10]int{5, 6, 7,3,5,1,6,2,4,9 }
for i:=0;i<len(nums);i++{
	for b:=0;b<len(nums);b++{
		 if nums[i]+nums[b]==target{
			fmt.Println("tartib raqami",i+1,"va",b+1)
			fmt.Println(nums[i],"+",nums[b],"=",target)
  
		 }
	}	
}
}
//b:=(len(nums)-1);b>=0;b--