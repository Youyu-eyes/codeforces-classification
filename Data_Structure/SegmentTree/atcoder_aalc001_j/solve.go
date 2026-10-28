// https://atcoder.jp/contests/aalc001/tasks/aalc001_j
// 洛谷绿 - 普及+/提高-

// 灵茶の试炼 2026/10/28

// 离散化所有下标，每个叶子节点 a[i] 保存 size 和 sum
// 初始化时 a[i] = sorted[i + 1] - sorted[i]
// 此时，对于查询区间 [l, r)，实际查询区间为 [l, r-1]
// 懒标记定义为 0/1，表示区间是否翻转
// 同时翻转偶数次相当于不反转，因此采用 异或 运算维护

package main

import (
	"slices"
	"sort"
	"math/bits"
)

func solve() {
	n := II()
	queries := make([][]int, n)
	sorted := make([]int, 2*n)
	for i := range n {
		op, l, r := II(), II() - 1, II()
		queries[i] = []int{op, l, r}
		sorted[i] = l
		sorted[n+i] = r
	}
	slices.Sort(sorted)
	sorted = slices.Compact(sorted)
	m := len(sorted)

	a := make([]info, m - 1)
	for i := range m-1 {
		a[i] = info{0, sorted[i + 1] - sorted[i]}
	}

	lsgt := newLazySegmentTree(a)
	for _, q := range queries {
		l := sort.SearchInts(sorted, q[1])
		r := sort.SearchInts(sorted, q[2])

		if q[0] == 1 {
			lsgt.update(1, l, r-1, 1)
		} else {
			Println(lsgt.query(1, l, r-1).sum)
		}
	}
}

const todoInit = 0
type info struct {
	sum, size int
}

type lazySeg []struct {
	l, r int
	todo int
	info
}

func (lazySeg) mergeInfo(a, b info) info {
	return info{a.sum + b.sum, a.size + b.size}
}

func (lazySeg) mergeTodo(a, b int) int {
	return a ^ b
}

func (t lazySeg) apply(o int, f int) {
	if f == 0 {
		return
	}
	cur := &t[o]
	cur.sum = cur.size - cur.sum
	cur.todo = t.mergeTodo(f, cur.todo)
}

func (t lazySeg) maintain(o int) {
	t[o].info = t.mergeInfo(t[o<<1].info, t[o<<1|1].info)
}

func (t lazySeg) spread(o int) {
	f := t[o].todo
	if f == todoInit {
		return
	}
	t.apply(o<<1, f)
	t.apply(o<<1|1, f)
	t[o].todo = todoInit
}

func (t lazySeg) build(a []info, o, l, r int) {
	t[o].l, t[o].r = l, r
	t[o].todo = todoInit
	if l == r {
		t[o].info = a[l]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

// 调用时 o=1  [l,r] 0<=l<=r<=n-1
func (t lazySeg) update(o, l, r int, f int) {
	if l <= t[o].l && t[o].r <= r {
		t.apply(o, f)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, f)
	}
	if m < r {
		t.update(o<<1|1, l, r, f)
	}
	t.maintain(o)
}

// 调用时 o=1  [l,r] 0<=l<=r<=n-1
func (t lazySeg) query(o, l, r int) info {
	if l <= t[o].l && t[o].r <= r {
		return t[o].info
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	lRes := t.query(o<<1, l, r)
	rRes := t.query(o<<1|1, l, r)
	return t.mergeInfo(lRes, rRes)
}

func newLazySegmentTree(a []info) lazySeg {
	n := len(a)
	if n == 0 {
		panic("slice can't be empty")
	}
	t := make(lazySeg, 2<<bits.Len(uint(n-1)))
	t.build(a, 1, 0, n-1)
	return t
}

func newLazySegmentTreeBySize(n int, defaultVal info) lazySeg {
	a := make([]info, n)
	for i := range a {
		a[i] = defaultVal
	}
	return newLazySegmentTree(a)
}
