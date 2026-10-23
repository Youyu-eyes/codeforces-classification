// https://codeforces.com/problemset/problem/453/A
// 1600
// 概率论，期望

// 投掷一次，点数小于等于 k 的概率为 k / m
// 投掷一次，点数恰好等于 k 的概率为 小于等于 k 的概率 - 小于等于 k - 1 的概率
// k / m - (k - 1) / m
// 枚举 k 为最大点数，则期望为

// sum{k = [1, m]}(k / m - (k - 1) / m) ^ n

package main
import (
	"math"
)

func solve() {
    m := II()
    n := II()

    if m == 1 {
        Println("1.000000000000")
        return
    }

    mf := float64(m)
    nf := float64(n)
    var sum float64

    for k := 1; k < m; k++ {
        ratio := float64(k) / mf
        sum += math.Pow(ratio, nf)
    }

    ans := mf - sum
    Printf("%.12f\n", ans)
}