package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    content, err := ioutil.ReadFile("file.txt")
    if err != nil {
        fmt.Println("Xato:", err)
        return
    }

    fmt.Println("Fayl ma'lumotlari:", string(content))
}
