// BM + kitamasa

// 时间复杂度：O(k^2 · logn)
// 空间复杂度：O(k · logn)

package main

import (
	"slices"
)

func solveInit(p []int) []int {
	x := 8
	f := make([]int, x + 3)
	inv := qpow(100 - p[0], -1, MOD)

	for i := x - 1; i >= 0; i-- {
		f[i] = 100
		for j := 1; j <= 3; j++ {
			f[i] = (f[i] + p[j] * f[i + j]) % MOD
		}
		f[i] = f[i] * inv % MOD
	}
	slices.Reverse(f)
	return f
}

func solveBM() {
	x, p1, p2, p3, p4 := II(), II(), II(), II(), II()
	p := [4]int{p4, p3, p2, p1}
	a := solveInit(p[:])
	coef := berlekampMassey(a)
	slices.Reverse(coef)
	ans := kitamasa(coef, a, x + 2)
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
