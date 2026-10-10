// https://www.spoj.com/problems/DQUERY/
// HH的项链
// 莫队做法

// 正常做法见树状数组题单

package main

import (
	"slices"
	"sort"
	"math"
	"cmp"
)

func solve() {
	n := II()
	nums := make([]int, n)
	for i := range n {
		nums[i] = II()
	}
	nums, _ = discretize(nums)

	type query struct {
		id   int
		l, r int
		qIdx int
	}

	m := II()
	B := int(math.Ceil(float64(n) / math.Sqrt(float64(m))))

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
	cnt := make([]int, n)
	
	add := func(x int) {
		if cnt[x] == 0 {
			res++
		}
		cnt[x]++
	}

	del := func(x int) {
		cnt[x]--
		if cnt[x] == 0 {
			res--
		}
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

func discretize(arr []int) ([]int, []int) {
    unique := slices.Clone(arr)
	slices.Sort(unique)
	unique = slices.Compact(unique)

    rank := make([]int, len(arr))
	for i, x := range arr {
		rank[i] = sort.SearchInts(unique, x)
	}

	return rank, unique
}
