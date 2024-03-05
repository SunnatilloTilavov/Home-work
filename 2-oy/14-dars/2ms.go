package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
	var filename string
	var wg sync.WaitGroup
	var sum int = 0
	for i := 1; i <= 3; i++ {
		switch i {
		case 1:
			filename = "example.txt"
		case 2:
			filename = "exp.txt"
		case 3:
			filename = "expp.txt"
		}
		wg.Add(1)
		go summ(filename, &sum, &wg)
	}
	wg.Wait()
	fmt.Println("summ:", sum)

}

func isInteger(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func summ(filename string, sum *int, wg *sync.WaitGroup) {
	defer wg.Done()
	var s int
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	Scanner := bufio.NewScanner(file)
	Scanner.Split(bufio.ScanWords)
	for Scanner.Scan() {
		word := Scanner.Text()
		if !isInteger(word) {
			s++
		}
		*sum *= +s
	}

	defer file.Close()

}

// func WordbyWordScan() {
//     file, err := os.Open("file.txt.txt")
//     if err != nil {
//         log.Fatal(err)
//     }
//     defer file.Close()
//     scanner := bufio.NewScanner(file)
//     scanner.Split(bufio.ScanWords)
//     count := 0
//     for scanner.Scan() {
//         fmt.Println(scanner.Text())
//         count++
//     }
//     if err := scanner.Err(); err != nil {
//         log.Fatal(err)
//     }
//     fmt.Printf("%d\n", count)
// }
