package main

import (
	"bufio"
	"fmt"
	"os"
	"math"
	"strings"
)

const (
	MOD = 998244353
	inf = 1 << 60
)

var (
	in  *bufio.Reader
	out *bufio.Writer
)

// n := II()
func II() int {
    n := 0
    neg := false
    b, _ := in.ReadByte()

    for b < '0' || b > '9' {
        if b == '-' {
            neg = true
            b, _ = in.ReadByte()
            break
        }
        b, _ = in.ReadByte()
    }

    for b >= '0' && b <= '9' {
        n = n*10 + int(b&15)
        b, _ = in.ReadByte()
    }
    if neg {
        n = -n
    }
    return n
}

// n := IF()
func IF() (x float64) {
	fmt.Fscan(in, &x)
	return
}

// s := IS()
func IS() (x string) {
	fmt.Fscan(in, &x)
	return
}

func Print(a ...any)            { fmt.Fprint(out, a...) }
func Println(a ...any)          { fmt.Fprintln(out, a...) }
func Printf(f string, a ...any) { fmt.Fprintf(out, f, a...) }

func main() {
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	T := 1

	for t := 0; t < T; t++ {
		solve()
	}
}

func abs[T int | int64 | float64](x T) T {
	if x >= 0 { return x }
	return -x
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// 数组要手动切片 sum(nums[:])
func sum[T int | int64 | float64](nums []T) T {
	var total T
	for _, v := range nums {
		total += v
	}
	return total
}

func isqrt(x int) int {
	rt := int(math.Sqrt(float64(x))) // 可能会算多一点点
	if rt*rt > x {
		rt--
	}
	return rt
}

func ReverseString(s string) string {
    var b strings.Builder
    b.Grow(len(s))
    for i := len(s) - 1; i >= 0; i-- {
        b.WriteByte(s[i])
    }
    return b.String()
}

func accumulate(nums []int) []int {
	n := len(nums)
	s := make([]int, n + 1)
	for i, x := range nums {
		s[i + 1] = s[i] + x
	}
	return s
}

// ----------------- io 模板结束 ----------------- //
// ----------------- 主函数部分  ----------------- //

// https://codeforces.com/problemset/problem/1265/E

var inv = qpow(100, -1, MOD)

func solve() {
	n := II()
	p := make([]int, n)
	for i := range p {
		p[i] = II()
		p[i] = p[i] * inv % MOD
	}

	a := make([]int, n + 1)
	b := make([]int, n + 1)
	for i := n - 1; i >= 0; i-- {
		a[i] = (p[i] * a[i + 1] + 1 - p[i]) % MOD
		b[i] = (p[i] * b[i + 1] + 1) % MOD
	}

	Println(b[0] * qpow(1 - a[0] + MOD, -1, MOD) % MOD)
}

func qpow(x, n, mod int) int {
	ans := 1
	base := x
	exp := n

	if mod > 0 {
		base %= mod
	}

	if exp < 0 {
		if mod == 0 {
			return 0
		}
		// 模质数下求逆元（费马小定理）
		inv, b, p := 1, base, mod-2
		for p > 0 {
			if p&1 == 1 {
				inv = (inv * b) % mod
			}
			b = (b * b) % mod
			p >>= 1
		}
		base = inv
		exp = -exp
	}

	for exp > 0 {
		if exp&1 == 1 {
			ans *= base
			if mod > 0 {
				ans %= mod
			}
		}
		base *= base
		if mod > 0 {
			base %= mod
		}
		exp >>= 1
	}
	return ans
}
