// https://codeforces.com/problemset/problem/467/C
// 原题难度 1700
// WQS 二分难度估计 2100

package main

type pair struct{ b, x int }

// p 比 q 更劣
func less(p, q pair) bool {
	return p.b < q.b || p.b == q.b && p.x > q.x
}

func solve() {
	n, m, limit := II(), II(), II()
	nums := make([]int, n)
	s := make([]int, n + 1)
	for i := range n {
		nums[i] = II()
		s[i + 1] = s[i] + nums[i]
	}

	dp := func(k int) pair {
		f := make([]pair, n + 1)
		for i := m; i <= n; i++ {
			cand1 := pair{f[i - m].b + s[i] - s[i - m] - k, f[i - m].x + 1}
			cand2 := f[i - 1]
			if less(cand1, cand2) {
				f[i] = cand2
			} else {
				f[i] = cand1
			}
		}
		return f[n]
	}

	var res pair

	maxSum := 0
	for i := m; i <= n; i++ {
		maxSum = max(maxSum, s[i]-s[i-m])
	}

	ans := 0
	left, right := 0, maxSum + 1
	for left + 1 < right {
		mid := left + (right - left) >> 1
		res = dp(mid)
		if res.x <= limit {
			ans = mid * limit + res.b
			right = mid
		} else {
			left = mid
		}
	}
	Println(ans)
}
