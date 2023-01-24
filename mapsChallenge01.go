package main

import (
    "fmt"
)

func main() {
    programmingLanguages := map[string]string{
        "Python": ".py",
        "Golang": ".go",
        "Java": ".java",
        "Ansible": ".yml",
        "C++": ".cpp",
    }

    fmt.Println("map:", programmingLanguages)

    delete(programmingLanguages, "C++")
    programmingLanguages["Julia"] = ".jl"

    fmt.Println("map:", programmingLanguages)
}

