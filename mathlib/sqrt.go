package mathlib
import (
	_"fmt"
	"math"
)

type Square struct {
	Width float64
	Height float64
}

func (s *Square) Sqrt() float64 {
	//return math.Sqrt((*s).Width*(*s).Width + (*s).Height*(*s).Height)
	return math.Sqrt((s.Width)*(s.Width) + (s.Height)*(s.Height))
}

