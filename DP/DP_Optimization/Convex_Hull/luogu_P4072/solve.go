// https://www.luogu.com.cn/problem/P4072
// 洛谷蓝 - 提高+/省选-

// 凸包优化 DP
// 将方差的式子展开并简化后
// 只有 \sum(x_{i}^2) 不是常数，且前面的系数为正
// 我们需要维护上述式子的最小值
// 考虑 f[i] 为处理完前 i 个数的最少平方和
// 定义 s[i+1] = s[i] + nums[i] 为前缀和

// 枚举左端点 j
// 则 x[i] = s[i] - [j]
// nf[i] = min(f[j] + (s[i] - s[j])^2)
// 展开后可写成向量点积的形式

// 另有 WQS 二分做法，见 WQS 二分题单

package main

import (
	"math/big"
)

func solve() {
	n, m := II(), II()
	s := make([]int, n+1)
	for i := range n {
		s[i+1] = s[i] + II()
	}
	f := make([]int, n+1)
	for i := range n+1 {
		f[i] = inf
	}
	f[0] = 0
	for k := 1; k <= m; k++ {
		nf := make([]int, n+1)
		lh := LowerHull{}
		lh.add(vec{s[k-1], f[k-1]+s[k-1]*s[k-1]})
		for i := k; i <= n-(m-k); i++ {
			v0 := vec{-2*s[i], 1}
			nf[i] = lh.queryMonotonic(v0) + s[i]*s[i]
			lh.add(vec{s[i], f[i]+s[i]*s[i]})
		}
		f = nf
	}
	Println(m*f[n] - s[n]*s[n])
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

func (lh *LowerHull) queryBinary(p vec) int {
	l, r := 0, len(lh.hull)-1
	for l < r {
		mid := (l + r) >> 1
		if p.dot(lh.hull[mid]) >= p.dot(lh.hull[mid+1]) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return p.dot(lh.hull[l])
}
