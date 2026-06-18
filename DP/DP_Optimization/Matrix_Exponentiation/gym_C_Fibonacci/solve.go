// https://codeforces.com/gym/102644/problem/C
// 经典题，不予评分

// 求斐波那契数列第 n 项
// 可以用矩阵快速幂
// 这里为了节约时间，直接用 kitamasa 算法快速计算第 n 项
// 注意本题 f[0] = 0

// 时间复杂度：O(logn)
// 空间复杂度：O(1)

package main

import (
	"slices"
)

func solve() {
	n := II()
	coef := [2]int{1, 1}
	a := [2]int{1, 0}
	ans := kitamasa(coef[:], a[:], n + 1)
	Println(ans)
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
