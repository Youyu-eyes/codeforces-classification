// https://codeforces.com/problemset/problem/1527/E
// 2500
// 贡献法 + Lazy 线段树优化 DP

// 前置题目：CF 833B

// 定义 f[k][i] 为前 i 个数选 k 个数组的最小代价
// 状态转移方程为
// f[k][i] = min{j = [k - 1, i - 1]}(f[k - 1][j] + w(j, i))

// 观察 w(j, i)，看到 nums[i] 对 w(j, i) 的贡献
// w(j, i + 1) = w(j, i) + (i - last[i])
// last[i] 表示 i 左侧与 nums[i] 相等的最近的数的下标，可以预处理

// 而 i - last[i] 与 j 无关，因此 i 的贡献可以直接写成对 (0, last[i]) 的贡献

// 问题变成了，区间查询最小值，区间更新，适用 Lazy 线段树

// 滚动数组优化空间

// 时间复杂度：O(k × n × logn)
// 空间复杂度：O(n)

package main

import (
	"math/bits"
)

func solve() {
	n, k := II(), II()
	nums := make([]int, n)
	last := make([]int, n)
	x2idx := make([]int, n + 1)
	for i := range x2idx { x2idx[i] = -1 }
	for i := range nums {
		nums[i] = II()
		if x2idx[nums[i]] < 0 {
			last[i] = -1
		} else {
			last[i] = x2idx[nums[i]]
		}
		x2idx[nums[i]] = i
	}

	f := make([]int, n + 1)
	for i := range f { f[i] = inf }
	f[0] = 0

	for K := 1; K <= k; K++ {
		Lsegtree := NewLazySegmentTree(f)
		if last[K - 1] >= 0 {
			Lsegtree.update(0, last[K - 1], K - 1 - last[K - 1])
		}
		for i := K; i <= n - (k - K); i++ {
			f[i] = Lsegtree.query(K - 1, i - 1)
			if i < n && last[i] >= 0 {
				Lsegtree.update(0, last[i], i - last[i])
			}
		}
	}

	Println(f[n])
}

// ------- Lazy线段树 ------- //
type LazyT = int
type LazyF = int

const TODO_INIT LazyF = 0 // **根据题目修改** 懒标记初始值

type LazyNode struct {
	val  LazyT
	todo LazyF
}

type LazySegmentTree struct {
	n    int
	tree []LazyNode
}

// 合并两个 val
func (st *LazySegmentTree) mergeVal(a, b LazyT) LazyT {
	return min(a, b) // **根据题目修改**
}

// 合并两个懒标记
func (st *LazySegmentTree) mergeTodo(a, b LazyF) LazyF {
	return a + b // **根据题目修改**
}

// 把懒标记作用到 node 子树（本例为区间加）
func (st *LazySegmentTree) apply(node, l, r int, todo LazyF) {
	// 计算 tree[node] 区间的整体变化
	// 如果是 max or min 写成 st.tree[node].val += todo;
	st.tree[node].val += todo // **根据题目修改**
	st.tree[node].todo = st.mergeTodo(todo, st.tree[node].todo)
}

// 把当前节点的懒标记下传给左右儿子
func (st *LazySegmentTree) spread(node, l, r int) {
	todo := st.tree[node].todo
	if todo == TODO_INIT { // 没有需要下传的信息
		return
	}
	m := (l + r) >> 1
	st.apply(node<<1, l, m, todo)
	st.apply(node<<1|1, m+1, r, todo)
	st.tree[node].todo = TODO_INIT // 下传完毕
}

// 合并左右儿子的 val 到当前节点的 val
func (st *LazySegmentTree) maintain(node int) {
	st.tree[node].val = st.mergeVal(st.tree[node<<1].val, st.tree[node<<1|1].val)
}

// 用 a 初始化线段树
// 时间复杂度 O(n)
func (st *LazySegmentTree) build(a []LazyT, node, l, r int) {
	st.tree[node].todo = TODO_INIT
	if l == r { // 叶子
		st.tree[node].val = a[l] // 初始化叶节点的值
		return
	}
	m := (l + r) >> 1
	st.build(a, node<<1, l, m)     // 初始化左子树
	st.build(a, node<<1|1, m+1, r) // 初始化右子树
	st.maintain(node)
}

