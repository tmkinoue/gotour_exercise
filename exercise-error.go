package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	z := 1.0
	n := 0.0
	for {
        z -= (z*z - x) / (2 * z)
        if math.Abs(z-n) < 0.0000001 {
            return z, nil
        }
        n = z
    }

    return z, nil
}

func run(x float64) {
	x, err := Sqrt(x)
	if err != nil {
		fmt.Println(err);
	} else {
		fmt.Println(x)
	}
}

func main() {
	run(2)
	run(-2)
}

