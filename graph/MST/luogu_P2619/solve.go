// https://www.luogu.com.cn/problem/P2619
// 洛谷蓝 - 提高+/省选-

// 最小生成树 + WQS 二分
// 个人觉得可以给到省选/NOI-

// 定义 f(x) 为选择 x 条白边的最小总边权
// 显然，随着 x 递增，f(x) 先递减后递增
// 形成下凸包

// 定义切线 y = k*x + b
// b = f(x) - k*x
// 二分斜率，寻找斜率固定时 b 的最小值
// 每选一条白边，减去 k
// 注意比较的时候白边也需要 -k
// 否则每次选择的边都是一样的

// 还要注意 k = 0 时，f(x - 1) = f(x) = f(x + 1) = ...
// 此时会形成一条与 x 轴平行的线，如果 limit 在直线内
// 此时如果我们 check 函数内部找的是 x 最大的那个
// 那么计算的 x 在 limit 右侧
// 此时我们需要减少中点斜率，使 x 左移

package main

import (
	"slices"
)

func solve() {
	n, m, need := II(), II(), II()
	white := [][]int{}
	black := [][]int{}
	right := 1
	for range m {
		u, v, wt, color := II(), II(), II(), II()
		if color == 0 {
			white = append(white, []int{u, v, wt})
		} else {
			black = append(black, []int{u, v, wt})
		}
		right += wt
	}
	left := -right
	slices.SortFunc(white, func(a, b []int) int {
		return a[2] - b[2]
	})
	slices.SortFunc(black, func(a, b []int) int {
		return a[2] - b[2]
	})

	kruskal := func(k int) (int, int) {
		uf := newUnionFind(n)
		sumWt := 0
		numWhite := 0
		i, j := 0, 0
		for uf.cc > 1 {
			for i < len(white) && uf.same(white[i][0], white[i][1]) {
				i++
			}
			for j < len(black) && uf.same(black[j][0], black[j][1]) {
				j++
			}
			if i < len(white) && (j == len(black) || white[i][2]-k < black[j][2]) {
				uf.merge(white[i][0], white[i][1])
				sumWt += white[i][2]
				numWhite++
				i++
			} else if j < len(black) {
				uf.merge(black[j][0], black[j][1])
				sumWt += black[j][2]
				j++
			}
		}

		return numWhite, sumWt - k*numWhite
	}

	var ans int
	for left + 1 < right {
		mid := (left + right) >> 1
		x, b := kruskal(mid)
		if x <= need {
			ans = mid*need + b
			left = mid
		} else {
			right = mid
		}
	}

	Println(ans)
}

type unionFind struct {
	fa []int // 代表元
	sz []int // 集合大小
	cc int   // 连通块个数
}

func newUnionFind(n int) unionFind {
	fa := make([]int, n)
	sz := make([]int, n)
	for i := range fa {
		fa[i] = i
		sz[i] = 1
	}
	return unionFind{fa, sz, n}
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

func (u *unionFind) merge(from, to int) bool {
	x, y := u.find(from), u.find(to)
	if x == y {
		return false
	}

	// 按秩合并
	if u.sz[x] > u.sz[y] {
	    x, y = y, x
	}

	u.fa[x] = y
	u.sz[y] += u.sz[x]
	u.cc--
	return true
}

func (u unionFind) size(x int) int {
	return u.sz[u.find(x)]
}
