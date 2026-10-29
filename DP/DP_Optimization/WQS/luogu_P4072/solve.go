// https://www.luogu.com.cn/problem/P4072
// 洛谷蓝 - 提高+/省选-

// 斜率优化 DP + WQS 二分
// 普通划分型 DP + 斜率优化 DP 见斜率优化 DP 题单

package main

import (
	"math/big"
)

type pair struct {
	x, b int
}

func solve() {
	n, limit := II(), II()
	s := make([]int, n+1)
	right := 1
	for i := range n {
		x := II()
		s[i+1] = s[i] + x
		right += x*x
	}
	left := -right

	dp := func(k int) pair {
		lh := LowerHull{}
		lh.add(vec{0, 0}, 0)
		f := make([]pair, n+1)
		for i := 1; i <= n; i++ {
			best, j := lh.queryMonotonic(vec{-2*s[i], 1})
			f[i] = pair{f[j].x+1, best+s[i]*s[i]-k}
			lh.add(vec{s[i], f[i].b+s[i]*s[i]}, i)
		}
		return f[n]
	}

	var ans int
	for left + 1 < right {
		mid := (left + right) >> 1
		res := dp(mid)
		if res.x <= limit {
			ans = mid*limit+res.b
			left = mid
		} else {
			right = mid
		}
	}

	Println(limit*ans-s[n]*s[n])
}

type vec struct{ x, y int }

func (a vec) sub(b vec) vec { return vec{a.x - b.x, a.y - b.y} }
func (a vec) dot(b vec) int { return a.x*b.x + a.y*b.y }

// a.det(b) > 0 => a 到 b 逆时针
// a.det(b) < 0 => a 到 b 顺时针
func (a vec) det(b vec) int { return a.x*b.y - a.y*b.x }

func (a vec) detCmp(b vec) int {
	v := new(big.Int).Mul(big.NewInt(int64(a.x)), big.NewInt(int64(b.y)))
	w := new(big.Int).Mul(big.NewInt(int64(a.y)), big.NewInt(int64(b.x)))
	return v.Cmp(w)
}

type point struct {
	v vec
	idx int
}

type LowerHull struct {
	hull []point
}

func (lh *LowerHull) add(p vec, i int) {
	for len(lh.hull) > 1 {
		back := lh.hull[len(lh.hull)-1].v
		prev := lh.hull[len(lh.hull)-2].v
		if back.sub(prev).detCmp(p.sub(back)) <= 0 {
			lh.hull = lh.hull[:len(lh.hull)-1]
		} else {
			break
		}
	}
	lh.hull = append(lh.hull, point{p, i})
}

// v0.x 单调递减
func (lh *LowerHull) queryMonotonic(p vec) (int, int) {
	for len(lh.hull) > 1 && p.dot(lh.hull[0].v) > p.dot(lh.hull[1].v) {
		lh.hull = lh.hull[1:]
	}
	return p.dot(lh.hull[0].v), lh.hull[0].idx
}
