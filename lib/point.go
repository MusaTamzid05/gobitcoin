package lib

import (
    "fmt"
)


type Point struct {
    a *FieldElement
    b *FieldElement
    x *FieldElement
    y *FieldElement
}

func NewPoint(x, y, a, b *FieldElement)  *Point {
    var err error
    _y, err := y.Pow(2)

    ErrorCheck(err)

    _x, err:= x.Pow(3)
    ErrorCheck(err)

    ax, err  := a.Mul(x)
    ErrorCheck(err)

    x_ax, err  := _x.Add(ax)
    ErrorCheck(err)

    x_ax_b, err:= x_ax.Add(b)
    ErrorCheck(err)

    if !_y.Equal(x_ax_b)  {
        panic("Is not on the curve")
    }

    return &Point {
        a: a,
        b: b,
        x: x,
        y: y,
    } 
}

func (p Point) Print() {
    fmt.Println("x = ", p.x, " y = ", p.y, " a = ", p.a, " b = ", p.b)
}


func PointEqual(p1, p2 *Point) bool {
    e1 := p1.x.Equal(p2.x)
    e2 := p1.y.Equal(p2.y)
    e3 := p1.a.Equal(p2.a)
    e4 := p1.b.Equal(p2.b)

    return e1 == true &&  e2 == true && e3 == true && e4 == true
}

func PointAdd(p1, p2 *Point) *Point {

    e1 := p1.a.Equal(p2.a)
    e2 := p1.b.Equal(p2.b)

    if !e1 || !e2 {
        panic("Is not on the curve")
    }

    if p1.x == nil {
        return p2
    }

    if p2.x == nil {
        return p1
    }

    zero,_ := NewFieldElement(0, p1.x.prime) // dont know whether this is right
    mul, _ :=  zero.Mul(p1.x)


    if PointEqual(p1, p2) && p1.y.Equal(mul) {
        return NewPoint(p1.a, p1.b, nil, nil)
    }

    panic("Add error")

}


