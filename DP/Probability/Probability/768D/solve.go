// https://codeforces.com/problemset/problem/768/D
// 2200
// 概率DP + 离线

// 定义 f[i][j] 为第 i 天生成了 j 个小球的概率
// 如果第 i - 1 天有 j     个小球，说明生成的仍然是 j 个小球中的一个，概率为 j/k
// 如果第 i - 1 天有 j - 1 个小球，说明生成的是 j - 1 个小球的另外一个，概率为 (k-j+1)/k

// 状态转移方程为
// f[i][j] = f[i - 1][j] * j / k + f[i - 1][j - 1] * (k - j + 1) / k

// 如果对于每一个查询都从第 0 天开始跑 DP 的话，会超时
// 注意到，随着 i 增大，f[i][k] 是递增的，因此我们可以对 p 排序，p 小的会先满足

// 时间复杂度：O(U × k) U 表示最大天数
// 空间复杂度：O(max(k, q))

package main

import (
	"slices"
	"cmp"
)

type Q struct {p float64; idx int}

func solve() {
	k, q := II(), II()
	queries := make([]Q, q)
	for i := range q {
		queries[i] = Q{IF(), i}
	}
	
	slices.SortFunc(queries, func(a, b Q) int {
		return cmp.Compare(a.p, b.p)
	})

	ans := make([]int, q)
	i := 0
	
	f := make([]float64, k + 1)
	f[0] = 1.0
	for _, query := range queries {
		for f[k] < query.p / float64(2000) {
			i++
			for j := k; j > 0; j-- {
				f[j] = f[j] * float64(j) / float64(k) + f[j - 1] * float64(k - j + 1) / float64(k)
			}
			f[0] = 0
		}
		ans[query.idx] = i
	}
	
	for _, res := range ans {
		Println(res)
	}
}