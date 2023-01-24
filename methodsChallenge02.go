package main

import (
    "fmt"
)

type Virtmach struct{
    ip string
    hostname string
    diskgb int
    ram int
}

func (a Virtmach) getDiskgb() int {
    return a.diskgb
}

func (a *Virtmach) increaseRam() {
    a.ram++
}

func main() {
    machine1 := Virtmach{ip: "10.10",hostname: "machine1",diskgb: 8,ram: 16}
    fmt.Println(machine1)
    fmt.Println(machine1.getDiskgb())
    machine1.increaseRam()
    fmt.Println(machine1)
}

