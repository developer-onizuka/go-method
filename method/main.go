package main
import (
	"fmt"
	"mathlib"
	"math"
)

type LocalSquare mathlib.Square
func (s *LocalSquare) Area() float64 {
	return math.Sqrt((*s).Width * (*s).Height)
}

func main() {
	fmt.Printf("----------\n")
	s := &mathlib.Square{ 1.0, 1.0 }
	t := &LocalSquare{ 1.0, 1.0 }
	fmt.Printf("Value is %v and %v\n", *s, *t)
	fmt.Printf("Type is %T and %T\n", s, t)
	fmt.Printf("Address is %p and %p\n", s, t)
	fmt.Printf("Sqrt is %v\n", (*s).Sqrt())
	fmt.Printf("Area is %v\n", (*t).Area())
	fmt.Printf("----------\n")

	var u mathlib.Square
	u = *s
	var v LocalSquare
	v = *t
	fmt.Printf("Value is %v and %v\n", u, v)
	fmt.Printf("Type is %T and %T\n", u, v)
	fmt.Printf("Address is %p and %p\n", &u, &v)
	fmt.Printf("Sqrt is %v\n", (u).Sqrt())
	fmt.Printf("Area is %v\n", (v).Area())
	fmt.Printf("----------\n")

	var w *mathlib.Square
	w = s
	var x *LocalSquare
	x = t
	fmt.Printf("Value is %v and %v\n", *w, *x)
	fmt.Printf("Type is %T and %T\n", w, x)
	fmt.Printf("Address is %p and %p\n", w, x)
	fmt.Printf("Sqrt is %v\n", (*w).Sqrt())
	fmt.Printf("Area is %v\n", (*x).Area())
	fmt.Printf("----------\n")
}

