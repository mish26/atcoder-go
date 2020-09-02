package main

import (
    "fmt"
    "math"
)

func main() {
    var n int
    fmt.Scan(&n)

    var sqrtN int = int(math.Ceil(math.Sqrt(float64(n))))
//    fmt.Println(sqrtN)

    for i := 2; i <= sqrtN; i++  {
        if n % i == 0 {

            fmt.Println("NO")
            return
        }
    }

    fmt.Println("YES")
}

