package main

import (
    "fmt"
    "sync"
)

var once sync.Once
var hello = 10

func main() {
    fmt.Println(" sync.Once")
    once.Do(addOnce)
    once.Do(addOnce)
    fmt.Println(" sync.Once", hello)
}

func addOnce() {
    hello = hello + 5
}

