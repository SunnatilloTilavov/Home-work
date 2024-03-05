// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	fmt.Println("hello1")
// 	go goodbye(&wg)
// 	go helloworld(&wg)
// 	fmt.Println("hello2")
// 	wg.Wait()
// }

// func helloworld(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println("Hello World!")
// }

// func goodbye(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println("Good Bye!")
// }
