package sqrt_gen_test

import (
	"testing"
	"sqrt_gen"
	"reflect"
	"math"
	//"errors"
	//"fmt"
)


func TestSqrt(t *testing.T) {
	tables := []struct {
			x		float64
			n 		int
			out		[]float64
	}{
			{2.0, 10, []float64{1, 1.5, 1.4166, 1.4142}},
			{25.0, 10, []float64{1, 13, 7.4615, 5.4060, 5.0152, 5.0000}},
			{1.0, 10, []float64{1}},
	}

	for _, table := range tables {
		res_arr := make([]float64, 0)
		for i := range sqrt_gen.Sqrt(table.x, table.n) {
			res_arr = append(res_arr, math.Floor(i*10000)/10000)
		}
		if !reflect.DeepEqual(res_arr, table.out) {
			t.Errorf("expected arr: %v, actual: %v",
					table.out, res_arr)
		}
		// if err != nil {

		// 	if err.Error() != table.err.Error() {
		// 		t.Errorf("expected error: %v, actual: %v",
		// 			table.err.Error(), err.Error())
		// 	}
		// } else {
	}
}