func (st *LazySegmentTree) updateNode(node, l, r, ql, qr int, f LazyF) {
	if ql <= l && r <= qr { // 当前子树完全在 [ql, qr] 内
		st.apply(node, l, r, f)
		return
	}
	st.spread(node, l, r)
	m := (l + r) >> 1
	if ql <= m { // 更新左子树
		st.updateNode(node<<1, l, m, ql, qr, f)
	}
	if qr > m { // 更新右子树
		st.updateNode(node<<1|1, m+1, r, ql, qr, f)
	}
	st.maintain(node)
}

func (st *LazySegmentTree) queryNode(node, l, r, ql, qr int) LazyT {
	if ql <= l && r <= qr { // 当前子树完全在 [ql, qr] 内
		return st.tree[node].val
	}
	st.spread(node, l, r)
	m := (l + r) >> 1
	if qr <= m { // [ql, qr] 在左子树
		return st.queryNode(node<<1, l, m, ql, qr)
	}
	if ql > m { // [ql, qr] 在右子树
		return st.queryNode(node<<1|1, m+1, r, ql, qr)
	}
	lRes := st.queryNode(node<<1, l, m, ql, qr)
	rRes := st.queryNode(node<<1|1, m+1, r, ql, qr)
	return st.mergeVal(lRes, rRes)
}

// 私有递归：查找第一个满足条件的下标
func (st *LazySegmentTree) findFirstNode(node, l, r, ql, qr int, f func(LazyT) bool) int {
	if r < ql || l > qr {
		return -1
	}
	if !f(st.tree[node].val) {
		return -1
	}
	if l == r {
		return l
	}
	st.spread(node, l, r) // 下传懒标记
	m := (l + r) >> 1
	leftRes := st.findFirstNode(node<<1, l, m, ql, qr, f)
	if leftRes != -1 {
		return leftRes
	}
	return st.findFirstNode(node<<1|1, m+1, r, ql, qr, f)
}

// 私有递归：查找最后一个满足条件的下标
func (st *LazySegmentTree) findLastNode(node, l, r, ql, qr int, f func(LazyT) bool) int {
	if r < ql || l > qr {
		return -1
	}
	if !f(st.tree[node].val) {
		return -1
	}
	if l == r {
		return l
	}
	st.spread(node, l, r)
	m := (l + r) >> 1
	rightRes := st.findLastNode(node<<1|1, m+1, r, ql, qr, f)
	if rightRes != -1 {
		return rightRes
	}
	return st.findLastNode(node<<1, l, m, ql, qr, f)
}

func NewLazySegmentTreeBySize(n int, initVal LazyT) *LazySegmentTree {
	a := make([]LazyT, n)
	for i := range a {
		a[i] = initVal
	}
	return NewLazySegmentTree(a)
}

func NewLazySegmentTree(a []LazyT) *LazySegmentTree {
	n := len(a)
	treeSize := 2 << bits.Len(uint(n-1))
	st := &LazySegmentTree{
		n:    n,
		tree: make([]LazyNode, treeSize),
	}
	if n > 0 {
		st.build(a, 1, 0, n-1)
	}
	return st
}

// ----- 外部接口 ----- //
func (st *LazySegmentTree) update(ql, qr int, f LazyF) {
	st.updateNode(1, 0, st.n-1, ql, qr, f)
}

func (st *LazySegmentTree) query(ql, qr int) LazyT {
	return st.queryNode(1, 0, st.n-1, ql, qr)
}

func (st *LazySegmentTree) findfirst(ql, qr int, f func(LazyT) bool) int {
	return st.findFirstNode(1, 0, st.n-1, ql, qr, f)
}

func (st *LazySegmentTree) findlast(ql, qr int, f func(LazyT) bool) int {
	return st.findLastNode(1, 0, st.n-1, ql, qr, f)
}