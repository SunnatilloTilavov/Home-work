package main
import "fmt"
func main() {
    var a int 
    fmt.Print("Qator soni: ")
    fmt.Scan(&a)

    for i := 1; i <= a; i++ {
        for j := 1; j <= i; j++ {
            fmt.Print("* ")
        }
        fmt.Println()
    }
}
