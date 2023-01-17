package main

import (
    "fmt"
    "musa.io/bitcoin/lib"
)

func main() {
    a, _ :=  lib.NewFieldElement(3, 13)
    b, _ :=  lib.NewFieldElement(1, 13)

    d, _ := a.Pow(-3)
    fmt.Println(d.Equal(b))

    _, err := lib.NewPoint(-1, -2, 5, 7)

    if err != nil {
        fmt.Println(err)
    }


}
