// https://www.luogu.com.cn/problem/P4513
// 最大子数组和
// 线段树

// 如果一个题的可以用分治解决，则这个题的带修版本可以用线段树解决

// 思考将数组分为两部分，分别维护两边的最大子数组和
// 则合并后，答案为两边的最大子数组和的最大值
// 同时需要考虑合并后新的子数组和
// 定义 f[0] 为最大子数组和，两端可选可不选
// 定义 f[1] 右端点必选，左端点可选可不选
// 定义 f[2] 左端点必选，右端点可选可不选
// 定义 f[3] 左右端点必选（子数组和）

// 状态转移方程见 mergeInfo 函数

package main

import (
	"math/bits"
)

func solve() {
	n, m := II(), II()
	nums := make([]int, n)
	for i := range n {
		nums[i] = II()
	}
	sgt := newSegmentTreeBySize(n, info{0, 0, 0, 0})
	for i, x := range nums {
		sgt.update(1, i, info{x, x, x, x})
	}
	for range m {
		op := II()
		if op == 1 {
			l, r := II() - 1, II() - 1
			if l > r {
				l, r = r, l
			}
			Println(sgt.query(1, l, r)[0])
			continue
		}
		idx, val := II() - 1, II()
		sgt.update(1, idx, info{val, val, val, val})
	}
}

type info = [4]int

type seg []struct {
	l, r int
	val  info
}

func (seg) mergeInfo(a, b info) (res info) {
	res[0] = max(a[0], b[0], a[1] + b[2])
	res[1] = max(b[1], a[1] + b[3])
	res[2] = max(a[2], a[3] + b[2])
	res[3] = a[3] + b[3]
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
