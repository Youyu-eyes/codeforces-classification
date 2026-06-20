package main

import (
	"math/big"
)

// ======= 可持久化李超线段树 模板开始 ======= //

// ------- 可持久化李超线段树 (浮点数版) ------- //
// 初始化：tree := NewPersistentLiChaoTree(minX, maxX)

type Line struct {
	k, b float64
	id   int
}

func newLineDefault() Line {
	return Line{k: 0, b: inf, id: 0}
	// 最大值：return Line{k: 0, b: -inf, id: 0}
}

func (l Line) calc(x int) float64 {
	return l.k*float64(x) + l.b
}

type PersistLiChaoNode struct {
	lc, rc *PersistLiChaoNode
	line   Line
}

type PersistentLiChaoTree struct {
	minX, maxX int
	history    []*PersistLiChaoNode // 维护历史版本根节点
}

func NewPersistentLiChaoTree(minX, maxX int) *PersistentLiChaoTree {
	return &PersistentLiChaoTree{
		minX:    minX,
		maxX:    maxX,
		history: []*PersistLiChaoNode{nil}, // 初始第 0 个版本为空树 (nil)
	}
}

func (st *PersistentLiChaoTree) better(a, b Line, x int) bool {
	if a.id == 0 {
		return false
	}
	if b.id == 0 {
		return true
	}
	va, vb := a.calc(x), b.calc(x)
	if va-vb < -eps {
		return true
	}
	// 最大值：if va - vb > eps { return true }
	if vb-va < -eps {
		return false
	}
	// 最大值：if vb - va > eps { return false }
	return a.id < b.id
}

func (st *PersistentLiChaoTree) insertNode(p *PersistLiChaoNode, l, r int, newLine Line) *PersistLiChaoNode {
	// 可持久化核心：不修改原节点，而是创建一个新节点（拷贝旧节点的状态）
	curr := &PersistLiChaoNode{}
	if p != nil {
		*curr = *p // 浅拷贝左右儿子指针和当前线段
	} else {
		curr.line = newLineDefault()
	}

	mid := (l + r) >> 1
	leftBetter := st.better(newLine, curr.line, l)
	midBetter := st.better(newLine, curr.line, mid)

	if midBetter {
		curr.line, newLine = newLine, curr.line // swap
		midBetter = st.better(newLine, curr.line, mid)
		leftBetter = st.better(newLine, curr.line, l)
	}

	if l == r {
		return curr
	}

	// 下放处于劣势的线段到对应的子树，并更新新节点的子节点指针
	if leftBetter != midBetter {
		curr.lc = st.insertNode(curr.lc, l, mid, newLine)
	} else {
		curr.rc = st.insertNode(curr.rc, mid+1, r, newLine)
	}

	return curr
}

func (st *PersistentLiChaoTree) queryNode(p *PersistLiChaoNode, l, r, x int) float64 {
	if p == nil {
		return inf
		// 最大值：return -inf
	}
	res := float64(inf)
	// 最大值：res := float64(-inf)
	if p.line.id != 0 {
		res = p.line.calc(x)
	}
	if l == r {
		return res
	}
	mid := (l + r) >> 1
	if x <= mid {
		return min(res, st.queryNode(p.lc, l, mid, x))
	}
	return min(res, st.queryNode(p.rc, mid+1, r, x))
	// 最大值：return max(res, ...)
}

// ----- 外部接口 ----- //

// insert 插入一条全局直线，生成一个新版本
// 时空复杂度：O(log C)
func (st *PersistentLiChaoTree) insert(line Line) {
	root := st.history[len(st.history)-1]
	newRoot := st.insertNode(root, st.minX, st.maxX, line)
	st.history = append(st.history, newRoot)
}

// query 查询点 x 的最值（在最新版本上查询）
// 时间复杂度：O(log C)
func (st *PersistentLiChaoTree) query(x int) float64 {
	root := st.history[len(st.history)-1]
	return st.queryNode(root, st.minX, st.maxX, x)
}

// rollBack 返回上一个版本（撤销最后一次插入的直线）
// 时间复杂度：O(1)
func (st *PersistentLiChaoTree) rollBack() {
	if len(st.history) > 1 {
		st.history = st.history[:len(st.history)-1]
	}
}

func (st *PersistentLiChaoTree) clear() {
	st.history = []*PersistLiChaoNode{nil}
}


// ======= 可持久化李超线段树 模板结束 ======= //

// ======= 凸包 模板开始 ======= //

// --- 向量模板 --- //
// 初始化：v := vec{x, y}

type vec struct{ x, y int }

func (a vec) sub(b vec) vec { return vec{a.x - b.x, a.y - b.y} }
func (a vec) dot(b vec) int { return a.x*b.x + a.y*b.y }

