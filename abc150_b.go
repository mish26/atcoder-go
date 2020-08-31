package main

import (
    "fmt"
    "strings"
)

func main() {
    var n int = 1
  //  fmt.Scan(&n)
    var s string = "test"
  //  fmt.Scan(&s)

    var result string = "ABC"

    var slice = strings.Split(s, "")
    fmt.Print(slice)

    for i := 0; i < n; i++  {
        fmt.Println(slice[i:i+3])
    }

}

