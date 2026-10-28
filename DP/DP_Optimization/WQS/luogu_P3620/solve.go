// https://www.luogu.com.cn/problem/P3620
// 洛谷蓝 - 提高+/省选-

// 有贪心结论显然成立：
// 选择的连边一定是相邻的两个公司
// 由于我们需要选 2*k 个公司
// 因此相邻的连边不能同时选择

// 将距离求差分，问题变成了
// 在差分数组中，选择恰好 k 个数的最小和，每个数不相邻

// WQS 二分去掉选择约束后，变成了选一些不相邻的数，使得和最小

// 这是打家劫舍问题，不过是 “廉洁” 的小偷
// 你可以用负数，但是更好的方法是更改 less 函数

package main

type pair struct {
	x, b int
}

func less(f0, f1 pair) bool {
	return f0.b > f1.b || f0.b == f1.b && f0.x > f1.x
}

func solve() {
	N, limit := II(), II()
	n := N-1
	dis := II()
	nums := make([]int, n)
	for i := range n {
		x := II()
		nums[i] = x - dis
		dis = x
	}

	dp := func(k int) pair {
		f0 := pair{0, 0}
		f1 := pair{0, 0}
		for _, x := range nums {
			f1.b += x - k
			f1.x++
			f := f0
			if less(f0, f1) {
				f = f1
			}
			f0, f1 = f, f0
		}
		return f0
	}

	left, right := -1, 1
	for _, x := range nums {
		right += x
	}

	var ans int
	for left + 1 < right {
		mid := (left + right) >> 1
		res := dp(mid)
		if res.x <= limit {
			ans = mid*limit + res.b
			left = mid
		} else {
			right = mid
		}
	}

	Println(ans)
}