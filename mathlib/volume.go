package mathlib
import (
	"fmt"
	_"math"
)

type Xyz struct {
	X float64
	Y float64
	Z float64
}

func Volume(a Xyz) float64 {
	x, y, z := a.X, a.Y, a.Z
	a.X = y
	a.Y = z
	a.Z = x * 10
	fmt.Printf("-----\n")
	fmt.Printf("a.X: %v, %p\n", a.X, &a.X)
	fmt.Printf("a.Y: %v, %p\n", a.Y, &a.Y)
	fmt.Printf("a.Z: %v, %p\n", a.Z, &a.Z)
	fmt.Printf("-------------------->")
	return a.X*a.Y*a.Z
}

func Volumep(a *Xyz) float64 {
	x, y, z := a.X, a.Y, a.Z
	a.X = y
	a.Y = z
	a.Z = x * 10
	fmt.Printf("-----\n")
	fmt.Printf("(*a).X: %v, %p\n", (*a).X, &(*a).X)
	fmt.Printf("(*a).Y: %v, %p\n", (*a).Y, &(*a).Y)
	fmt.Printf("(*a).Z: %v, %p\n", (*a).Z, &(*a).Z)
	fmt.Printf("-------------------->")
	return a.X*a.Y*a.Z
}
