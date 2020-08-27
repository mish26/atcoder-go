package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Scan(&n)

    for i := 0;i < 10; i++  {
        for j := 0;j < 10; j++  {
            var result int = i * j
            if result == n {
                fmt.Println("Yes")
                return
            }
        }
    }
    fmt.Println("No")
}

