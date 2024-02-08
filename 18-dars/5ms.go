// package main  //omitepty--- bo'sh bolsa tashlab ketadi 

// import (
// 	"fmt"
// 	"os"
// )
// func main(){
// 	var words []int
// 	var nums []int
// 	filename:="example.txt"
//     readData, err := os.ReadFile(filename)
// 	if err != nil {
// 		fmt.Println("Fayldan o'qishda xatolik:", err)
// 		return
// 	}
// 	defer readData

// }

package main
 
import (
    "fmt"
    "os"
    "bufio"
    "log"
	"strconv"
)
 
func main() {
	var words []string
	var nums []int
 
    file, err := os.Open("example.txt")
    if err != nil {
        log.Fatal(err)
    }
 
    Scanner := bufio.NewScanner(file)
    Scanner.Split(bufio.ScanWords)
 
	for Scanner.Scan() {
		word := Scanner.Text()

		if isInteger(word) {
			num, err := strconv.Atoi(word)
			if err != nil {
				fmt.Println(err)
			} else {
				nums = append(nums, num)
			}
		} else {
			words = append(words, word)
		}
	}
    if err := Scanner.Err(); err != nil {
        log.Fatal(err)
    }
	fmt.Println("words:",words)
fmt.Println("nums",nums)
}



func isInteger(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}