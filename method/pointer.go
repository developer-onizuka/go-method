package main

import (
	"fmt"
	"mathlib"
	"reflect"
)

func main() {
	fmt.Printf("----------\n")
	var ptr uintptr = reflect.ValueOf(mathlib.Volume).Pointer()
	fmt.Printf("Volume: 0x%x\n", ptr)
	var ptrp uintptr = reflect.ValueOf(mathlib.Volumep).Pointer()
	fmt.Printf("Volumep: 0x%x\n", ptrp)
	fmt.Printf("----------\n")
	fmt.Printf("\n")
	p := mathlib.Xyz{1.0, 2.0, 3.0}
	fmt.Printf("Address:\n")
	fmt.Printf("p.X: %v, %p\n", p.X, &p.X)
	fmt.Printf("p.Y: %v, %p\n", p.Y, &p.Y)
	fmt.Printf("p.Z: %v, %p\n", p.Z, &p.Z)
	fmt.Printf("Volume: %v\n", mathlib.Volume(p))
	fmt.Printf("\n")
	pp := &mathlib.Xyz{3.0, 4.0, 5.0}
	fmt.Printf("Address:\n")
	fmt.Printf("(*pp).X: %v, %p\n", (*pp).X, &(*pp).X)
	fmt.Printf("(*pp).Y: %v, %p\n", (*pp).Y, &(*pp).Y)
	fmt.Printf("(*pp).Z: %v, %p\n", (*pp).Z, &(*pp).Z)
	fmt.Printf("Volume: %v\n", mathlib.Volumep(pp))
}