// a.det(b) > 0 => a 到 b 逆时针
// a.det(b) < 0 => a 到 b 顺时针
func (a vec) det(b vec) int { return a.x*b.y - a.y*b.x } // 如果乘法会溢出，用 detCmp
func (a vec) detCmp(b vec) int {
	v := new(big.Int).Mul(big.NewInt(int64(a.x)), big.NewInt(int64(b.y)))
	w := new(big.Int).Mul(big.NewInt(int64(a.y)), big.NewInt(int64(b.x)))
	return v.Cmp(w)
}

// --- 泛型双端队列 --- //

type deque[T any] struct{ l, r []T }

func (q deque[T]) empty() bool {
	return len(q.l) == 0 && len(q.r) == 0
}

func (q deque[T]) size() int {
	return len(q.l) + len(q.r)
}

func (q *deque[T]) pushFront(v T) {
	q.l = append(q.l, v)
}

func (q *deque[T]) pushBack(v T) {
	q.r = append(q.r, v)
}

func (q *deque[T]) popFront() (v T) {
	if len(q.l) > 0 {
		q.l, v = q.l[:len(q.l)-1], q.l[len(q.l)-1]
	} else {
		v, q.r = q.r[0], q.r[1:]
	}
	return
}

func (q *deque[T]) popBack() (v T) {
	if len(q.r) > 0 {
		q.r, v = q.r[:len(q.r)-1], q.r[len(q.r)-1]
	} else {
		v, q.l = q.l[0], q.l[1:]
	}
	return
}

func (q deque[T]) front() T {
	if len(q.l) > 0 {
		return q.l[len(q.l)-1]
	}
	return q.r[0]
}

func (q deque[T]) back() T {
	if len(q.r) > 0 {
		return q.r[len(q.r)-1]
	}
	return q.l[0]
}

// 0 <= i < q.size()
func (q deque[T]) get(i int) T {
	if i < len(q.l) {
		return q.l[len(q.l)-1-i]
	}
	return q.r[i-len(q.l)]
}

// --- 下凸包（求最小值） --- //
// lh := LowerHull

type LowerHull struct {
	hull deque[vec]
}

func (lh *LowerHull) add(p vec) {
	for lh.hull.size() > 1 {
		back := lh.hull.back()
		prev := lh.hull.get(lh.hull.size() - 2)
		if back.sub(prev).detCmp(p.sub(back)) <= 0 {
			lh.hull.popBack()
		} else {
			break
		}
	}
	lh.hull.pushBack(p)
}

// 保证 v0.x 单调
// 如果 v0.x 单调递增，则 LowerHull.queryMonotonic(v0,  1)
// 如果 v0.x 单调递减，则 LowerHull.queryMonotonic(v0, -1)
// 复杂度 O(n)
func (lh *LowerHull) queryMonotonic(p vec, dir int) int {
	if dir < 0 {
		for lh.hull.size() > 1 && p.dot(lh.hull.get(0)) >= p.dot(lh.hull.get(1)) {
			lh.hull.popFront()
		}
		return p.dot(lh.hull.front())
	} else {
		for lh.hull.size() > 1 && p.dot(lh.hull.back()) >= p.dot(lh.hull.get(lh.hull.size()-2)) {
			lh.hull.popBack()
		}
		return p.dot(lh.hull.back())
	}
}

// 二分查询最小值，复杂度 O(nlogn)
func (lh *LowerHull) queryBinary(p vec) int {
	l, r := 0, lh.hull.size()-1
	for l < r {
		mid := (l + r) >> 1
		if p.dot(lh.hull.get(mid)) >= p.dot(lh.hull.get(mid+1)) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return p.dot(lh.hull.get(l))
}

func (lh *LowerHull) empty() bool {
	return lh.hull.empty()
}

func (lh *LowerHull) clear() {
	lh.hull = deque[vec]{}
}

// addFront 向下凸包头部插入元素（用于启发式合并）
func (lh *LowerHull) addFront(p vec) {
	for lh.hull.size() > 1 {
		front := lh.hull.front()
		next := lh.hull.get(1)

		if front.sub(p).detCmp(next.sub(front)) <= 0 {
			lh.hull.popFront()
		} else {
			break
		}
	}
	lh.hull.pushFront(p)
}

// merge 启发式合并两个下凸包
// 要求：lh1 中的所有点的 x 坐标 严格小于 lh2 中所有点的 x 坐标
func merge(lh1, lh2 LowerHull) LowerHull {
	if lh1.hull.size() <= lh2.hull.size() {
		for !lh1.empty() {
			lh2.addFront(lh1.hull.popBack())
		}
		return lh2
	} else {
		for !lh2.empty() {
			lh1.add(lh2.hull.popFront())
		}
		return lh1
	}
}

// ======= 凸包 模板结束 ======= //