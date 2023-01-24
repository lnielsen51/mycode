package main

import "fmt"

type Astro struct {
    name string
    age  int
    mission string
    isneeded bool
}

type nasaMission struct {
    people []Astro
    number int
    message string
}

func main() {

    astro1 := Astro{name: "John", age: 20, mission: "mission1", isneeded: true}
    astro2 := Astro{name: "Sally", age: 21, mission: "mission1", isneeded: true}
    astro3 := Astro{name: "Frank", age: 40, mission: "mission2", isneeded:false}
    fmt.Println(astro1)
    fmt.Println(astro2)
    fmt.Println(astro3)

    p := []Astro{astro1, astro2, astro3}

    fmt.Println(p)

    fmt.Println(p[2].mission)

    s := nasaMission{people: p, number: 3, message: "success"}

    fmt.Println(s)
    fmt.Printf("%+v", s)

}

