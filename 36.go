package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
    ret := make([][]uint8, dy)
    for i := 0; i < len(ret); i++ {
        child := make([]uint8, dx)
        for t := 0; t < len(child); t++ {
            child[t] = uint8(i*t)
        }
      ret[i] = child
    }

    return ret
}

func main() {
    pic.Show(Pic)
}
