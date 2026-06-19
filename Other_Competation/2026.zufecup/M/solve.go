// https://codeforces.com/gym/680541/attachments
// 贡献法 + Lazy 线段树

// 核心思路：
// 题目的意思是，[l, r] 区间的更新是从左往右依次更新的，翻转一个数就记录一次，比如 [1, 1, 0, 0]，我们定义 a[0] = 0，则为 a = [0, 1, 1, 0, 0]，然后更新 [1, 4]
// 第一次 a = [0, 0, 1, 0, 0]，然后 (1, 2) 贡献一个前沿
// 第二次 a = [0, 0, 0, 0, 0]，没有贡献
// 第三次 a = [0 ,0 ,0 ,1, 0]，此时 (2, 3) 贡献一个前沿
// 第四次 a = [0, 0, 0, 1, 1]，没有贡献
// 我们的思路是，翻转前 a[i] = 1 时，讨论 a[i + 1] 翻转前的状态
// a[i + 1] = 1，则翻转到 a[i] = 0，还没有翻转 a[i + 1] 时，a[i] = 0, a[i + 1] = 1，贡献一次
// a[i + 1] = 0，则翻转到 a[i + 1] = 1，此时 a[i] 已经被翻转为 a[i] = 0 了，a[i] = 0, a[i + 1] = 1，贡献一次
// 因此 a[i] = 1 时，无论 a[i + 1] 初始是什么，a[i] 一定会对答案做出一次贡献

// 我们需要一个数据结构支持维护区间和，区间更新，区间查询，单点查询
// 适用 Lazy 线段树
// 实现细节：
// 定义懒标记为区间是否需要执行翻转操作

// 如果整个区间都需要翻转，则区间和为 区间长度 - 现在的区间和
// 下传懒标记时，如果区间是 [l, r] 的子区间，执行翻转操作

// 连续两次翻转等于没翻转，因此用异或合并懒标记

// 时间复杂度：O(mlogn)
// 空间复杂度：O(n)

package main

import (
	"math/bits"
)

func solve() {
	n, m := II(), II()
	Lsegtree := NewLazySegmentTreeBySize(n + 1, 0)
	ans := 0
	var resl, resr int
	for i := 0; i < m; i++ {
		l, r := II(), II()
		resl, resr = 0, 0
		if Lsegtree.query(l - 1, l - 1) == 0 && Lsegtree.query(l, l) == 0 {
			resl = 1
		}
		if r < n && Lsegtree.query(r, r) == 1 && Lsegtree.query(r + 1, r + 1) == 1 {
			resr = 1
		}
		sumIn := 0
		if l < r {
			sumIn = Lsegtree.query(l, r - 1)
		}
		ans += resl + resr + sumIn
		Lsegtree.update(l, r, 1)
	}
	Println(ans)
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
	return a + b // **根据题目修改**
}

// 合并两个懒标记
func (st *LazySegmentTree) mergeTodo(a, b LazyF) LazyF {
	return a ^ b // **根据题目修改**
}

// 把懒标记作用到 node 子树（本例为区间加）
func (st *LazySegmentTree) apply(node, l, r int, todo LazyF) {
	// 计算 tree[node] 区间的整体变化
	// 如果是 max or min 写成 st.tree[node].val += todo;
	if todo == 1 {
		st.tree[node].val = LazyT(r-l+1) - st.tree[node].val // **根据题目修改**
	}
	
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
