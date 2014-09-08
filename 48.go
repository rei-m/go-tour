package main

import "fmt"
import "math/cmplx"


func Cbrt(x complex128) (z complex128) {
    z = complex128(1)
    for{
        z2 := ((z*z*z)-x)/(3*z*z)
        if cmplx.Abs(z2) < 0.000000001 {
            return z
        }
        z = z - z2
    }

  return z
}

func main() {
    fmt.Println(Cbrt(2))
    fmt.Println(cmplx.Pow(Cbrt(2), 3))
}
