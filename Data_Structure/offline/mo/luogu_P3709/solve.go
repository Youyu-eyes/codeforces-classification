// https://www.luogu.com.cn/problem/P3709
// 回滚莫队
// 维护区间众数的出现次数
// 也有普通莫队做法，见 [CF 1514D](https://codeforces.com/contest/1514/problem/D)

package main

import (
	"math"
	"slices"
	"sort"
	"cmp"
)

func solve() {
	n, m := II(), II()
	nums := make([]int, n)
	for i := range n {
		nums[i] = II()
	}
	nums, _ = discretize(nums)

	B := int(math.Ceil(float64(n) / math.Sqrt(float64(m))))
	
	type query struct {
		id   int
		l, r int
		qIdx int
	}

	res := 0
	cnt := make([]int, n)
	add := func(x int) {
		cnt[x]++
		res = max(res, cnt[x])
	}

	qs := []query{}
	ans := make([]int, m)
	for qi := range m {
		l, r := II() - 1, II()
		if r - l > B {
			qs = append(qs, query{l/B, l - 1, r, qi})
			continue
		}

		res = 0
		clear(cnt)
		for i := l; i < r; i++ {
			add(nums[i])
		}
		ans[qi] = -res
	}

	slices.SortFunc(qs, func(a, b query) int {
		return cmp.Or(
			cmp.Compare(a.id, b.id),
			cmp.Compare(a.r, b.r),
		)
	})

	var l, r int
	for i, b := range qs {
		start := (b.id + 1) * B
		if b.id == 0 || b.id > qs[i - 1].id {
			l = start - 1
			r = start
			res = 0
			clear(cnt)
		}

		for ; r < b.r; r++ {
			add(nums[r])
		}

		tmp := res
		
		for ; l > b.l; l-- {
			add(nums[l])
		}

		ans[b.qIdx] = -res
		
		res = tmp
		l = start - 1
		for j := b.l + 1; j < start; j++ {
			cnt[nums[j]]--
		}
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
