package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {

    ret := make(map[string]int)
    slice := strings.Fields(s)
    for _, v := range slice {
        ret[v] += 1
    }
    return ret
}

func main() {
   wc.Test(WordCount)
}