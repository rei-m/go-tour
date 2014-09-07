package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {

    ret := make(map[string]int)
    slice := strings.Fields(s)
    for i := 0; i < len(slice); i++ {
        v, j := ret[slice[i]]
        if j {
          ret[slice[i]] += v
        }else{
          ret[slice[i]] = 1
        }
    }
    return ret
}

func main() {
   wc.Test(WordCount)
}