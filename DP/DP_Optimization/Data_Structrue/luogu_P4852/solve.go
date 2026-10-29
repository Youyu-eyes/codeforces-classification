// https://www.luogu.com.cn/problem/P4852
// 洛谷紫 - 省选/NOI-

// 单调队列优化 DP
// DP 输出具体方案

// 相比于本题的优化
// 本题的状态定义更值得研究

// 如果你定义成选了 k 个连抽，处理完前 i 个数，已经连续抽了 x 个单抽
// f[k][i][x]
// 虽然状态转移可以被优化为 O(1) 的，但是这个状态定义已经爆炸了

// 考虑每次连抽一次后连续单抽，这样的情况划分为一段
// 即定义 f[k][i] 为已经处理完前 i 个数，已经抽了 k 个连抽
// 假设上一次左端点为 j，则
// a[j:j+c] 为连抽，欧气值为 a[j]
// a[j+c:i] 为单抽，欧气值为 a[j+c]+...+a[i-1]，前缀和优化为 s[i] - s[j+c]

// f[k][i] = max_j f[k-1][j] + s[i] - s[j+c] + a[j]
// 其中，单抽不超过 d
// i - j - c <= d
// j >= i - c - d

// 至少一次连抽
// i - j >= c
// j <= i - c
// 上下限均与 i 有关，且单调递增，采用单调队列优化
// 如果不需要输出具体方案，则有更优的 WQS 二分做法，留给读者思考

// 时间复杂度：O(n(n*(c-1)+m))

package main

func solve() {
	n, m, c, d := II(), II(), II(), II()
	nums := make([]int, c*n+m)
	s := make([]int, c*n+m+1)
	for i := range c*n+m {
		nums[i] = II()
		s[i+1] = s[i] + nums[i]
	}
	f := make([]int, c*n+m+1)
	for i := range d+1 {
		f[i] = s[i]
	}
	parents := make([][]int, n)
	for i := range n {
		parents[i] = make([]int, n*c+m+1)
	}
	for k := 1; k <= n; k++ {
		nf := make([]int, c*n+m+1)
		q := []int{(k-1)*c}
		for i := k*c; i <= c*n+m-c*(n-k); i++ {
			j := i - c
			for len(q) > 0 && f[q[len(q)-1]] - s[q[len(q)-1]+c] + nums[q[len(q)-1]] <= f[j] - s[j+c] + nums[j] {
				q = q[:len(q)-1]
			}
			q = append(q, j)

			if q[0] < i-c-d {
				q = q[1:]
			}
			
			nf[i] = f[q[0]] - s[q[0]+c] + nums[q[0]] + s[i]
			parents[k-1][i] = q[0]
		}
		f = nf
	}

	ans := make([]int, n)
	pos := n*c+m
	for k := n-1; k >= 0; k-- {
		ans[k] = parents[k][pos]
		pos = ans[k]
	}

	Println(f[c*n+m])
	for k := range n {
		Print(ans[k]+1)
		Print(" ")
	}
}