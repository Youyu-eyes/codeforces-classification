// https://codeforces.com/problemset/problem/1265/E

package main

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
