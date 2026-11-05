// https://codeforces.com/problemset/problem/2050/F
// 1700

// a[i] % MOD = a[i+1] % MOD
// abs(a[i] - a[i+1]) % MOD = 0

// 定义一个 a 的差分数组，问题变成了
// 给你一些数，求区间内数的 gcd

// 做法很多
// 离线-回滚莫队
// st 表（见 st 表题单）
// 线段树
// 分治

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

	sgt := newSegmentTree(a)
	for range m {
		l, r := II()-1, II()-1
		if l == r {
			Print(0, " ")
		} else {
			Print(sgt.query(1, l, r-1), " ")
		}
	}
	Println()
}

type info = int

type seg []struct {
	l, r int
	val  info
}

// 可维护 max(default = -inf), min(default = inf), gcd, +, &(-1), |, ^, ( * ) % MOD(1) 等
func (seg) mergeInfo(a, b info) info {
	return gcd(a, b)
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
