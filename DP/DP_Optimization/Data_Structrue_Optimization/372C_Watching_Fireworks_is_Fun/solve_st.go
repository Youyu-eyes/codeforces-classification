// https://codeforces.com/contest/372/problem/C
// 2100
// RMQ 优化 DP
// 乍一看很难看出转移，本题需要把状态定义出来之后才能很好的看到
// 定义 f[i][p] 为放第 i 次烟花，人在 p 位置的最大开心度
// 思考上一次放烟花，假设在 q 位置，那么 p 与 q 满足：
// 如果 q 在 p 左边：
// q = p - (t[i] - t[i - 1]) * d
// 如果 q 在 p 右边
// q = p + (t[i] - t[i - 1]) * d
// 左边与 1 取最大，右边与 n 取最小，保证下标不越界，则我们可以得到 q 的取值范围
// q in [max(1, p - (t[i] - t[i - 1]) * d), min(n, p + (t[i] - t[i - 1]) * d)]

// 状态转移方程
// f[i][p] = max(f[i - 1][q]) + b[i] - abs(a[i] - p)
// 滚动数组空间优化
// nf[p] = max(f[q]) + b[i] - abs(a[i] - p)
// q 的取值是一段连续区间，可以采用 ST 表优化 RMQ 的过程
// 在更新 nf 之前，对 f 数组建 ST 表

// 普通 ST 表也能过，但是 ST_fast 更快
// 本题采用普通 ST 表
// 普通 ST 表见数据结构模板，ST_fast 见 Go 语言算法模板库

// 时间复杂度 O(n·m·logn)
// 空间复杂度 O(n)

// 另有双端队列做法，见下

package main

import (
	"math/bits"
	"slices"
)

func solveST() {
	n := II()
	m := II()
	d := II()

	type Firework struct{ a, b, t int }
	fw := make([]Firework, m)
	for i := 0; i < m; i++ {
		fw[i].a = II()
		fw[i].b = II()
		fw[i].t = II()
	}

	f := make([]int, n + 1)
	for p := 1; p <= n; p++ {
		f[p] = fw[0].b - abs(fw[0].a - p)
	}

	for i := 1; i < m; i++ {
		st := newSparseTable(f, func(a, b int) int { return max(a, b) })
		dist := (fw[i].t - fw[i - 1].t) * d
		nf := make([]int, n+1)
		for p := 1; p <= n; p++ {
			best := st.query(max(1, p - dist), min(n, p + dist) + 1)
			nf[p] = best + fw[i].b - abs(fw[i].a - p)
		}
		copy(f, nf)
	}

	Println(slices.Max(f[1:]))
}


// ------- 数值 ST 表 ------- //
type SparseTable[T any] struct {
	st [][]T
	op func(T, T) T
}

// op 需满足结合律（如 min, max, gcd 等）
func newSparseTable[T any](nums []T, op func(T, T) T) *SparseTable[T] {
	n := len(nums)
	w := bits.Len(uint(n))
	st := make([][]T, w)
	for i := range st {
		st[i] = make([]T, n)
	}
	copy(st[0], nums)
	for i := 1; i < w; i++ {
		step := 1 << (i - 1)
		for j := 0; j+step < n && j+(1<<i) <= n; j++ {
			st[i][j] = op(st[i-1][j], st[i-1][j+step])
		}
	}
	return &SparseTable[T]{st: st, op: op}
}

func (st *SparseTable[T]) query(l, r int) T {
	k := bits.Len(uint(r-l)) - 1
	return st.op(st.st[k][l], st.st[k][r-1<<k])
}
