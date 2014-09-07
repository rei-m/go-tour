package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    z := float64(1)
    for i := 0; ; i++ {
        z1 := z
        z = z - (z * z - x) / 2 * z
        if math.Abs(z - z1) <= 0.001 {
            return z
        }
    }
}

func main() {
    fmt.Println(Sqrt(2))
}
