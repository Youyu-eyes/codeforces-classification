package main

import (
	"bufio"
	"fmt"
	"os"
	"math/bits"
)

const (
	MOD = 1_000_000_007
	inf = 0x3f3f3f3f
)

var (
	in  *bufio.Reader
	out *bufio.Writer
)

// n := II()
func II() (x int) {
	fmt.Fscan(in, &x)
	return
}

// n := IF()
func IF() (x float64) {
	fmt.Fscan(in, &x)
	return
}

// s := IS()
func IS() (x string) {
	fmt.Fscan(in, &x)
	return
}

func Print(a ...any)            { fmt.Fprint(out, a...) }
func Println(a ...any)          { fmt.Fprintln(out, a...) }
func Printf(f string, a ...any) { fmt.Fprintf(out, f, a...) }


func solve() {
	n := II()
	k := II()
	nums := make([]int, n)
	f := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = II()
		f[i] = 1
	}

	for K := 1; K <= k; K++ {
		segtree := NewSegmentTreeBySize(n, 0)
		nf := make([]int, n)
		for i := 0; i < n; i++ {
			nf[i] = segtree.query(0, nums[i] - 1)
			segtree.update(nums[i], f[i])
		}
		copy(f, nf)
	}
	ans := 0
	for i := 0; i < n; i++ { ans += f[i] }
	Println(ans)
}


func main() {
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	T := 1

	for t := 0; t < T; t++ {
		solve()
	}
}


func abs[T int | int64 | float64](x T) T {
	if x >= 0 { return x }
	return -x
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// ------- 线段树 ------- //
type SegT = int

type SegmentTree struct {
	n    int
	tree []SegT
}

func (st *SegmentTree) mergeVal(a, b SegT) SegT {
	return a + b // **根据题目修改** 可维护 max(default = -inf), min(default = inf), gcd, +, &(-1), |, ^, ( * ) % MOD(1) 等
}

func (st *SegmentTree) maintain(node int) {
	st.tree[node] = st.mergeVal(st.tree[node<<1], st.tree[node<<1|1])
}

func (st *SegmentTree) build(a []SegT, node, l, r int) {
	if l == r {
		st.tree[node] = a[l]
		return
	}
	m := (l + r) >> 1
	st.build(a, node<<1, l, m)
	st.build(a, node<<1|1, m+1, r)
	st.maintain(node)
}

func (st *SegmentTree) updateNode(node, l, r, i int, val SegT) {
	if l == r {
		st.tree[node] = val
		return
	}
	m := (l + r) >> 1
	if i <= m {
		st.updateNode(node<<1, l, m, i, val)
	} else {
		st.updateNode(node<<1|1, m+1, r, i, val)
	}
	st.maintain(node)
}

func (st *SegmentTree) queryNode(node, l, r, ql, qr int) SegT {
	if ql <= l && r <= qr {
		return st.tree[node]
	}
	m := (l + r) >> 1
	if qr <= m {
		return st.queryNode(node<<1, l, m, ql, qr)
	}
	if ql > m {
		return st.queryNode(node<<1|1, m+1, r, ql, qr)
	}
	lRes := st.queryNode(node<<1, l, m, ql, qr)
	rRes := st.queryNode(node<<1|1, m+1, r, ql, qr)
	return st.mergeVal(lRes, rRes)
}

func (st *SegmentTree) findFirstNode(node, l, r, ql, qr int, f func(SegT) bool) int {
	if r < ql || l > qr {
		return -1
	}
	if !f(st.tree[node]) {
		return -1
	}
	if l == r {
		return l
	}
	m := (l + r) >> 1
	leftRes := st.findFirstNode(node<<1, l, m, ql, qr, f)
	if leftRes != -1 {
		return leftRes
	}
	return st.findFirstNode(node<<1|1, m+1, r, ql, qr, f)
}

func (st *SegmentTree) findLastNode(node, l, r, ql, qr int, f func(SegT) bool) int {
	if r < ql || l > qr {
		return -1
	}
	if !f(st.tree[node]) {
		return -1
	}
	if l == r {
		return l
	}
	m := (l + r) >> 1
	rightRes := st.findLastNode(node<<1|1, m+1, r, ql, qr, f)
	if rightRes != -1 {
		return rightRes
	}
	return st.findLastNode(node<<1, l, m, ql, qr, f)
}

// NewSegmentTreeBySize 线段树维护一个长为 n 的数组（下标从 0 到 n-1），元素初始值为 init_val
func NewSegmentTreeBySize(n int, initVal SegT) *SegmentTree {
	a := make([]SegT, n)
	for i := range a {
		a[i] = initVal
	}
	return NewSegmentTree(a)
}

// NewSegmentTree 线段树维护数组 a
func NewSegmentTree(a []SegT) *SegmentTree {
	n := len(a)
	// 2 << bits.Len(uint(len(a)-1)) 等同于 C++ 中的 2 << bit_width(a.size() - 1)
	treeSize := 2 << bits.Len(uint(n-1))
	st := &SegmentTree{
		n:    n,
		tree: make([]SegT, treeSize),
	}
	if n > 0 {
		st.build(a, 1, 0, n-1)
	}
	return st
}

// ----- 外部接口 ----- //
//
// Update 单点更新 将 tree[node] 改成 merge_val(tree[node], val)，需要直接覆盖要修改私有函数
func (st *SegmentTree) update(i int, val SegT) {
	st.updateNode(1, 0, st.n-1, i, val)
}

func (st *SegmentTree) query(ql, qr int) SegT {
	return st.queryNode(1, 0, st.n-1, ql, qr)
}

func (st *SegmentTree) get(i int) SegT {
	return st.queryNode(1, 0, st.n-1, i, i)
}

func (st *SegmentTree) findfirst(ql, qr int, f func(SegT) bool) int {
	return st.findFirstNode(1, 0, st.n-1, ql, qr, f)
}

func (st *SegmentTree) findlast(ql, qr int, f func(SegT) bool) int {
	return st.findLastNode(1, 0, st.n-1, ql, qr, f)
}