package main

import (
    "fmt"
    "math/rand"
    "time"
)

func init() {
    rand.Seed(time.Now().Unix())
}

func main() {
    names := []string{"Tom","Freddy","Jason","Michael","Pennywise","Chucky","Pinhead"}

    if x := names[rand.Intn(len(names))]; x== "Freddy" {
        fmt.Println("Friday the 13th", x)
    } else if x == "Michael" {
        fmt.Println("Halloween", x)
    } else if x == "Jason" {
        fmt.Println("Jason", x)
    } else if x == "Pennywise" {
        fmt.Println("IT", x)
    } else if x == "Chucky" {
        fmt.Println("Chucky", x)
    } else if x == "Pinhead" {
        fmt.Println("Hellraiser", x)
    } else {
        fmt.Println("none", x)
    }
}
