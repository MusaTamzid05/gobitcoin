package lib

import (
    "testing"
    )

type ThisIsSick struct {
    x int
    y int
}

func TestECC(t *testing.T) {
    prime := 223
    a, _  := NewFieldElement(0, prime)
    b, _  := NewFieldElement(7, prime)


    var err error

    valid := []ThisIsSick {
        ThisIsSick{x: 192, y: 105},
        ThisIsSick{x: 17, y: 56},
        ThisIsSick{x: 1, y: 193},
    }

    for _, sick := range valid {

        x, _  := NewFieldElement(sick.x, prime)
        y, _  := NewFieldElement(sick.y, prime)

        _, err = NewPoint(x, y, a, b)

        if err != nil {
            t.Errorf("%s", err)
        }
    }


    inValids := []ThisIsSick {
        ThisIsSick{x: 200, y: 109},
        ThisIsSick{x: 42, y: 99},
    }

    for _, sick := range inValids{

        x, _  := NewFieldElement(sick.x, prime)
        y, _  := NewFieldElement(sick.y, prime)

        _, err = NewPoint(x, y, a, b)

        if err == nil {
            t.Errorf("Invalid point validated")
        }
    }


}
