package main

import(
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)


func fetchURL(url string , wg *sync.WaitGroup){
	defer wg.Done()

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
	fmt.Printf("Fetched %s:%d bytes\n",url,len(body))

}

func main(){
	now :=time.Now()

	urls:=[]string{
		"http://www.google.com",
		"http://www.google.com",
		"http://www.github.com",
		"http://www.wikipedia.org",

	}
var wg sync.WaitGroup

for _,url:=range urls{
	wg.Add(1)
	go fetchURL(url,&wg)
}

wg.Wait()
	fmt.Println("done it", time.Since(now).Seconds())

}