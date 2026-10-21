// https://www.luogu.com.cn/problem/P3834
// 整体二分做法
// 另有主席树做法，见线段树题单

// 2026/9/22 的茶

package main

import (
	"slices"
)

func solve() {
	n, m := II(), II()
	a := make([]int, n)
	idx := make([]int, n)
	for i := range n {
		a[i] = II()
		idx[i] = i
	}
	type query struct {
		l, r, k int
	}
	qs := make([]query, m)
	qIdx := make([]int, m)
	for i := range m {
		qs[i] = query{II() - 1, II() - 1, II()}
		qIdx[i] = i
	}
	bit := newFenwickTree(n)

	unique := slices.Clone(a)
	slices.Sort(unique)
	unique = slices.Compact(unique)
	
	var obs func([]int, []int, int, int)
	obs = func(idx, qIdx []int, low, high int) {
		if len(qIdx) == 0 {
			return
		}
		if low + 1 == high {
			for _, qi := range qIdx {
				qs[qi].k = unique[high]
			}
			return
		}
		
		mid := (low + high) >> 1
		x := unique[mid]

		var b, c []int
		for _, i := range idx {
			if a[i] <= x {
				b = append(b, i)
				bit.update(i, 1)
			} else {
				c = append(c, i)
			}
		}

		var d, e []int
		for _, qi := range qIdx {
			q := &qs[qi]
			cnt := bit.query(q.l, q.r)
			if cnt >= q.k {
				d = append(d, qi)
			} else {
				q.k -= cnt
				e = append(e, qi)
			}
		}

		for _, i := range idx {
			if a[i] <= x {
				bit.update(i, -1)
			}
		}

		obs(b, d, low, mid)
		obs(c, e, mid, high)
	}
	obs(idx, qIdx, -1, n)
	for qi := range m {
		Println(qs[qi].k)
	}
}

type fenwick []int

func newFenwickTree(n int) fenwick {
	return make(fenwick, n+1)
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
