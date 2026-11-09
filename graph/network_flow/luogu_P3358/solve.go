// https://www.luogu.com.cn/problem/P3358
// 洛谷紫 - 省选/NOI-

// 网络流 24 题

// 某一点不能被覆盖 k 次
// 可以转换为选 k 组不重叠的边
// 将区间左端点于右端点（离散化下标）连边，容量为 1 表示只能选一次，费用为距离的负数

// 将 x 轴离散化后相邻的连边，容量为 k 或 inf 表示可以任意选，费用为 0 表示跳过这个区间
// 设立超级源点 src，与源点 0 连边，容量为 k，费用为 0，表示总共选 k 条边
// 建立超级汇点 dst，与汇点 m-1 连边，容量为 k，费用为 0，作用同理

// 最后跑最小费用最大流，最小费用的相反数即为最大区间和

package main

import (
	"slices"
	"sort"
	"container/heap"
)

type inter struct {
	l, r int
}

func solve() {
	n, k := II(), II()
	inters := make([]inter, n)
	xs := make([]int, 0, 2*n)
	for i := range n {
		l, r := II(), II()
		inters[i] = inter{l, r}
		xs = append(xs, l, r)
	}
	slices.Sort(xs)
	sorted := slices.Compact(xs)
	m := len(sorted)

	src, dst := m, m+1

	edges := [][]int{}

	add := func(u, v, cap, cost int) {
		edges = append(edges, []int{u, v, cap, cost})
	}
	for _, p := range inters {
		l := sort.SearchInts(sorted, p.l)
		r := sort.SearchInts(sorted, p.r)
		add(l, r, 1, -(p.r - p.l))
	}
	for i := range len(sorted)-1 {
		add(i, i+1, k, 0)
	}
	add(src, 0, k, 0)
	add(m-1, dst, k, 0)
	_, cost := zkw(m+2, edges, src, dst)
	Println(-cost)
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

    // spfa 求初始势能函数
    for i := range hi {
        hi[i] = inf
    }
    hi[s] = 0
    inq := make([]bool, n)
    q := []int{s}
    inq[s] = true
    for len(q) > 0 {
        u := q[0]
        q = q[1:]
        inq[u] = false
        for _, e := range g[u] {
            if e.wt > 0 && hi[u]+e.ct < hi[e.to] {
                hi[e.to] = hi[u] + e.ct
                if !inq[e.to] {
                    q = append(q, e.to)
                    inq[e.to] = true
                }
            }
        }
    }

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

type pair struct{ dis, x int }
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }