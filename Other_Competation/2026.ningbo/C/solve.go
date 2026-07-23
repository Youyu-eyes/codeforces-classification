package main

func solve() {
	n, x := II(), II()
	f := make([]int, 12 * n + 1)
	f[0] = 1
	for i := 1; i <= n; i++ {
		nf := make([]int, 12 * n + 1)
		k := II()
		invK := qpow(k, MOD - 2, MOD)
		points := make([]int, k)
		for j := range points {
			points[j] = II()
		}

		for _, v := range points {
			for j := v; j <= 12 * i; j++ {
				nf[j] = (nf[j] + f[j - v] * invK) % MOD
			}
		}
		copy(f, nf)
	}

	ans := 0
	for i := x; i <= 12 * n; i++ {
		ans = (ans + f[i]) % MOD
	}
	Println(ans)
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
