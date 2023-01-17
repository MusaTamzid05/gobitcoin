package main

import (
    "musa.io/bitcoin/lib"
)

func main() {
    a, _ :=  lib.NewFieldElement(0, 223)
    b, _ :=  lib.NewFieldElement(7, 223)
    x, _ :=  lib.NewFieldElement(192, 223)
    y, _ :=  lib.NewFieldElement(105, 223)


    result := lib.NewPoint(x, y, a, b)
    result.Print()


}
