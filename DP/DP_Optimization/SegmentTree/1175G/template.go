package main

import (
	"math/big"
)

type line struct {
	k, b float64
	id   int
}

func newLine() line {
	return line{k: 0, b: inf, id: 0}
	// 最大值：return line{k: 0, b: -inf, id: 0}
}

func (l line) calc(x int) float64 {
	return l.k*float64(x) + l.b
}

type pLcNode struct {
	lo, ro *pLcNode
	l, r   int
	line   line
}

func buildPlc(l, r int) *pLcNode {
	o := &pLcNode{l: l, r: r, line: newLine()}
	if l == r {
		return o
	}
	m := (l + r) >> 1
	o.lo = buildPlc(l, m)
	o.ro = buildPlc(m+1, r)
	return o
}

func (o pLcNode) updateLine(nl line) *pLcNode {
	l, r := o.l, o.r
	m := (l + r) >> 1

	if nl.calc(m) < o.line.calc(m)-eps { // 最大值：if nl.calc(m) > o.line.calc(m)+eps
		o.line, nl = nl, o.line
	}
	if l == r {
		return &o
	}

	if nl.calc(l) < o.line.calc(l)-eps { // 最大值：if nl.calc(l) > o.line.calc(l)+eps
		o.lo = o.lo.updateLine(nl)
	} else if nl.calc(r) < o.line.calc(r)-eps { // 最大值：if nl.calc(r) > o.line.calc(r)+eps
		o.ro = o.ro.updateLine(nl)
	}
	return &o
}

// 插入线段 [ql, qr] 上的直线 nl
func (o pLcNode) insert(nl line, ql, qr int) *pLcNode {
	if ql > o.r || qr < o.l {
		return &o
	}
	if ql <= o.l && o.r <= qr {
		return o.updateLine(nl)
	}
	m := (o.l + o.r) >> 1
	if ql <= m {
		o.lo = o.lo.insert(nl, ql, qr)
	}
	if m < qr {
		o.ro = o.ro.insert(nl, ql, qr)
	}
	return &o
}

// 查询 x 处的最值
func (o *pLcNode) query(x int) (res float64) {
	if o == nil {
		return inf // 最大值：return -inf
	}
	res = inf // 最大值：res := -inf
	
	if o.line.id != 0 {
		res = o.line.calc(x)
	}
	if o.l == o.r {
		return res
	}
	m := (o.l + o.r) >> 1
	if x <= m {
		return min(res, o.lo.query(x)) // 最大值：return max(res, o.lo.query(x))
	}
	return min(res, o.ro.query(x)) // 最大值：return max(res, o.ro.query(x))
}

func newPlc(l, r int) []*pLcNode {
	return []*pLcNode{buildPlc(l, r)}
}

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
