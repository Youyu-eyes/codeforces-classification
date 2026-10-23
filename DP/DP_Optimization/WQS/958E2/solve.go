// https://codeforces.com/problemset/problem/958/E2
// 2200
// 买卖股票模型
// 选择一些子数组，每个子数组不相交，求子数组的头尾之差，然后求和
// 买卖股票是求最大，本题是求最小

// 本质上就是 leetcode 188. 买卖股票的最佳时机 IV
// 可以看 零神 的有关 WQS 二分的题解

// 时间复杂度和空间复杂度见 零神 题解

package main

import "slices"

// leetcode 188. 买卖股票的最佳时机 IV
// 求最小
func minProfit(k int, prices []int) int {
    dp := func(fee int) pair {
        f0 := pair{0, 0}
        f1 := pair{inf, 0}
        var best pair
        for _, p := range prices {
            nf0, nf1 := f0, f1
            best = pair{f1.b + p - fee, f1.x + 1}
            if less(f0, best) {
                nf0 = best
            }

            best = pair{f0.b - p, f0.x}
            if less(f1, best) {
                nf1 = best
            }
            f0, f1 = nf0, nf1
        }

        return f0
    }

    ans := 0
    left, right := -sum(prices) - 1, sum(prices) + 1
    for left + 1 < right {
        mid := left + (right - left) >> 1
        res := dp(mid)
        if res.x <= k {
            ans = k * mid + res.b
        	left = mid
        } else {
            right = mid
        }
    }

    return ans
}

type pair struct{ b, x int }

// p 比 q 更劣
func less(p, q pair) bool {
	return p.b > q.b || p.b == q.b && p.x > q.x
}

func solve() {
	limit, n := II(), II()
	nums := make([]int, n)
	for i := range nums { nums[i] = II() }

	slices.Sort(nums)
	Println(minProfit(limit, nums))
}