package main

import (
    "musa.io/bitcoin/lib"
    "fmt"
)

func main() {
    a, _ :=  lib.NewFieldElement(0, 223)
    b, _ :=  lib.NewFieldElement(7, 223)
    x1, _ :=  lib.NewFieldElement(192, 223)
    y1, _ :=  lib.NewFieldElement(105, 223)
    x2, _ :=  lib.NewFieldElement(17, 223)
    y2, _ :=  lib.NewFieldElement(56, 223)


    result, _ := lib.NewPoint(x1, y1, a, b)
    result.Print()


    result2, _ := lib.NewPoint(x2, y2, a, b)
    result2.Print()


    _, err  := lib.PointAdd(result, result2)

    fmt.Println("Error was ", err)


}
