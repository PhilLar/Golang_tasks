package sqrt_default

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	var old float64
	for {
		fmt.Println(z)
		old = z
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-old) <= 0.0001 {
			break
		}
	}
	return z
}
