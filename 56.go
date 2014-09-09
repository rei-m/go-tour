package main

import (
    "fmt"
    "math"
)

type ErrNegativeSqrt struct {
    Value float64
}
func (e *ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e.Value))
}

func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return 0, &ErrNegativeSqrt{
      f,
      }
    }else{
        z := float64(1)
        for i := 0; ; i++ {
            z1 := z
            z = z - (z * z - f) / 2 * z
            if math.Abs(z - z1) <= 0.001 {
                return z, nil
            }
        }
    }
    return 0, nil
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}