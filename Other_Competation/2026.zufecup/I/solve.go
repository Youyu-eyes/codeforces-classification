// https://codeforces.com/gym/680541/attachments
// 期望 DP
// 定义 f[i] 从 i 颗星开始超过 x 颗星的最小天数
// f[i] = (f[i] = (100 + \sum_{j = 1}^{3} (p[j] * f[i + j])) * inv % MOD)
// inv 是 (100 - p[0]) 在 MOD 意义下的乘法逆元
// 初始化 f[x], f[x + 1], f[x + 2] = 0
// 答案为 f[0]

// 时间复杂度：O(k · x)，k 表示加星的可能数量，用前缀和可优化至 O(x)
// 空间复杂度：O(x) 空间优化可以做到 O(1)

// 更快的做法见 solve_matrix.go

package main

func solve() {
	x, p1, p2, p3, p4 := II(), II(), II(), II(), II()
	p := [4]int{p4, p3, p2, p1}
	f := make([]int, x + 3)
	inv := qpow(100 - p[0], -1, MOD)

	for i := x - 1; i >= 0; i-- {
		f[i] = 100
		for j := 1; j <= 3; j++ {
			f[i] = (f[i] + p[j] * f[i + j]) % MOD
		}
		f[i] = f[i] * inv % MOD
	}
	Println(f[0])
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
