package main

import(
	"fmt"
"time"
)

func  worker(id int,done chan struct{}){
	defer func(){
done <-struct{}{}
	}()
fmt.Printf("worker %d started\n",id)
time.Sleep(time.Second)
fmt.Printf("worker %d finish \n",id)
}


func main(){
numWorkers:=5
done:=make(chan struct{})

for i:=0; i<numWorkers;i++{
go worker(i,done)
}

timeout:=time.After(2*time.Second)
for i:=0;i<numWorkers;i++{
	select{
	case<-done:
		fmt.Println("received signal from a worker. That worker is done")
	case <-timeout:
		fmt.Println("Time out")
		return
	}
}
fmt.Println("All workers finished")
}