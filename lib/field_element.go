package lib


import (
    "errors"
    "math"
)

type FieldElement struct {
    num int
    prime int
}

func NewFieldElement(num, prime int) (*FieldElement, error){

    if num >= prime || num < 0  {
        return nil, errors.New("Invalid!!")
    }
    return &FieldElement{num: num, prime: prime}, nil
}

func (f FieldElement) Add(other *FieldElement) ( *FieldElement, error) {
    if f.prime != other.prime {
        return nil, errors.New("Cannot add two numbere in different field")
    }

    return &FieldElement{
        num: (f.num + other.num) % f.prime,
        prime: f.prime,
    }, nil
}

func (f FieldElement) Sub(other *FieldElement) ( *FieldElement, error) {
    if f.prime != other.prime {
        return nil, errors.New("Cannot add two numbere in different field")
    }

    return &FieldElement{
        num: (f.num - other.num) % f.prime,
        prime: f.prime,
    }, nil
}


func (f FieldElement) Mul(other *FieldElement) ( *FieldElement, error) {
    if f.prime != other.prime {
        return nil, errors.New("Cannot add two numbere in different field")
    }

    return &FieldElement{
        num: (f.num * other.num) % f.prime,
        prime: f.prime,
    }, nil
}


func (f FieldElement) Pow(exponent int) ( *FieldElement, error) {
    n := exponent 

    for n < 0 {
        n += f.prime - 1
    }

    num := math.Pow(float64(f.num), float64(n))

    return &FieldElement{
        num: int(num)  % f.prime,
        prime: f.prime,
    }, nil
}


func (f FieldElement) Equal(other* FieldElement) bool {
    return f.num == other.num && f.prime == other.prime
}
