// https://www.luogu.com.cn/problem/P2709
// 【国家集训队】小 b 的询问
// 莫队模板

package main

import (
	"slices"
	"cmp"
	"math"
)

func solve() {
	n, m, k := II(), II(), II()
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
	cnt := make([]int, k + 1)
	
	add := func(x int) {
		res += 2 * cnt[x] + 1
		cnt[x]++
	}

	del := func(x int) {
		res -= 2 * cnt[x] - 1
		cnt[x]--
	}

	ans := make([]int, m)
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
		ans[b.qIdx] = res
	}
	
	for i := range ans { Println(ans[i]) }
}