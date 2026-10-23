// https://codeforces.com/contest/2117/problem/G
// 1900
// 最短路，并查集

// 核心思想是用并查集维护连通块的边权最值
// 当 0 和 n - 1 在一个连通块内时，可以更新答案

// 需要按边权从小到大连边，保证最大边权递增
// 如果不从小到大连边，则会出现先连了一条无用的边权较大的边，使得连通块最大边权虚大
// 如果先连了一条无用的边权较小的边，边权较大的边有用，则不会影响答案

package main

import (
	"slices"
	"cmp"
)

type edge struct {
	o1 int
	o2 int
	wt int
}

func solve() {
	n, m := II(), II()
	uf := newUnionFind(n)
	edges := make([]edge, m)

	for i := range m {
		u, v, w := II() - 1, II() - 1, II()
		edges[i] = edge{u, v, w}
	}

	slices.SortFunc(edges, func(a, b edge) int {
		return cmp.Compare(a.wt, b.wt)
	})

	ans := inf

	for _, e := range edges {
		uf.merge(e.o1, e.o2, e.wt)
		if uf.same(0, n - 1) {
			ans = min(ans, uf.get(0))
		}
	}
	Println(ans)
}

// ------- 并查集 ------- //
// 初始化：uf := newUnionFind(n)

type unionFind struct {
	fa []int // 代表元
	sz []int // 集合大小
	mx []int // 连通块边权最大值
	mn []int // 连通块边权最小值
	cc int   // 连通块个数
}

func newUnionFind(n int) unionFind {
	fa := make([]int, n)
	sz := make([]int, n)
	mx := make([]int, n)
	mn := make([]int, n)
	for i := range fa {
		fa[i] = i
		sz[i] = 1
		mx[i] = 0
		mn[i] = inf
	}
	return unionFind{fa, sz, mx, mn, n}
}

func (u unionFind) find(x int) int {
	if u.fa[x] != x {
		u.fa[x] = u.find(u.fa[x])
	}
	return u.fa[x]
}

func (u unionFind) same(x, y int) bool {
	return u.find(x) == u.find(y)
}

func (u *unionFind) merge(from, to int, weight int) bool {
	x, y := u.find(from), u.find(to)
	if x == y {
		u.mx[x] = max(u.mx[x], weight)
		u.mn[x] = min(u.mn[x], weight)
		return false
	}

	// 按秩合并
	if u.sz[x] > u.sz[y] {
		x, y = y, x
	}

	u.fa[x] = y
	u.sz[y] += u.sz[x]
	u.mx[y] = max(max(u.mx[x], u.mx[y]), weight)
	u.mn[y] = min(min(u.mn[x], u.mn[y]), weight)
	u.cc--
	return true
}

func (u unionFind) size(x int) int {
	return u.sz[u.find(x)]
}

func (u unionFind) get(x int) int {
	f := u.find(x)
	return u.mx[f] + u.mn[f]
}