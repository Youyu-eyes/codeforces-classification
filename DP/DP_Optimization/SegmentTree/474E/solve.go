// https://codeforces.com/contest/474/problem/E
// 2000
// 动态开点线段树优化 DP + 输出具体方案

// 在 DP 转移的同时记录前驱节点，类似 Dijkstra 的输出具体方案
// 因为 DP 本质上是在一张有向无环图 (DAG) 上更新

// 题目本身很简单，最长子序列问题
// 定义 f[i] 为以 i 结尾的最长子序列长度
// 转移方程
// f[i] = max{j | j < i, h[j] <= h[i] - d or h[j] >= h[i] + d}(f[j])
// 典型二维偏序问题，第一维度有序无需排序，第二维度用线段树维护区间最值
// 由于 h 的值域为 10^15，我们需要动态开点线段树

// 时间复杂度: O(n·logMX)
// 空间复杂度: O(n·logMX)

package main

import(
	"fmt"
)

type ValType struct {
	maxLen int
	idx    int
}

func (v ValType) better(other ValType) ValType {
	if v.maxLen >= other.maxLen {
		return v
	}
	return other
}

// ---------- 动态开点线段树 ----------
type DynSegNode struct {
	left, right *DynSegNode
	l, r        int
	val         ValType
}

type DynamicSegmentTree struct {
	root       *DynSegNode
	empty      *DynSegNode
	minX, maxX int
	defaultVal ValType
}

func NewDynamicSegmentTree(minX, maxX int) *DynamicSegmentTree {
	def := ValType{0, -1}
	empty := &DynSegNode{l: 0, r: 0, val: def}
	empty.left = empty
	empty.right = empty

	root := &DynSegNode{l: minX, r: maxX, val: def, left: empty, right: empty}

	return &DynamicSegmentTree{
		root:       root,
		empty:      empty,
		minX:       minX,
		maxX:       maxX,
		defaultVal: def,
	}
}

func (st *DynamicSegmentTree) maintain(o *DynSegNode) {
	o.val = o.left.val.better(o.right.val)
}

func (st *DynamicSegmentTree) updateNode(o *DynSegNode, i int, val ValType) {
	if o.l == o.r {
		// 题目要求取最长，所以保留较优值（长度相同时保留谁都可以）
		o.val = o.val.better(val)
		return
	}
	m := (o.l + o.r) >> 1
	if i <= m {
		if o.left == st.empty {
			o.left = &DynSegNode{l: o.l, r: m, val: st.defaultVal, left: st.empty, right: st.empty}
		}
		st.updateNode(o.left, i, val)
	} else {
		if o.right == st.empty {
			o.right = &DynSegNode{l: m + 1, r: o.r, val: st.defaultVal, left: st.empty, right: st.empty}
		}
		st.updateNode(o.right, i, val)
	}
	st.maintain(o)
}

func (st *DynamicSegmentTree) queryNode(o *DynSegNode, ql, qr int) ValType {
	if o == st.empty || ql > o.r || qr < o.l {
		return st.defaultVal
	}
	if ql <= o.l && o.r <= qr {
		return o.val
	}
	return st.queryNode(o.left, ql, qr).better(st.queryNode(o.right, ql, qr))
}

// 外部接口
func (st *DynamicSegmentTree) update(i int, val ValType) {
	if i < st.minX || i > st.maxX {
		return
	}
	st.updateNode(st.root, i, val)
}

func (st *DynamicSegmentTree) query(ql, qr int) ValType {
	if ql > qr || ql > st.maxX || qr < st.minX {
		return st.defaultVal
	}
	return st.queryNode(st.root, ql, qr)
}

func solve() {
	n := II()
	d := II()
	h := make([]int, n)
	for i := 0; i < n; i++ {
		h[i] = II()
	}

	mxH := 0
	for _, v := range h {
		if v > mxH {
			mxH = v
		}
	}

	f := make([]int, n)
	pre := make([]int, n)
	for i := range pre {
		pre[i] = -1
	}

	seg := NewDynamicSegmentTree(0, mxH)

	f[0] = 1
	pre[0] = -1
	seg.update(h[0], ValType{1, 0})

	for i := 1; i < n; i++ {
		left := ValType{0, -1}
		if h[i]-d >= 0 {
			left = seg.query(0, h[i]-d)
		}
		right := ValType{0, -1}
		if h[i]+d <= mxH {
			right = seg.query(h[i]+d, mxH)
		}
		best := left.better(right)
		f[i] = best.maxLen + 1
		if best.idx != -1 {
			pre[i] = best.idx
		} else {
			pre[i] = -1
		}
		seg.update(h[i], ValType{f[i], i})
	}

	bestLen := 0
	bestIdx := 0
	for i := 0; i < n; i++ {
		if f[i] > bestLen {
			bestLen = f[i]
			bestIdx = i
		}
	}

	Println(bestLen)

	// 回溯路径
	path := make([]int, 0, bestLen)
	for cur := bestIdx; cur != -1; cur = pre[cur] {
		path = append(path, cur)
	}
	// 反转得到从前往后的顺序
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	// 输出（1‑based 下标）
	for i, idx := range path {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, idx + 1)
	}
	fmt.Fprintln(out)
}