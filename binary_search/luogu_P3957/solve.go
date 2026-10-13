// https://www.luogu.com.cn/problem/P3957
// 二分答案 + 单调队列优化 DP

// 二分 g，然后状态转移方程为
// f[i] = max_j f[j] + score[i]
// j 需要满足
// d - g <= dis[i] - dis[j] <= d + g
// dis[i] - (d + g) <= dis[j] <= dis[i] - max(1, d - g)
// 由于 dis 单调，则左右端点单调，可以用单调队列优化 DP
// 由于 j 单调，因此可以用双指针优化二分

// 本题注意的细节是，需要提前判断是否可以到达
// 必须中途退出，我们定义的是跳跃到 f[i] 可以获得的最大得分
// 本题 f[n-1] 不一定可以跳到，也就是说 f[i] > f[n-1] 是可能的

// 时间复杂度：O(nlogMX)，其中 MX = dis[n-1]

package main

func solve() {
	n, d, k := II() + 1, II(), II()
	dis := make([]int, n)
	score := make([]int, n)
	s := 0
	for i := 1; i < n; i++ {
		dis[i] = II()
		score[i] = II()
		if score[i] > 0 {
			s += score[i]
		}
	}

	if s < k {
		Println(-1)
		return
	}

	check := func(g int) bool {
		f := make([]int, n)
		for i := range f {
			f[i] = -inf
		}
		f[0] = 0
		
		q := []int{}
		r := 0
		for i := 1; i < n; i++ {
			for ; r < i && dis[r] <= dis[i] - max(1, d - g); r++ {
				for len(q) > 0 && f[q[len(q)-1]] <= f[r] {
					q = q[:len(q)-1]
				}
				q = append(q, r)
			}

			for len(q) > 0 && dis[q[0]] < dis[i] - (d + g) {
				q = q[1:]
			}

			if len(q) > 0 {
				f[i] = f[q[0]] + score[i]
				if f[i] >= k {
					return true
				}
			}
		}

		return false
	}

	left, right := -1, dis[n-1]
	for left + 1 < right {
		mid := (left + right) >> 1
		if check(mid) {
			right = mid
		} else {
			left = mid
		}
	}
	Println(right)
}
