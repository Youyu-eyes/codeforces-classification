// https://www.spoj.com/problems/DQUERY/
// HH的项链

package main

import (
	"slices"
	"sort"
)

func solve() {
	n := II()
	nums := make([]int, n)
	for i := range n {
		nums[i] = II()
	}
	nums, _ = discretize(nums)

	m := II()
	type query struct {
		l    int
		qIdx int
	}

	qs := make([][]query, n)
	for qi := range m {
		l, r := II() - 1, II() - 1
		qs[r] = append(qs[r], query{l, qi})
	}

	ans := make([]int, m)
	last := make([]int, n)
	bit := newFenwickTree(n)
	for i := range n {
		last[i] = -1
	}
	for r, x := range nums {
		if last[x] >= 0 {
			bit.update(last[x], -1)
		}
		last[x] = r
		bit.update(r, 1)
		for _, q := range qs[r] {
			ans[q.qIdx] = bit.query(q.l, r)
		}
	}
	for i := range ans { Println(ans[i]) }
}

type fenwick []int

func newFenwickTree(n int) fenwick {
	return make(fenwick, n+1) // 使用下标 1 到 n
}

func (f fenwick) update(i, val int) {
	for i++; i < len(f); i += i & -i {
		f[i] += val
	}
}

func (f fenwick) pre(i int) (res int) {
	for i++; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}

func (f fenwick) query(l, r int) int {
	return f.pre(r) - f.pre(l-1)
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
