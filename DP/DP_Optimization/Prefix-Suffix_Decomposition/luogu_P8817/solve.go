// https://www.luogu.com.cn/problem/P8817
// BFS + 前后缀分解

// 首先，本题是图论，但是一些图论专属的算法（最短路，网络流，并查集等）都和本题没什么太大的关系
// 且有环，不考虑 DP

// 观察到我们只选 4 个点，路径为 1 -> A -> B -> C -> D -> 1
// 将路径从 B C 中间断开，考虑合法的 A B 点和 C D 点
// 枚举 B 点，考虑合法的 A 点
// 合法的定义是 dis[0][A] <= k && dis[A][B] <= k
// 其中 dis[i][j] 定义为从 i 到 j 的距离，本题的距离定义为途径点数，这可以用 BFS 线性求解

// 由于 n <= 2500，因此我们暴力 O(n^2) 枚举 B, A，对于 B 而言，用 pre[B] 保存前 4 大的 score(A) 的 A 点
// 本题是无向图，因此无需额外对 C D 点进行处理
// 再次暴力枚举 B, C 两点，如果 B, C 合法，则枚举 A in pre[B] 和 D in pre[C]，如果 4 点均不同，则更新答案

// 为什么是保存 4 个点？
// 我们暴力枚举 B C 两点，此时 B C 两点合法，枚举 A 点，A B 合法，且保证 A 与 C 不同（如果相同则跳过）
// 此时需要从 pre[C] 中选一点 D，使得 D 与 A, B, C 三点均不同
// 根据鸽巢原理，pre[C] 中保存 4 个点，则必然有一点与 A B C 不同
// 且此时选出来的 A B C D 之间均合法

// 时间复杂度：O(n(n+m) + 4*n^2) 其中 n 表示点的数量，m 表示边的数量

package main

func solve() {
	n, m, k := II(), II(), II()
	score := make([]int, n)
	for i := 1; i < n; i++ {
		score[i] = II()
	}

	g := make([][]int, n)
	for range m {
		x, y := II() - 1, II() - 1
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	dis := make([][]bool, n)
	for i := range n {
		dis[i] = make([]bool, n)
	}

	for i := range n {
		q := []int{i}
		d := 0
		for len(q) > 0 && d <= k {
			tmp := q
			q = []int{}
			for _, x := range tmp {
				for _, y := range g[x] {
					if !dis[i][y] {
						q = append(q, y)
						dis[i][y] = true
					}
				}
			}
			d++
		}
	}

	pre := make([][]int, n)
	upd := func(nums []int, x int) []int {
		n := len(nums)
		res := make([]int, 0, 4)
		i := 0

		for i < n && score[nums[i]] >= score[x] {
			res = append(res, nums[i])
			i++
		}

		if len(res) < 4 {
			res = append(res, x)
		}

		for i < n && len(res) < 4 {
			res = append(res, nums[i])
			i++
		}
		return res
	}

	for B := 1; B < n; B++ {
		for A := 1; A < n; A++ {
			if A == B {
				continue
			}
			if dis[0][A] && dis[A][B] {
				pre[B] = upd(pre[B], A)
			}
		}
	}

	ans := 0
	for B := 1; B < n; B++ {
		for C := 1; C < n; C++ {
			if B == C {
				continue
			}
			if dis[B][C] {
				for _, A := range pre[B] {
					if A == C {
						continue
					}
					for _, D := range pre[C] {
						if B == D || A == D {
							continue
						}
						ans = max(ans, score[A] + score[B] + score[C] + score[D])
					}
				}
			}
		}
	}
	Println(ans)
}
