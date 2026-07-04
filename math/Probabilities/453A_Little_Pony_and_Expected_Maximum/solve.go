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