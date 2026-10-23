// https://codeforces.com/contest/1996/problem/F
// 1900
// 二分答案

// leetcode 3971

package main

import (
	"slices"
)

func solve() {
	n, m := II(), II()
	value := make([]int, n)
	decay := make([]int, n)
	for i := range value {
		value[i] = II()
	}
	for i := range decay {
		decay[i] = II()
	}

    check := func(x int) bool {
        res := m
        for i, v := range value {
            if v >= x {
                d := decay[i]
                res -= (v - x) / d + 1
                if res < 0 {
                    return true
                }
            }
        }
        return false
    }

    left := 0; right := slices.Max(value) + 1
    for left + 1 < right {
        mid := left + (right - left) >> 1
        if check(mid) {
            left = mid
        } else {
            right = mid
        }
    }
    low := left
    ans := 0
    for i, v := range value {
        if v > low {
            d := decay[i]
            x := (v - low - 1) / d + 1
            m -= x
            ans += (v * 2 - d * (x - 1)) * x
        }
    }
    ans /= 2
    ans += m * low
    
	Println(ans)
}