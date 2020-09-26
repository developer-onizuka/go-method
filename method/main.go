package main
import (
	"fmt"
	"mathlib"
	"math"
)

type LocalSquare mathlib.Square
func (s *LocalSquare) Area() float64 {
	//return math.Sqrt((*s).Width * (*s).Height)
	return math.Sqrt(s.Width * s.Height)
}

func main() {
	fmt.Printf("----------\n")
	s := &mathlib.Square{ 1.0, 1.0 }
	t := &LocalSquare{ 1.0, 1.0 }
	fmt.Printf("Value is %v and %v\n", s, t)
	fmt.Printf("  (*s).Width: %v (*s).Height: %v\n", (*s).Width, (*s).Height)
	fmt.Printf("  (*t).Width: %v (*t).Height: %v\n", (*t).Width, (*t).Height)
	fmt.Printf("  s.Width: %v    s.Height: %v\n", s.Width, s.Height)
	fmt.Printf("  t.Width: %v    t.Height: %v\n", t.Width, t.Height)
	fmt.Printf("Type is %T and %T\n", s, t)
	fmt.Printf("  (*s).Width: %T (*s).Height: %T\n", (*s).Width, (*s).Height)
	fmt.Printf("  (*t).Width: %T (*t).Height: %T\n", (*t).Width, (*t).Height)
	fmt.Printf("  s.Width: %T    s.Height: %T\n", s.Width, s.Height)
	fmt.Printf("  t.Width: %T    t.Height: %T\n", t.Width, t.Height)
	fmt.Printf("Address is %p and %p\n", s, t)
	fmt.Printf("  (*s).Width: %p (*s).Height: %p\n", &(*s).Width, &(*s).Height)
	fmt.Printf("  (*t).Width: %p (*t).Height: %p\n", &(*t).Width, &(*t).Height)
	fmt.Printf("  s.Width: %p    s.Height: %p\n", &s.Width, &s.Height)
	fmt.Printf("  t.Width: %p    t.Height: %p\n", &t.Width, &t.Height)
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

