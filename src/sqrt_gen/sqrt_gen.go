package sqrt_gen

import (
  //"fmt"
  "math"
  )

func Sqrt(x float64, n int) chan float64 {
	ch := make(chan float64)
	
	go func () {
    z := 1.0
		for i := 0; i < n; i++ {
			ch <- z
			if math.Floor(z*10000)/10000 == math.Floor((z-(z*z - x) / (2*z))*10000)/10000 {
					break
				}
      		z -= (z*z - x) / (2*z)
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