package main

import (
    "fmt"
    "math"
)

func main() {
    var binary, reversed string // declaring strings to store user input and reversed input
    fmt.Print("Enter binary number: ") // input message and scanning
    fmt.Scanln(&binary)

    // reversing the input string 
    for r := len(binary) - 1; r >= 0; r-- { 
        reversed += string(binary[r])
    }

    // deferred recovery function in case of panic cause by non-binary input
    defer func() { 
        if r := recover(); r != nil {
            fmt.Println("Error:\n", r)
        }
    }()

    var total int32 // storing the accumulated total 
    for i, v := range reversed {
        if v != '0' && v != '1' { // determining if input is binary, panic otherwise 
            panic("Invalid input: only 0's and 1's are accepted ")
        }
        if v == '1' { // incrementing total by 2^index at every occurence of a 1 in the input
            total += int32(math.Pow(2, float64(i)))
        }
    }
    fmt.Println(total) // printing total
}