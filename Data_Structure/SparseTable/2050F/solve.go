// https://codeforces.com/problemset/problem/2050/F
// 1700

// 此为 st 表做法
// 解析见线段树题单

package main

import (
	"math/bits"
)

func solve() {
	n, m := II(), II()
	cur := II()

	if n == 1 {
		for range m {
			II()
			II()
			Print(0, " ")
		}
		Println()
		return
	}

	a := make([]int, n-1)
	for i := range a {
		x := II()
		a[i] = abs(x - cur)
		cur = x
	}

	buildSTGcd(a)
	for range m {
		l, r := II()-1, II()-1
		if l == r {
			Print(0, " ")
		} else {
			Print(queryGcd(l, r), " ")
		}
	}
	Println()
}

// ------- GCD ST 表 ------- //
// 初始化：buildSTGcd(nums)
// 查询：  queryGcd(l, r)

const maxN = 200_005
var maxLog = bits.Len(uint(maxN))

var stGcd = make([]int, maxLog*maxN)

func buildSTGcd(nums []int) {
	n := len(nums)
	w := bits.Len(uint(n))

	for i := range n {
		stGcd[i] = nums[i] 
	}

	for k := 1; k < w; k++ {
		step := 1 << (k - 1)
		currOffset := k * maxN
		prevOffset := (k - 1) * maxN
		
		for i := 0; i + step + (1 << (k - 1)) <= n; i++ {
			v1 := stGcd[prevOffset + i]
			v2 := stGcd[prevOffset + i + step]
			stGcd[currOffset + i] = gcd(v1, v2)
		}
	}
}

func queryGcd(l, r int) int {
	k := bits.Len(uint(r-l)) - 1
	offset := k * maxN
	
	v1 := stGcd[offset + l]
	v2 := stGcd[offset + r - (1 << k)]
	
	return gcd(v1, v2)
}
