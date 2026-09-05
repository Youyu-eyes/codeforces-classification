package main

import (
	"math/big"
)

func solve() {
	n, L := II(), II()
	nums := make([]int, n)
	for i := range nums {
		nums[i] = II()
	}

	ans := vec{1, -inf}
	s := make([]int, n + 1)
	lh := LowerHull{}
	lh.add(vec{0, 0})
	for i, x := range nums {
		s[i + 1] = s[i] + x
		point := vec{i + 1, s[i + 1]}

		if i - L + 1 >= 0 {
			mx := lh.queryBinary(point)
			cur := vec{point.x - mx.x, point.y - mx.y}
			if ans.detCmp(cur) > 0 {
				ans = cur
			}
			lh.add(vec{i - L + 2, s[i - L + 2]})
		}
	}
	Println(float64(ans.y)/float64(ans.x))
}

// 在使用模板时要注意：
// 首先，在维护凸包时要保证查询向量 v0 的 y 值 v0.y > 0，这样我们才能用 max or min 简单地判断用 上凸包 or 下凸包
// 其次，在维护凸包时要保证待维护向量 v1 的 x 值 v1.x 单调递增

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
// lh := LowerHull{}

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

// 二分查询最小值，复杂度 O(nlogn)
func (lh *LowerHull) queryBinary(p vec) vec {
	l, r := 0, lh.hull.size()-1
	for l < r {
		mid := (l + r) >> 1
		m0 := lh.hull.get(mid)
		m1 := lh.hull.get(mid+1)

		dy0, dx0 := p.y - m0.y, p.x - m0.x
		dy1, dx1 := p.y - m1.y, p.x - m1.x
		if dy1 * dx0 >= dy0 * dx1 {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return lh.hull.get(l)
}