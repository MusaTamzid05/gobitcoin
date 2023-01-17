package lib

import (
    "fmt"
    "errors"
)


type Point struct {
    a *FieldElement
    b *FieldElement
    x *FieldElement
    y *FieldElement
}

func NewPoint(x, y, a, b *FieldElement)  (*Point, error){
    var err error
    _y, err := y.Pow(2)

    if err != nil {
        return nil, err
    }

    _x, err:= x.Pow(3)

    if err != nil {
        return nil, err
    }

    ax, err  := a.Mul(x)

    if err != nil {
        return nil, err
    }

    x_ax, err  := _x.Add(ax)

    if err != nil {
        return nil, err
    }


    x_ax_b, err:= x_ax.Add(b)

    if err != nil {
        return nil, err
    }

    if !_y.Equal(x_ax_b)  {
        return nil, errors.New("Is not on the curve")
    }

    if err != nil {
        return nil, err
    }

    return &Point {
        a: a,
        b: b,
        x: x,
        y: y,
    } , nil
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

func PointAdd(p1, p2 *Point) (*Point, error){

    e1 := p1.a.Equal(p2.a)
    e2 := p1.b.Equal(p2.b)

    if !e1 || !e2 {
        return nil, errors.New("Is not on the curve")
    }

    if p1.x == nil {
        return p2, nil
    }

    if p2.x == nil {
        return p1, nil
    }

    zero,_ := NewFieldElement(0, p1.x.prime) // dont know whether this is right
    mul, err  :=  zero.Mul(p1.x)

    if err != nil {
        return nil, err
    }



    if PointEqual(p1, p2) && p1.y.Equal(mul) {
        return NewPoint(nil, nil,p1.a, p1.b)
    }

    if !p1.x.Equal(p2.x) {
        diffY , _ := p2.y.Sub(p1.y)
        diffX , _ := p2.x.Sub(p1.x)
        s, _ := diffY.Div(diffX)

        xpart1, _ := s.Pow(2)
        xpart2, _ := xpart1.Sub(xpart1)
        x, _ := xpart2.Sub(p2.y)
        ypart2, _ := p1.x.Sub(x)
        ypart1, _ := s.Mul(ypart2)
        y, _ := ypart1.Sub(p1.y)

        return NewPoint(x, y,p1.a, p1.b)

    }


    return nil, nil
}


