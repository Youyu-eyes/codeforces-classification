// https://atcoder.jp/contests/aalc001/tasks/aalc001_e
// 1500
// 分治思想
// 能分治解决的问题，其带修版本可以用线段树解决

package main

import (
	"math/bits"
)

func solve() {
	n, m := II(), II()
	a := make([]int, n)
	for i := range n {
		a[i] = II()
	}
	s := IS()
	sgt := newSegmentTreeBySize(n, info{0, 0, 0, 0, 0, 0})

	update := func(i, x int) {
		switch s[i] {
		case 'R':
			sgt.update(1, i, info{r: x})
		case 'G':
			sgt.update(1, i, info{g: x})
		case 'B':
			sgt.update(1, i, info{b: x})
		}
	}

	for i, x := range a {
		update(i, x)
	}

	for range m {
		op := II()
		if op == 2 {
			l, r := II() - 1, II() - 1
			res := sgt.query(1, l, r)
			Println(res.rgb)
			continue
		}
		idx, val := II() - 1, II()
		update(idx, val)
	}
}

type info struct {
	rgb int
	rg int
	gb int
	r int
	g int
	b int
}

type seg []struct {
	l, r int
	val  info
}

// 可维护 max(default = -inf), min(default = inf), gcd, +, &(-1), |, ^, ( * ) % MOD(1) 等
func (seg) mergeInfo(a, b info) (res info) {
	res.rgb = (a.rgb + b.rgb + a.rg * b.b + a.r * b.gb) % MOD
	res.rg = (a.rg + b. rg + a.r * b.g) % MOD
	res.gb = (a.gb + b.gb + a.g * b.b) % MOD
	res.r = (a.r + b.r) % MOD
	res.g = (a.g + b.g) % MOD
	res.b = (a.b + b.b) % MOD
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
