package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Scan(&n)

    // 整数Nが与えられます。N!の正の約数の個数を109+7で割った余りを求めてください。
    // 約数の個数を求めるために、素因数分解を行う

    var nFact int =  factorial(n)
    fmt.Println(nFact)
    var exs []int = make([]int, 0)

    for i := 2; i * i <= n; i++  {
        if nFact % i != 0 {
            continue
        }

        var ex int = 0
        for {
            // 指数をカウントするために割り切れなくなるまで割る
            if nFact % i == 0 {
                ex++
                nFact = nFact / i
                fmt.Println("nFact:", nFact, ", i:", i)
            } else {
                break
            }
        }
        // 指数を配列に格納し、ループ完了後に取り出して掛け算する
        exs = append(exs, ex)
    }
    if nFact != 0 {
        exs = append(exs, 1)
    }
    var divisorCount int = 1
    fmt.Println("exs:", exs)
    for _, v := range exs{
        divisorCount = (v+1) * divisorCount
    }
    fmt.Println(divisorCount % 1000000007)
}

func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n - 1)
}
