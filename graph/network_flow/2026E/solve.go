// https://codeforces.com/problemset/problem/2026/E
// 2500

// 最大权闭合子图

package main

import (
	"math/bits"
)

func solve() {
	n := II()
	all := 0
	nums := make([]int, n)
	for i := range n {
		x := II()
		nums[i] = x
		all |= x
	}
	tot := bits.Len(uint(all))
	edges := [][]int{}
	src, dst := n + tot, n + tot + 1
	for b := range tot {
		flag := false
		for i, x := range nums {
			if x >> b & 1 > 0 {
				edges = append(edges, []int{n + b, i, inf})
				flag = true
			}
		}
		if flag {
			edges = append(edges, []int{src, n + b, 1})
		}
	}
	for i := range n {
		edges = append(edges, []int{i, dst, 1})
	}
	cut := dinic(n + tot + 2, edges, src, dst)
	Println(n - cut)
}

func dinic(n int, edges [][]int, s, t int) (maxFlow int) {
	type edge struct { to, wt, rev int }
	g := make([][]edge, n)
	for _, e := range edges {
		u, v, wt := e[0], e[1], e[2]
		g[u] = append(g[u], edge{v, wt, len(g[v])})
		g[v] = append(g[v], edge{u, 0, len(g[u]) - 1}) // 无向图将 0 改成 wt
	}

	level := make([]int, n)

	bfs := func() bool {
		for i := range level {
			level[i] = -1
		}
		level[s] = 0
		q := []int{s}
		for len(q) > 0 {
			u := q[0]
			q = q[1:]
			for _, e := range g[u] {
				v, wt := e.to, e.wt
				if level[v] < 0 && wt > 0 {
					level[v] = level[u] + 1
					q = append(q, v)
				}
			}
		}
		return level[t] > 0
	}

	var cur = make([]int, n)

	var dfs func(int, int) int
	dfs = func(u, flow int) (pushed int) {
		if u == t || flow == 0 {
			return flow
		}

		for ; cur[u] < len(g[u]); cur[u]++ {
			e := &g[u][cur[u]]
			v, wt, rev := e.to, e.wt, e.rev

			if level[v] == level[u] + 1 && wt > 0 {
				push := dfs(v, min(flow - pushed, wt))
				if push > 0 {
					e.wt -= push
					g[v][rev].wt += push
					pushed += push
					
					if pushed == flow {
						break
					}
				}
			}
		}
		if pushed == 0 {
			level[u] = -1
		}
		return
	}

	for bfs() {
		cur = make([]int, n)
		for {
			push := dfs(s, inf)
			if push == 0 {
				break
			}
			maxFlow += push
		}
	}
	return maxFlow
}
