package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {

    cnt := 0
    num1 := 0
    num2 := 1

    return func() int {
        var ret int
        if cnt == 0{
            ret = num1
        }else if cnt == 1 {
            ret = num2
        }else{
            ret = num1 + num2
            num1 = num2
            num2 = ret
        }
        cnt++
        return ret
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}