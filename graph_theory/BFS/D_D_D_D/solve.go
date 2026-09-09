// https://codeforces.com/problemset/problem/2109/D
// 1900
// 贪心 + BFS 最短路

// 我们需要从 A 中选择一些数使得这些数的和为路径长度
// 本题的关键在于图是无向图，且允许走重复路径
// 这就意味着虽然我们不能选择比 0 到 i 的最短路更小的子集和
// 但是我们可以选一个更大的子集和，这样会多出一些步数

// 从 i 出发到任意一个端点，来回走，这样走偶数步就能回到 i 点
// 可以用这样的思想消耗掉多余的步数
// 这样带来的要求是：多余步数一定为偶数

// 多余步数 = 子集和 - 最短路长度

// 我们要求子集和尽可能大，同时要求子集和与最短路长度同奇偶
// 则我们维护 最大奇数子集和 与 最大偶数子集和
// 同时维护   最小奇数最短路 与 最小偶数最短路
// 奇数或偶数子集和比对应的最短路长度大，即满足要求
// 注意这里必须维护两种最短路长度，因为最短路多走一步，可能子集和就能多选择一个更大的奇数来改变奇偶性

// 1. 子集和
// 首先我们算 A 集合的和 total，分类讨论：
// - 如果 A 集合的和为偶数，则 mxEven = total
//   - 此时 mxOdd 为整个集合中删除一个最小的奇数，即 total - minOdd
// - 如果 A 集合的和为奇数，则 mxOdd = total
//   - 此时 mxEven 为整个集合中删除一个最小的奇数，即 total - minOdd

// 2. 最短路
// 本题无边权，相当于所有边权全部为 1
// 直接用 BFS 即可
// 注意多用一个维度维护奇偶性
// 相当于分层图最短路

// 复杂度
// 时间复杂度：O(n + m + l)，其中 n 表示点的个数，m 表示边的条数，l 表示集合大小
// 空间复杂度：O(n)

package main

func solve() {
	n, m, l := II(), II(), II()
	total, minOdd := 0, inf
	for range l {
		x := II()
		total += x
		if x & 1 == 1 {
			minOdd = min(minOdd, x)
		}
	}
	mxOdd, mxEven := 0, 0
	if total & 1 > 0 {
		mxOdd = total
		mxEven = total - minOdd
	} else {
		mxEven = total
		mxOdd = total - minOdd
	}

	g := make([][]int, n)
	for range m {
		x, y := II() - 1, II() - 1
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	dis := make([][2]int, n)
	for i := range dis {
		dis[i] = [2]int{-1, -1}
	}
	dis[0][0] = 0

	q := []int{0}
	d := 1
	for len(q) > 0 {
		tmp := make([]int, len(q))
		copy(tmp, q)
		q = q[:0]

		for _, u := range tmp {
			for _, v := range g[u] {
				if dis[v][d & 1] < 0 {
					dis[v][d & 1] = d
					q = append(q, v)
				}
			}
		}
		d++
	}

	ans := make([]byte, n)
	for i := range n {
		ans[i] = '0'
		if dis[i][0] >= 0 && mxEven >= dis[i][0] {
			ans[i] = '1'
		}
		if dis[i][1] >= 0 && mxOdd >= dis[i][1] {
			ans[i] = '1'
		}
	}

	Println(string(ans))
}
