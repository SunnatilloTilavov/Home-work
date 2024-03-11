package main

import(
	"fmt"
	"io"
	"net/http"
	"time"
	"os"
)


func fetchURL(url string,statusChan chan string,statusChan1 chan string){

	resp,err:=http.Get(url)
	if err !=nil{
		statusChan1<-fmt.Sprintf("error",url,err)
	}

	defer resp.Body.Close()

	body,err:=io.ReadAll(resp.Body)
	if err !=nil{
		statusChan1<-fmt.Sprintf("error",url,err)
		return
	}
	statusChan<-fmt.Sprintf("Fetched %s:%d bytes\n",url,len(body))

}

func main(){
	Filename, err := os.Create("Error.txt")
	if err != nil {
	 fmt.Println("Crate error", err)
	 return
	}
	Filename1, err := os.Create("successful.txt")
	if err != nil {
	 fmt.Println("Crate error", err)
	 return
	}
	urls:=[]string{
		"http://www.google.com",
		"http://www.google.com",
		"http://www.github.com",
		"http://www.wikipedia.org",

	}

	statusChan:=make(chan string)
	statusChan1:=make(chan string)
	timeout:=time.After(3 * time.Second)

for _,url:=range urls{

	go fetchURL(url,statusChan,statusChan1)
}

for{
	select{
	case status :=<-statusChan1:
		Filename.Write([]byte(status))
	case error1 :=<-statusChan:
		Filename1.Write([]byte(error1))
	case <-timeout:
		fmt.Println("time out")
	}
}

	

}