package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}
func(r13 *rot13Reader) Read(p []byte) (n int, err error){

    // Readerから読み出す
    n, err = r13.r.Read(p)
    for i := 0; i < n; i++ {
        // アルファベットか判定
        if 'A' <= p[i] && p[i] <= 'z' {
            // 大文字/小文字判定
            var cStart, cEnd byte
            if 'A' <= p[i] && p[i] <= 'Z' {
                cStart = 'A'
                cEnd = 'Z'
            }else{
                cStart = 'a'
                cEnd = 'z'
            }
            // Rot13変換をかける
            c13 := p[i] + 13
            if cEnd < c13 {
                p[i]= cStart + (c13 - cEnd) - 1
            }else{
                p[i]=c13
            }
        }
    }

    return
}

func main() {
    s := strings.NewReader(
        "Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}