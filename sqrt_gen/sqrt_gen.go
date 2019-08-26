package sqrt_gen

import (
	//"fmt"
	"math"
)

func Sqrt(x float64, n int) chan float64 {
	ch := make(chan float64)

	go func() {
		z := 1.0
		var old float64
		for {
			ch <- z
			old = z
			z -= (z*z - x) / (2 * z)
			if math.Abs(z-old) <= 0.0001 {
				break
			}
		}
		close(ch)
	}()

	return ch
}

// func main() {
// 	for i := range Sqrt(0.25, 10) {
// 		fmt.Println("Counted", i)
// 	}
// }
