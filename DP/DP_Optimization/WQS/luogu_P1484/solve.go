// https://www.luogu.com.cn/problem/P1484
// 洛谷绿 - 提高
// 另有反悔贪心做法

// WQS 二分做法
// 具体简介见主仓库 笔记·浅析 WQS 二分

// 二分切线斜率，内层就变成了无数量约束的种树
// 这是 leetcode 198.打家劫舍
// 转移时统计选了多少数

// 右端点：所有正数和
// 左端点：所有负数和

package main

type pair struct {
	x, b int
}

func less(f0, f1 pair) bool {
	return f0.b < f1.b || f0.b == f1.b && f0.x > f1.x
}

func solve() {
	n, limit := II(), II()
	nums := make([]int, n)
	left, right := -1, 1
	for i := range n {
		x := II()
		nums[i] = x
		if x < 0 {
			left += x
		} else {
			right += x
		}
	}

	dp := func(k int) pair {
		f0 := pair{0, 0}
		f1 := pair{0, 0}
		for _, x := range nums {
			cand1 := f1
			cand1.b += x - k
			f := f0
			if less(f0, cand1) {
				f = cand1
				f.x++
			}
			f0, f1 = f, f0
		}
		return f0
	}

	res := dp(0)
	if res.x <= limit {
		Println(res.b)
		return
	}

	var ans int
	for left + 1 < right {
		mid := (left + right) >> 1
		res := dp(mid)
		if res.x <= limit {
			ans = mid * limit + res.b
			right = mid
		} else {
			left = mid
		}
	}

	Println(ans)
}