package sqrt_default

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i:=0; i<10; i++ {
		fmt.Println(z)
		if math.Floor(z*1000)/1000 == math.Floor((z-(z*z - x) / (2*z))*1000)/1000 {
			return z
		}
		z -= (z*z - x) / (2*z)
	}
	return z
}