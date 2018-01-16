package main

import "fmt"
import "github.com/sgoodliff/hello"

func main() {
    fmt.Printf("Hello, World")
    var x string
    x= hello.Hello()
    fmt.Printf(x)
}
