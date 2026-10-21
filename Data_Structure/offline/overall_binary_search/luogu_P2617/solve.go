// https://www.luogu.com.cn/problem/P2617
// 整体二分
// 带修区间第 k 小

// 2026/9/23 的茶

package main

import (
	"slices"
)

func solve() {
	n, m := II(), II()
	a := make([]int, n)
	type query struct {
		l, r, k int
	}
	qs := make([]query, n, n + m * 2)
	qIdx := make([]int, n, n + m * 2)
	unique := make([]int, n, n + m * 2)

	for i := range n {
		a[i] = II()
		qs[i] = query{i, a[i], -1}
		qIdx[i] = i
		unique[i] = a[i]
	}

	for range m {
		op := IB()
		if op == 'Q' {
			l, r, k := II() - 1, II() - 1, II()
			qs = append(qs, query{l, r, k})
			qIdx = append(qIdx, len(qIdx))
		} else {
			idx, val := II() - 1, II()
			qIdx = append(qIdx, len(qIdx), len(qIdx) + 1)
			qs = append(qs, query{idx, a[idx], -3})
			qs = append(qs, query{idx, val, -1})
			a[idx] = val
			unique = append(unique, val)
		}
	}

	slices.Sort(unique)
	unique = slices.Compact(unique)
	bit := newFenwickTree(n)

	var obs func([]int, int, int)
	obs = func(qIdx []int, low, high int) {
		for _, qi := range qIdx {
			if qs[qi].k > 0 {
				goto next
			}
		}
		return

	next:
		if low + 1 == high {
			for _, qi := range qIdx {
				if qs[qi].k > 0 {
					qs[qi].k = unique[high]
				}
			}
			return
		}

		mid := (low + high) >> 1
		x := unique[mid]

		var d, e []int
		for _, qi := range qIdx {
			q := &qs[qi]
			if q.k < 0 {
				if q.r <= x {
					bit.update(q.l, q.k + 2)
					d = append(d, qi)
				} else {
					e = append(e, qi)
				}
			} else {
				cnt := bit.query(q.l, q.r)
				if cnt >= q.k {
					d = append(d, qi)
				} else {
					q.k -= cnt
					e = append(e, qi)
				}
			}
		}

		for _, qi := range qIdx {
			q := &qs[qi]
			if q.k < 0 && q.r <= x {
				bit.update(q.l, -q.k-2)
			}
		}

		obs(d, low, mid)
		obs(e, mid, high)
	}
	obs(qIdx, -1, len(unique)-1)
	for _, qi := range qIdx {
		if qs[qi].k >= 0 {
			Println(qs[qi].k)
		}
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
