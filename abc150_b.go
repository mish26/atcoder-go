package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Scan(&n)
    var s string
    fmt.Scan(&s)

    var result string = "ABC"

//    var slice = strings.Split(s, "")

    var j int = 0
    for i := 0; i < n-3; i++  {
        var char string = s[i:i+3]
        if char == result {
            j++
        }
    }
    fmt.Println(j)

}

