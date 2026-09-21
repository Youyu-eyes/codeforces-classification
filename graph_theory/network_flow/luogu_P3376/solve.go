// https://www.luogu.com.cn/problem/P3376

package main

func solve() {
	n, m, s, t := II(), II(), II() - 1, II() - 1

	type edge struct { to, wt, rev int }
	g := make([][]edge, n)
	for range m {
		u, v, wt := II() - 1, II() - 1, II()
		g[u] = append(g[u], edge{v, wt, len(g[v])})
		g[v] = append(g[v], edge{u,  0, len(g[u]) - 1})
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
	dfs = func(u, flow int) int {
		if u == t || flow == 0 {
			return flow
		}
		for i := cur[u]; i < len(g[u]); i++ {
			cur[u] = i
			e := &g[u][i]
			v, wt, rev := e.to, e.wt, e.rev
			if level[v] == level[u] + 1 && wt > 0 {
				push := dfs(v, min(flow, wt))
				if push > 0 {
					e.wt -= push
					g[v][rev].wt += push
					return push
				}
			}
		}
		return 0
	}

	maxFlow := 0
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
	Println(maxFlow)
}