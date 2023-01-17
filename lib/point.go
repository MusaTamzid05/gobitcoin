package lib

import (
    "math"
    "errors"
)

type PointInt struct {
    i int
}

type Point struct {
    a int
    b int
    x *PointInt
    y *PointInt
}

func NewPoint(x, y, a, b int)  (*Point, error) {
    if math.Pow(float64(y), 2.0) != math.Pow(float64(x), 3) + (float64(a) * float64(x)) + float64(b) {
        return &Point{}, errors.New("Is not on the curve")
    }
    return &Point {
        a: a,
        b: b,
        x: &PointInt{i: x},
        y: &PointInt{i: y},
    }, nil
}

func NewInfinityPoint(a, b int)  *Point {
    return &Point {
        a: a,
        b: b,
        x: nil,
        y: nil,
    }
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

    if p1.x == nil {
        return p2, nil
    }

    if p2.x == nil {
        return p1, nil
    }

    if PointEqual(p1, p2) && p1.y.i == 0 * p1.x.i {
        return NewInfinityPoint(p1.a, p1.b), nil

    }

    return &Point{}, errors.New("Add error")
}

