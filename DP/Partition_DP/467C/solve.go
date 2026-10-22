// https://codeforces.com/problemset/problem/467/C
// 1700
// 划分型 DP
// 有复杂度更优的 WQS 二分做法，见 DP 优化题单

package main

func solve() {
	n, m, k := II(), II(), II()
	nums := make([]int, n)
	s := make([]int, n + 1)
	for i := range n {
		nums[i] = II()
		s[i + 1] = s[i] + nums[i]
	}

	f := make([]int, n + 1)
	for K := 1; K <= k; K++ {
		nf := make([]int, n + 1)
		for i := K * m; i <= n - (k - K) * m; i++ {
			nf[i] = max(nf[i - 1], f[i - m] + s[i] - s[i - m])
		}
		f = nf
	}
	Println(f[n])
}
