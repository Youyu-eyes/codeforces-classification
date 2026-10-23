// https://www.luogu.com.cn/problem/U103816
// 最大权闭合子图

package main

func solve() {
	n, m := II(), II()
	weights := make([]int, n)
	for i := range n {
		weights[i] = II()
	}
	edges := [][]int{}
	for range m {
		edges = append(edges, []int{II() - 1, II() - 1, inf})
	}
	src := n
	dst := n + 1
	ans := 0
	for i, w := range weights {
		if weights[i] >= 0 {
			edges = append(edges, []int{src, i, w})
			ans += w
		} else {
			edges = append(edges, []int{i, dst, -w})
		}
	}
	mnCut := dinic(n + 2, edges, src, dst)
	Println(ans - mnCut)
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
