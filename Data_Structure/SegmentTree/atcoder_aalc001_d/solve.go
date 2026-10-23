// https://atcoder.jp/contests/aalc001/tasks/aalc001_d
// 1700
// 分治思想/线段树

package main

import (
	"math/bits"
)

func solve() {
	n, m := II(), II()
	sgt := newSegmentTreeBySize(n, info{0, [3]int{0, 0, 0}})
	update := func(i, x int) {
		cur := [3]int{0, 0, 0}
		cur[x] = 1
		sgt.update(1, i, info{cnt: cur})
	}

	for i := range n {
		x := II()
		update(i, x)
	}

	for range m {
		op := II()
		if op == 1 {
			i, x := II() - 1, II()
			update(i, x)
			continue
		}
		l, r := II() - 1, II() - 1
		Println(sgt.query(1, l, r).inv)
	}
}

type info struct {
	inv int
	cnt [3]int
}

type seg []struct {
	l, r int
	val  info
}

// 可维护 max(default = -inf), min(default = inf), gcd, +, &(-1), |, ^, ( * ) % MOD(1) 等
func (seg) mergeInfo(a, b info) (res info) {
	res.inv = a.inv + b.inv + (a.cnt[2] + a.cnt[1]) * b.cnt[0] + a.cnt[2] * b.cnt[1]
	res.cnt[0] = a.cnt[0] + b.cnt[0]
	res.cnt[1] = a.cnt[1] + b.cnt[1]
	res.cnt[2] = a.cnt[2] + b.cnt[2]
	return
}

func (t seg) maintain(o int) {
	t[o].val = t.mergeInfo(t[o<<1].val, t[o<<1|1].val)
}

func (t seg) build(a []info, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].val = a[l]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

// 调用时 o=1  0<=i<=n-1
func (t seg) update(o, i int, val info) {
	if t[o].l == t[o].r {
		t[o].val = val
		return
	}
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		t.update(o<<1, i, val)
	} else {
		t.update(o<<1|1, i, val)
	}
	t.maintain(o)
}

// 调用时 o=1  [l,r] 0<=l<=r<=n-1
func (t seg) query(o, l, r int) info {
	if l <= t[o].l && t[o].r <= r {
		return t[o].val
	}
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

func newSegmentTree(a []info) seg {
	n := len(a)
	if n == 0 {
		panic("slice can't be empty")
	}
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(a, 1, 0, n-1)
	return t
}

func newSegmentTreeBySize(n int, defaultVal info) seg {
	a := make([]info, n)
	for i := range a {
		a[i] = defaultVal
	}
	return newSegmentTree(a)
}
