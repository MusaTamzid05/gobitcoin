package lib

import (
    "math"
    "errors"
)

type Point struct {
    a int
    b int
    x int
    y int
}

func NewPoint(x, y, a, b int)  (*Point, error) {

    
    if math.Pow(float64(y), 2.0) != math.Pow(float64(x), 3) + (float64(a) * float64(x)) + float64(b) {
        return &Point{}, errors.New("Is not on the curve")
    }
    return &Point {
        a: a,
        b: b,
        x: x,
        y: y,
    }, nil
}


func PointEqual(p1, p2 *Point) bool {
    return p1.x == p2.x &&
           p1.y == p2.y &&
           p1.a == p2.a &&
           p1.b == p2.b 

}
func PointAdd(p1, p2 *Point) (*Point, error) {
    if (p1.a != p2.a ) || (p1.b != p2.b) {
        return &Point{}, errors.New("Is not on the curve")
    }

    if p1.x == -1 {
        return p2, nil
    }

    if p2.x == -1 {
        return p1, nil
    }

    if PointEqual(p1, p2) && p1.y == 0 * p1.x {
        return &Point{-1, -1, p1.a, p1.b}, nil

    }

    return &Point{}, errors.New("Add error")
}

