// https://atcoder.jp/contests/aalc001/tasks/aalc001_g
// 洛谷青 - 提高
// 2026/10/29 灵茶の试炼

// Lazy 线段树模拟
// 考虑区间每个数增加 x 后，区间总和会增加多少
// 由于 k 最大只有 4，我们直接维护 4 种情况
// 区间维护 4 种情况的和
// 作用懒标记时，注意中途取模
// 本题极易加法溢出

package main

import (
	"math/bits"
)

func solve() {
	n, m := II(), II()
	a := make([]int, n)
	abuild := make([]info, n)
	for i := range n {
		x := II() % MOD
		x2 := x * x % MOD
		x3 := x2 * x % MOD
		x4 := x3 * x % MOD
		abuild[i] = [4]int{x, x2, x3, x4}
		a[i] = x
	}
	lsgt := newLazySegmentTree(abuild)
	for range m {
		op, l, r := II(), II() - 1, II() - 1
		if op == 1 {
			v := II() % MOD
			lsgt.update(1, l, r, v)
		} else {
			k := II()
			sum := lsgt.query(1, l, r)
			Println(sum[k-1])
		}
	}
}

const todoInit = 0
type info = [4]int

type lazySeg []struct {
	l, r int
	s info
	todo int
}

func (lazySeg) mergeInfo(a, b info) (res info) {
	for i := range 4 {
		res[i] = (a[i] + b[i]) % MOD
	}
	return
}

func (lazySeg) mergeTodo(a, b int) int {
	return (a + b) % MOD
}

func (t lazySeg) apply(o int, f int) {
	cur := &t[o]
	x1 := f
	x2 := x1 * f % MOD
	x3 := x2 * f % MOD
	x4 := x3 * f % MOD
	sz := cur.r - cur.l + 1
	cur.s[3] = (cur.s[3] + 4 * (cur.s[2] * x1 + cur.s[0] * x3) % MOD + 6 * cur.s[1] * x2 % MOD + sz * x4 % MOD) % MOD
	cur.s[2] = (cur.s[2] + 3 * (cur.s[1] * x1 + cur.s[0] * x2) % MOD + sz * x3 % MOD) % MOD
	cur.s[1] = (cur.s[1] + 2 * cur.s[0] * x1 % MOD + sz * x2 % MOD) % MOD
	cur.s[0] = (cur.s[0] + sz * x1 % MOD) % MOD
	cur.todo = t.mergeTodo(f, cur.todo)
}

func (t lazySeg) maintain(o int) {
	t[o].s = t.mergeInfo(t[o<<1].s, t[o<<1|1].s)
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
		t[o].s = a[l]
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
		return t[o].s
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
