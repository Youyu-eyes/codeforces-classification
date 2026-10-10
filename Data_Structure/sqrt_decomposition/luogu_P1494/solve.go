// https://www.luogu.com.cn/problem/P1494
// 【国家集训队】小 z 的袜子

// 莫队 + 组合数学
// \dfrac{\sum\binom{cnt[i]}{2}}{\binom{L}{2}}
// 化简式子后，发现我们需要查询的实际上是 \sum\cnt[i]^2

// 用莫队维护这个信息，那么问题就变成了
// 【国家集训队】小 b 的询问

package main

import (
	"math"
	"slices"
	"cmp"
	"strconv"
)

func solve() {
	n, m := II(), II()
	nums := make([]int, n)
	for i := range n {
		nums[i] = II()
	}

	B := int(math.Ceil(float64(n) / math.Sqrt(float64(m))))

	type query struct {
		id   int
		l, r int
		qIdx int
	}

	qs := make([]query, m)
	for qi := range m {
		l, r := II() - 1, II()
		qs[qi] = query{l/B, l - 1, r, qi}
	}

	slices.SortFunc(qs, func(a, b query) int {
		if a.id != b.id {
			return cmp.Compare(a.id, b.id)
		}
		if a.id & 1 > 0 {
			return cmp.Compare(b.r, a.r)
		}
		return cmp.Compare(a.r, b.r)
	})

	res := 0
	cnt := make([]int, n + 1)
	
	add := func(x int) {
		res += 2 * cnt[x] + 1
		cnt[x]++
	}

	del := func(x int) {
		res -= 2 * cnt[x] - 1
		cnt[x]--
	}

	ans := make([]string, m)
	l, r := -1, 0
	for _, b := range qs {
		for ; r < b.r; r++ {
			add(nums[r])
		}
		for ; l > b.l; l-- {
			add(nums[l])
		}
		for r > b.r {
			r--
			del(nums[r])
		}
		for l < b.l {
			l++
			del(nums[l])
		}
		L := b.r - b.l - 1
		A := res - L
		B := L * (L - 1)
		if A == 0 || B == 0 {
			A, B = 0, 1
		} else {
			g := gcd(A, B)
			A /= g
			B /= g
		}
		ans[b.qIdx] = strconv.Itoa(A) + "/" + strconv.Itoa(B)
	}
	
	for i := range ans { Println(ans[i]) }
}
