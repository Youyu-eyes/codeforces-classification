// https://www.luogu.com.cn/problem/P3381
// 最小费用最大流
// 费用流

package main

import (
	"container/heap"
)

func solve() {
	n, m, s, t := II(), II(), II() - 1, II() - 1
	edges := make([][]int, m)
	for i := range m {
		edges[i] = []int{II() - 1, II() - 1, II(), II()}
	}
	maxFlow, minCost := zkw(n, edges, s, t)
	Println(maxFlow, minCost)
}

func zkw(n int, edges [][]int, s, t int) (maxFlow int, minCost int) {
	type edge struct{ to, wt, ct, rev int }
	g := make([][]edge, n)
	for _, e := range edges {
		x, y, wt, ct := e[0], e[1], e[2], e[3]
		g[x] = append(g[x], edge{y, wt, ct, len(g[y])})
		g[y] = append(g[y], edge{x, 0, -ct, len(g[x]) - 1})
	}

	dis := make([]int, n)
	hi := make([]int, n)

	dijkstra := func() bool {
		for i := range dis {
			dis[i] = inf
		}
		dis[s] = 0
		h := &hp{{0, s}}

		for h.Len() > 0 {
			p := heap.Pop(h).(pair)
			disX, x := p.dis, p.x
			if disX > dis[x] {
				continue
			}
			for _, e := range g[x] {
				y := e.to
				if e.wt > 0 {
					newDisY := disX + e.ct + hi[x] - hi[y]
					if newDisY < dis[y] {
						dis[y] = newDisY
						heap.Push(h, pair{newDisY, y})
					}
				}
			}
		}
		return dis[t] < inf
	}

	var cur = make([]int, n)
	vis := make([]bool, n)

    var dfs func(int, int) int
	dfs = func(u, flow int) (pushed int) {
		if u == t || flow == 0 {
			return flow
		}
		vis[u] = true
		
		for ; cur[u] < len(g[u]); cur[u]++ {
			e := &g[u][cur[u]]
			v, wt, ct, rev := e.to, e.wt, e.ct, e.rev
			
			if !vis[v] && wt > 0 && ct + hi[u] - hi[v] == 0 {
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
		vis[u] = false
		return
	}

	for dijkstra() {
		for i := range n {
			if dis[i] < inf {
				hi[i] += dis[i]
			}
			cur[i] = 0
		}

		for {
			push := dfs(s, inf)
			if push == 0 {
				break
			}
			maxFlow += push
			minCost += push * hi[t]
		}
	}
	return
}

// 需要用堆的时候，需要定义的东西
type pair struct{ dis, x int }
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }