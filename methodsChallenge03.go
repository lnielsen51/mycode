package main

import (
    "fmt"
)

type Player struct{
    Lives int
    Stage int
    Inventory []string
}

func (p *Player) Greenmushroom() {
    p.Lives++
}

func (p *Player) Pickup(item string) {
    p.Inventory = append(p.Inventory, item)
}

func (p Player) CanWhistle() bool {
    if p.Stage >= 5 {
        return true
    }
    return false
}

func main() {
    mario := Player{3, 1, []string{"mushroom"}}
    fmt.Println(mario.Lives)
    mario.Greenmushroom()
    fmt.Println(mario.Lives)

    fmt.Println(mario.Inventory)
    mario.Pickup("leaf")
    fmt.Println(mario.Inventory)

    fmt.Println(mario.CanWhistle())
    mario.Stage = 5
    fmt.Println(mario.CanWhistle())
}
