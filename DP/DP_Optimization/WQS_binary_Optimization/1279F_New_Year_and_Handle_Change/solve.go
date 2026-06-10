// https://codeforces.com/problemset/problem/1279/F
// 2800
// WQS 二分

// leetcode 2209. 用地毯覆盖后的最少白色砖块
// 数据升级版
// 详见
// 我的题解 
// WQS 二分笔记 https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-v/solutions/3981282/xiao-bai-si-lu-bi-ji-xi-lie-qian-xi-wqse-nhw1/

package main

import(
	"strings"
)

// leetcode 2209. 用地毯覆盖后的最少白色砖块
func minimumWhiteTiles(floor string, numCarpets int, carpetLen int) int {
	n := len(floor)
	s := make([]int, n + 1)
	for i, x := range floor {
		s[i + 1] = s[i] + int(x - '0')
	}

	dp := func(k int) pair {
		f := make([]pair, n + 1)
		f[0] = pair{0, 0}
		for i := 1; i < carpetLen && i <= n; i++ {
			opt0 := pair{s[i], 0}
			opt1 := pair{-k, 1}
			if less(opt0, opt1) {
				f[i] = opt1
			} else {
				f[i] = opt0
			}
		}

		mn := pair{inf, 0}
		for i := carpetLen; i <= n; i++ {
			j := i - carpetLen
			
			cand := pair{f[j].b - s[i], f[j].x}
			if less(mn, cand) {
				mn = cand
			}
			
			opt0 := pair{s[i], 0}
			optCarpet := pair{mn.b + s[i] - k, mn.x + 1}
			
			if less(opt0, optCarpet) {
				f[i] = optCarpet
			} else {
				f[i] = opt0
			}
		}
		return f[n]
	}

	res := dp(0)
	if res.x <= numCarpets {
		return res.b
	}

	ans := 0
	left, right := -s[n] - 1, 0
	for left+1 < right {
		mid := left + (right - left) >> 1
		res = dp(mid)
		if res.x <= numCarpets {
			ans = mid*numCarpets + res.b
			left = mid
		} else {
			right = mid
		}
	}

	return ans
}

type pair struct{ b, x int }

// p 比 q 更劣 (返回 true 意味着要被替换掉)
func less(p, q pair) bool {
	return p.b > q.b || p.b == q.b && p.x > q.x
}

func solve() {
	n, k, l := II(), II(), II()
	s := IS()
	var floor1, floor2 strings.Builder
	floor1.Grow(n)
	floor2.Grow(n)
	for _, c := range s {
		if c >> 5 & 1 > 0 {
			floor1.WriteByte('1')
			floor2.WriteByte('0')
		} else {
			floor1.WriteByte('0')
			floor2.WriteByte('1')
		}
	}

	ans1 := minimumWhiteTiles(floor1.String(), k, l)
	ans2 := minimumWhiteTiles(floor2.String(), k, l)
	Println(min(ans1, ans2))
}