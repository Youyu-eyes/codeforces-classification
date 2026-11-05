// https://atcoder.jp/contests/aalc001/tasks/aalc001_m
// 难度未定级 2026/11/5
// 2026/11/5 灵茶の试炼

// 将位数拆分
// 区间维护一个 mp 数组，视为一个函数
// 接收区间左侧的一个值作为输入，对于某一位 i
// mp[i][0] = 1
// mp[i][1] = (x >> i & 1) ^ 1

// 区间合并时
// 左侧输入 0，mpa 输出 mpa[i][0]，mpb 输出 mpb[i][mpa[i][0]]
// 左侧输入 1，mpa 输出 mpa[i][1]，mpb 输出 mpb[i][mpb[i][1]]

// 时间复杂度：O(nUlog(nU))，其中 n 表示 a 数组长度，U 表示位数

package main

import (
	"math/bits"
)

func solve() {
	n, m := II(), II()
	a := make([]int, n)
	useBuild := make([]info, n)
	for i := range n {
		x := II()
		a[i] = x
		mp := info{}
		for bit := range 30 {
			mp[bit][0] = 1
			mp[bit][1] = (x >> bit & 1) ^ 1
		}
		useBuild[i] = mp
	}
	sgt := newSegmentTree(useBuild)
	for range m {
		op := II()
		if op == 1 {
			idx, val := II() - 1, II()
			mp := info{}
			for bit := range 30 {
				mp[bit][0] = 1
				mp[bit][1] = (val >> bit & 1) ^ 1
			}
			sgt.update(1, idx, mp)
			a[idx] = val
		} else {
			l, r := II() - 1, II() - 1
			if l == r {
				Println(a[l])
			} else {
				mp := sgt.query(1, l+1, r)
				ans := 0
				for i := range 30 {
					ans |= (mp[i][a[l] >> i & 1]) << i
				}
				Println(ans)
			}
		}
	}
}

type info = [30][2]int

type seg []struct {
	l, r int
	val  info
}

// 可维护 max(default = -inf), min(default = inf), gcd, +, &(-1), |, ^, ( * ) % MOD(1) 等
func (seg) mergeInfo(a, b info) (res info) {
	for i := range 30 {
		res[i][0] = b[i][a[i][0]]
		res[i][1] = b[i][a[i][1]]
	}
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
