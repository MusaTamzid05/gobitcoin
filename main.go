package main

import (
    "errors"
    "fmt"
)

type FieldElement struct {
    num int
    prime int
}

func NewFieldElement(num, prime int) (*FieldElement, error){

    if num >= prime || num < 0  {
        return nil, errors.New("")
    }
    return &FieldElement{num: num, prime: prime}, nil
}

func main() {
    data, _ :=  NewFieldElement(1, 32)
    fmt.Println(data)
}
