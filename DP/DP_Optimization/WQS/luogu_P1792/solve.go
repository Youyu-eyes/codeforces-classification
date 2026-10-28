// https://www.luogu.com.cn/problem/P1792
// 洛谷青 - 提高

// WQS 二分做法
// 同 P1484 题，只不过内层无约束 DP 是环形数组
// 这是 leetcode 213.打家劫舍3

// 考虑 nums[0] 选或不选，在 0 处断环

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
		if x > 0 {
			right += x
		} else {
			left += x
		}
	}

	if n >> 1 < limit {
		Println("Error!")
		return
	}

	rob1 := func(start, end, k int) pair {
		f0 := pair{0, 0}
		f1 := pair{0, 0}
		for i := start; i < end; i++ {
			f1.b += nums[i] - k
			f1.x++
			f := f0
			if less(f0, f1) {
				f = f1
			}
			f0, f1 = f, f0
		}
		return f0
	}

	rob := func(k int) pair {
		n := len(nums)
		f0 := rob1(2, n-1, k)
		f0.b += nums[0] - k
		f0.x++
		f1 := rob1(1, n, k)
		if less(f0, f1) {
			return f1
		}
		return f0
	}

	var ans int
	for left + 1 < right {
		mid := (left + right) >> 1
		res := rob(mid)
		if res.x <= limit {
			ans = mid * limit + res.b
			right = mid
		} else {
			left = mid
		}
	}

	Println(ans)
}