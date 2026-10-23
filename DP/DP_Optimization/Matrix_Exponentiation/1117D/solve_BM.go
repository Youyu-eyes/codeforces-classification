// https://codeforces.com/problemset/problem/1117/D
// 2100
// BM 优化 DP

// 状态转移方程见 solve
// 通过一维 DP 打表可以预处理出数列
// 通过前 2m 项计算通向公式

// 然后通过 kitamasa 算法快速计算第 n 项

// 时间复杂度：O(m^2 logn)
// 空间复杂度：O(m)

package main

import(
	"slices"
)

func solveInit(m int) []int {
	f := make([]int, 2 * m)
	for i := 0; i < m; i++ {
		f[i] = 1
	}
	for i := m; i < 2 * m; i++ {
		f[i] = (f[i - 1] + f[i - m]) % MOD
	}
	return f
}

func solveBM() {
	n, m := II(), II()
	a := solveInit(m)

	coef := berlekampMassey(a)
	slices.Reverse(coef)
	ans := kitamasa(coef, a, n)
	Println(ans)
}

// ------- bM 算法：推导数列通项 ------- //
func berlekampMassey(a []int) (coef []int) {
	var preC []int
	preI, preD := -1, 0
	for i, v := range a {
		d := v
		for j, c := range coef {
			d = (d - c*a[i-1-j]) % MOD
		}
		if d == 0 {
			continue
		}

		if preI < 0 {
			coef = make([]int, i+1)
			preI, preD = i, d
			continue
		}

		bias := i - preI
		oldLen := len(coef)
		newLen := bias + len(preC)
		var tmp []int
		if newLen > oldLen {
			tmp = slices.Clone(coef)
			coef = slices.Grow(coef, newLen-oldLen)[:newLen]
		}

		delta := d * qpow(preD, -1, MOD) % MOD
		coef[bias-1] = (coef[bias-1] + delta) % MOD
		for j, c := range preC {
			coef[bias+j] = (coef[bias+j] - delta*c) % MOD
		}

		if newLen > oldLen {
			preC = tmp
			preI, preD = i, d
		}
	}

	for i, c := range coef {
		if c < -MOD / 2 {
			c += MOD
		} else if c > MOD / 2 {
			c -= MOD
		}
		coef[i] = c
	}

	return
}


// ------- kitamasa：快速计算第 n 项 ------- //

func kitamasa(coef, a []int, n int) (ans int) {
	defer func() { ans = (ans % MOD + MOD) % MOD }()
	if n < len(a) {
		return a[n]
	}

	k := len(coef)
	if k == 0 {
		return
	}
	if k == 1 {
		return a[0] * qpow(coef[0], n, MOD) % MOD
	}

	// 已知 f(n) 的各项系数为 a，f(m) 的各项系数为 b
	// 计算并返回 f(n+m) 的各项系数 c
	compose := func(a, b []int) []int {
		c := make([]int, k)
		for _, v := range a {
			// 累加 a[i] * f(m+i) 的各项系数
			for j, w := range b {
				c[j] = (c[j] + v*w) % MOD
			}
			// 从 f(m+i) 到 f(m+i+1)
			bk1 := b[k-1]
			for j := k - 1; j > 0; j-- {
				b[j] = (b[j-1] + bk1*coef[j]) % MOD
			}
			b[0] = bk1 * coef[0] % MOD
		}
		return c
	}

	// 计算 resC，以表出 f(n) = resC[k-1] * a[k-1] + resC[k-2] * a[k-2] + ... + resC[0] * a[0]
	resC := make([]int, k)
	resC[0] = 1
	c := make([]int, k)
	c[1] = 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			resC = compose(c, resC)
		}
		c = compose(c, slices.Clone(c))
	}

	for i, c := range resC {
		ans = (ans + c*a[i]) % MOD
	}
	return
}

// ------- 快速幂模板，支持负指数计算逆元 ------- //
// 必须要传 MOD 参数，如果不取模传 MOD = 0

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
