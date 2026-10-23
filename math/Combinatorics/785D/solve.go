// https://codeforces.com/problemset/problem/785/D
// 2300

// 见 solution.md

package main

var (
	F          []int
	INV_F      []int
)

func init() {
	F = make([]int, MX)
	INV_F = make([]int, MX)

	F[0] = 1
	for i := 1; i < MX; i++ {
		F[i] = F[i-1] * i % MOD
	}
	INV_F[MX-1] = qpow(F[MX-1], MOD-2)
	for i := MX - 1; i > 0; i-- {
		INV_F[i-1] = INV_F[i] * i % MOD
	}
}

func solve() {
	s := IS()
	A, B := 0, 0
	for _, c := range s {
		if c == ')' {
			B++
		}
	}

	ans := 0
	for _, c := range s {
		if c == '(' {
			A++
		} else {
			B--
			ans = (ans + comb(A+B, A-1)) % MOD
		}
	}

	Println(ans)
}

func qpow(x, n int) int {
	ans := 1
	x %= MOD

	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			ans = ans * x % MOD
		}
		x = x * x % MOD
	}
	return ans
}

// ------- 组合数学 ------- //

// 从 n 个数中选 m 个数的方案数
func comb(n, m int) int {
	if m < 0 || m > n {
		return 0
	}
	return F[n] * INV_F[m] % MOD * INV_F[n-m] % MOD
}
