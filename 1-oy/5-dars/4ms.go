package main

import "fmt"

func main() {
    for i := 100; i <= 999; i++ {
        a:= i / 100
        b:= (i / 10) % 10
        c:= i % 10
        if (a == b||c == b||c == a)&&!(a==b&&b==c&&c==a){
            fmt.Println(i, )
        } else {
            
        }
    }
}