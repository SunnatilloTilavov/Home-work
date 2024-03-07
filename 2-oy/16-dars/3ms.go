package main

import(
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func main(){
	
   
	now :=time.Now()

	urls:=[]string{
		"http://www.google.com",
		"http://www.google.com",
		"http://www.github.com",
		"http://www.wikipedia.org",

	}
for _,url:=range urls{
	randomch, readablech := time.fetchURL1()
}
for {
	select {
	case test := <-randomch:
		fmt.Println(test)
		// time.Sleep(1 *time.Second)
	case status := <-readablech:
		fmt.Print(status)
	}
}
}

func fetchURL1(url string ,chan int, <-chan string){
	readablech := make(chan string)

	resp,err:=http.Get(url)
	if err !=nil{
		fmt.Println("error",url,err)
	}

	defer resp.Body.Close()

	body,err:=io.ReadAll(resp.Body)
	if err !=nil{
		fmt.Println("error",url,err)
		return
	}
	
	go func() {
		defer close(randomch)
		defer close(readablech)
		for {
			select {
			case randomch <- body:
				readablech <- "text:"
			case <-time.After(time.Second):
				readablech <- "time out"
				return
			}

		}
	}()
	return randomch, readablech

}

