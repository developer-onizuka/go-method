package main
import (
	"fmt"
	"mathlib"
	_"math"
)

type LocalSquare mathlib.Square
func (s *LocalSquare) Area() float64 {
	//return (*s).Width * (*s).Height
	return s.Width * s.Height
}
// https://qiita.com/tatachiy/items/7856931f4cacf2c5b743

func main() {
	fmt.Printf("----------\n")
	s := &mathlib.Square{ 1.0, 1.0 }
	t := &LocalSquare{ 1.0, 1.0 }
	fmt.Printf("Value is %v and %v\n", s, t)
	//fmt.Printf("  (*s).Width: %v (*s).Height: %v\n", (*s).Width, (*s).Height)
	//fmt.Printf("  (*t).Width: %v (*t).Height: %v\n", (*t).Width, (*t).Height)
	fmt.Printf("  s.Width: %v    s.Height: %v\n", s.Width, s.Height)
	fmt.Printf("  t.Width: %v    t.Height: %v\n", t.Width, t.Height)
	fmt.Printf("Type is %T and %T\n", s, t)
	//fmt.Printf("  (*s).Width: %T (*s).Height: %T\n", (*s).Width, (*s).Height)
	//fmt.Printf("  (*t).Width: %T (*t).Height: %T\n", (*t).Width, (*t).Height)
	fmt.Printf("  s.Width: %T    s.Height: %T\n", s.Width, s.Height)
	fmt.Printf("  t.Width: %T    t.Height: %T\n", t.Width, t.Height)
	fmt.Printf("Address is %p and %p\n", s, t)
	//fmt.Printf("  (*s).Width: %p (*s).Height: %p\n", &(*s).Width, &(*s).Height)
	//fmt.Printf("  (*t).Width: %p (*t).Height: %p\n", &(*t).Width, &(*t).Height)
	fmt.Printf("  &s.Width: %p    &s.Height: %p\n", &s.Width, &s.Height)
	fmt.Printf("  &t.Width: %p    &t.Height: %p\n", &t.Width, &t.Height)
	fmt.Printf("\n")
	//fmt.Printf("Sqrt is %v\n", (*s).Sqrt())
	fmt.Printf("s.Sqrt is %v\n", s.Sqrt())
	//fmt.Printf("Area is %v\n", (*t).Area())
	fmt.Printf("t.Area is %v\n", t.Area())
	fmt.Printf("----------\n")
	// https://qiita.com/tikidunpon/items/2d9598f33817a6e99860

	var u mathlib.Square
	u = *s
	u.Width = 2.0
	var v LocalSquare
	v = *t
	v.Width = 2.0
	fmt.Printf("Value is %v and %v\n", u, v)
	fmt.Printf("Type is %T and %T\n", u, v)
	fmt.Printf("Address is %p and %p\n", &u, &v)
	fmt.Printf("\n")
	//fmt.Printf("Sqrt is %v\n", u.Sqrt())
	fmt.Printf("(&u).Sqrt is %v\n", (&u).Sqrt()) //2.2360679
	fmt.Printf("s.Sqrt is %v\n", s.Sqrt())       //1.4142135
	//fmt.Printf("Area is %v\n", v.Area())
	fmt.Printf("(&u).Area is %v\n", (&v).Area()) //2.0
	fmt.Printf("t.Area is %v\n", t.Area())       //1.0
	fmt.Printf("----------\n")
	// https://qiita.com/tikidunpon/items/2d9598f33817a6e99860

	var w *mathlib.Square
	w = s
	w.Width = 2.0
	var x *LocalSquare
	x = t
	x.Width = 2.0
	fmt.Printf("Value is %v and %v\n", *w, *x)
	fmt.Printf("Type is %T and %T\n", w, x)
	fmt.Printf("Address is %p and %p\n", w, x)
	fmt.Printf("\n")
	//fmt.Printf("Sqrt is %v\n", (*w).Sqrt())
	fmt.Printf("w.Sqrt is %v\n", w.Sqrt()) //2.2360679
	fmt.Printf("s.Sqrt is %v\n", s.Sqrt()) //2.2360679
	//fmt.Printf("Area is %v\n", (*x).Area())
	fmt.Printf("x.Area is %v\n", x.Area()) //2.0
	fmt.Printf("t.Area is %v\n", t.Area()) //2.0
	fmt.Printf("----------\n")
}

