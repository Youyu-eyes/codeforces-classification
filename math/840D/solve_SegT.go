// https://codeforces.com/contest/840/problem/D
// 2500
// 线段树 + 数学
// Misra-Gries 算法，扩展摩尔投票

// 比较复杂的算法，前置知识 leetcode 169.多数元素 摩尔投票
// 169 我的题解 https://leetcode.cn/problems/majority-element/solutions/3978521/xiao-bai-si-lu-mo-er-tou-piao-fu-misra-g-18fe/

// 由题解中的结论，我们可以用 4 个元素维护一个数组中的 k 多数元素
// 现在我们考虑如何处理区间查询

// 由于本题只需要维护 k 多数元素，不需要维护所有数字的出现次数
// 如果本题需要返回满足频率 f > (r - l + 1) / k 的元素中 f 最大的元素，那就可以直接用回滚莫队
// 但是本题只要满足条件的元素就处于候选集中，然后选择候选集中最小的元素
// 这很难直接用最大出现次数判断

// 我们观察到 (r - l + 1) / k 是一个很特殊的数，表示 n/k，即元素数量需要满足 > n/k
// 这正是 Misra-Gries 算法，扩展摩尔投票 解决的问题
// 由于 Misra-Gries 算法 对顺序无严格限制
// 我们每次选择抵消的组都是集合，而每个组之间也是无序的，因此整个算法即使改变 nums 的元素顺序，也是成立的
// 因此我们可以执行区间合并操作，等价于先在后半段计算候选值，再在前半段计算候选值

// 浅证明一下，这样不会遗漏
// 反证法，如果遗漏了 t，即在合并后的数组中 t 的出现次数 f > n / k
// 假设前半段 a 长度为 n1，后半段 b 长度为 n2
// t 不在 a 和 b 的候选集里，则 f1 <= n1/k; f2 <= n2/k
// f = f1 + f2 <= n1/k + n2/k = n/k
// 矛盾，因此不会遗漏

// 可以区间合并，并且候选集 k - 1 最大为 4，合并时只需要维护 4 长度的候选集即可
// 合并时间复杂度 O(k^2)，但是这里的 k 很小，因此可以暴力合并

// 这样，现在需要区间查询操作，且区间合并的代价很小
// 可以采用线段树

// 线段树维护一个结构体，结构体中存储
// val  表示候选集的值
// hp   表示候选集的值的剩余血量
// size 表示候选集大小

// 当 hp 归零，说明此时为无效数据

// 当我们得到候选集后，需要判断元素在区间中的出现次数，在 169 题解的最后也提到了方法，具体为方法三

// 时间复杂度 O((q + n) × logn × k^2)
// 空间复杂度 O(n)

// 还有一个思路，见 solve_Mo
// 莫队结合懒删除堆的写法，但是复杂度太大没过


package main

import (
	"math/bits"
	"slices"
)

var (
	pos [][]int
)

func count(left, right, x int) int {
	l, _ := slices.BinarySearch(pos[x], left)
	r, _ := slices.BinarySearch(pos[x], right + 1)
	return r - l
}

func solve_SegT() {
	n, q := II(), II()
	nums := make([]SegT, n)
	pos = make([][]int, n + 1)
	for i := range nums {
		x := II()
		pos[x] = append(pos[x], i)
		nums[i] = SegT{[4]int{x}, [4]int{1}, 1}
	}

	segtree := NewSegmentTree(nums)

	for Q := 0; Q < q; Q++ {
		l, r, k := II() - 1, II() - 1, II()
		cur := segtree.query(l, r)
		ans := inf
		for i := 0; i < cur.size; i++ {
			if count(l, r, cur.val[i]) > (r - l + 1) / k {
				ans = min(ans, cur.val[i])
			}
		}
		if ans == inf {
			ans = -1
		}
		Println(ans)
	}
}

// ------- 线段树 ------- //
type SegT struct {
	val [4]int
	hp [4]int
	size int
}

type SegmentTree struct {
	n    int
	tree []SegT
}

func (st *SegmentTree) mergeVal(a, b SegT) SegT {
	for i := 0; i < b.size; i++ {
		found := false
		for j := 0; j < a.size; j++ {
			if b.val[i] == a.val[j] {
				a.hp[j] += b.hp[i]
				found = true
			}
		}
		if found {
			continue
		}
		if a.size < 4 {
			a.val[a.size] = b.val[i]
			a.hp[a.size] = b.hp[i]
			a.size++
		} else {
			minHP := b.hp[i]
			for _, hp := range a.hp {
				minHP = min(minHP, hp)
			}
			for j := 0; j < 4; j++ {
				a.hp[j] -= minHP
			}
			b.hp[i] -= minHP
			if b.hp[i] > 0 {
				for j, hp := range a.hp {
					if hp == 0 {
						a.val[j] = b.val[i]
						a.hp[j] = b.hp[i]
						break
					}
				}
			}

			for j, hp := range a.hp {
				if hp == 0 && j < a.size - 1 {
					a.val[j] = a.val[a.size - 1]
					a.hp[j] = a.hp[a.size - 1]
					a.size--
				}
			}
		}
	}
	return a
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
		st.tree[node] = st.mergeVal(st.tree[node], val) // 如果想直接覆盖就改成 st.tree[node] = val;
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

// [ql, qr] 双闭
func (st *SegmentTree) query(ql, qr int) SegT {
	return st.queryNode(1, 0, st.n-1, ql, qr)
}

func (st *SegmentTree) get(i int) SegT {
	return st.queryNode(1, 0, st.n-1, i, i)
}
