// https://vjudge.net/problem/HDU-7906

package main

import (
	"math/big"
)

func solve() {
	t, n, m := II(), II(), II()
	a := make([]int, n + 1)
	for i := range n - 1 {
		a[i + 1] = II()
	}
	a[n] = n * t

	nums := make([]int, n)
	nums[0] = a[1]
	for i := 1; i < n; i++ {
		nums[i] = a[i + 1] - a[i]
	}
	
	s := make([]int, n + 1)
	for i := range n {
		s[i + 1] = s[i] + nums[i]
	}

	f := make([][]int, n*t+1)
	for x := range f {
		f[x] = make([]int, m + 1)
		for c := range m + 1 {
			f[x][c] = inf
		}
	}

	f[0][0] = 0
	for i := 1; i <= n; i++ {
		nf := make([][]int, n*t+1)
		for x := range nf {
			nf[x] = make([]int, m + 1)
			for c := range m + 1 {
				nf[x][c] = inf
			}
		}
		
		for c := range m + 1 {
			lh := LowerHull{}
			for x := range n*t+1 {
				lh.add(vec{x, f[x][c] + x * x})
				best := lh.queryMonotonic(vec{-2 * x, 1})
				nc := c + abs(s[i] - x)
				if nc <= m {
					nf[x][nc] = min(nf[x][nc], best + x * x)
				}
			}
		}
		
		f = nf
	}

	ans := inf
	for c := range m + 1 {
		ans = min(ans, f[n*t][c])
	}
	Println(ans)
}

// v0.y > 0
// v1.x 单调递增

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

// --- 下凸包 --- //
// lh := LowerHull{}

type LowerHull struct {
	hull []vec
}

func (lh *LowerHull) add(p vec) {
	for len(lh.hull) > 1 {
		back := lh.hull[len(lh.hull)-1]
		prev := lh.hull[len(lh.hull)-2]
		if back.sub(prev).detCmp(p.sub(back)) <= 0 {
			lh.hull = lh.hull[:len(lh.hull)-1]
		} else {
			break
		}
	}
	lh.hull = append(lh.hull, p)
}

// v0.x 单调递减
func (lh *LowerHull) queryMonotonic(p vec) int {
	for len(lh.hull) > 1 && p.dot(lh.hull[0]) >= p.dot(lh.hull[1]) {
		lh.hull = lh.hull[1:]
	}
	return p.dot(lh.hull[0])
}
