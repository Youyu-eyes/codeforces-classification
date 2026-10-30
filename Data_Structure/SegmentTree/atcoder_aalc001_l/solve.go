// https://atcoder.jp/contests/aalc001/tasks/aalc001_l
// 洛谷青 - 提高
// 2026/10/30 灵茶の试炼

// lazy 线段树
// 用线段树维护区间前缀最大值和区间和
// 懒标记为 [0, 1, 2] 表示当前区间在 [a, b, c] 数组中的变换
// 如 [1, 2, 0] 表示当前 a 区间的数是 b 数组的

// 线段树二分
// 定义 s 为 [ql, t[o].r] 的和
// 计算 [ql, t[o].m] 的和
// 如果比 0 大，说明 j 在 m 左侧
// 如果比 0 小，说明 j 在 m 右侧

// 剪枝
// 递归计算 [ql, t[o].m] 的和
// 如果返回的 idx >= 0，则直接返回
// 如果 idx < 0，则说明左半区间没有

// 计算 s + pre([t[o].l, t[o].r]) 的和
// 如果此时没有答案，则 j 一定不在这个区间中，剪枝

package main

import (
	"math/bits"
)

func solve() {
	n, m := II(), II()
	usebuild := make([]info, n)
	for k := range 3 {
		for i := range n {
			x := II()
			usebuild[i][k] = pair{x, x}
		}
	}
	lsgt := newLazySegmentTree(usebuild)
	f := [3][3]int{{1, 0, 2}, {0, 2, 1}, {2, 1, 0}}
	for range m {
		op := II()
		if op == 4 {
			ql := II() - 1
			i, _ := lsgt.findfirst(1, ql, 0)
			if i < 0 {
				i = n
			}
			Println(i+1)
		} else {
			l, r := II() - 1, II() - 1
			lsgt.update(1, l, r, f[op-1])
		}
	}
}

var todoInit = [3]int{0, 1, 2}

type pair struct {
	sum, pre int
}

type info = [3]pair

type lazySeg []struct {
	l, r int
	info
	todo [3]int
}

func (lazySeg) mergeInfo(a, b info) (res info) {
	for i := range 3 {
		res[i].sum = a[i].sum + b[i].sum
		res[i].pre = max(a[i].pre, a[i].sum + b[i].pre)
	}
	return
}

func (lazySeg) mergeTodo(a, b [3]int) (res [3]int) {
	for i := range 3 {
		res[i] = a[b[i]]
	}
	return
}

func (t lazySeg) apply(o int, f [3]int) {
	cur := &t[o]
	old := cur.info
	for i := range 3 {
		cur.info[i] = old[f[i]]
	}
	cur.todo = t.mergeTodo(cur.todo, f)
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
func (t lazySeg) update(o, l, r int, f [3]int) {
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

func (t lazySeg) findfirst(o, ql, s int) (int, int) {
	if t[o].l == t[o].r {
		s += t[o].info[0].sum
		if s > 0 {
			return t[o].l, 0
		}
		return -1, s
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if m < ql {
		return t.findfirst(o<<1|1, ql, s)
	}
	idx, s := t.findfirst(o<<1, ql, s)
	if idx >= 0 {
		return idx, 0
	}

	cur := t[o<<1|1].info[0]
	if s + cur.pre <= 0 {
		return -1, s + cur.sum
	}
	return t.findfirst(o<<1|1, ql, s)
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
