// https://codeforces.com/contest/617/problem/E
// 2200
// 莫队 + 前缀和 + 枚举右维护左

// 设想一下，如果没有多个查询，只有一个查询，你会做吗？

// 区间异或和可以用前缀异或和求解，然后问题就变成了
// pre[i + 1] ^ pre[j] = k
// 求 (j, i) 点对，j <= i
// 这就是两数之和，力扣全站第一题，枚举右维护左
// 我们需要在更新答案之后，将 pre[i + 1] ^ k 记录在 cnt 中

// 现在让我们思考如何计算多个查询

// 没有更新，多个查询相互独立，可以考虑离线算法
// 针对多个查询的离线优化，考虑莫队

// cnt 统计的区间中 pre[i] 的频率
// 因此在区间左移的时候，更新操作是一样的

// 考虑撤销操作，由于我们只维护频率，没有维护最值，可以撤销
// 为了避免出现 a[i] = k 的情况，我们需要先将 pre[i + 1] 从 cnt 中删除
// 然后再撤销 i + 1 对答案的贡献

// 分析复杂度
// 枚举右维护左是线性复杂度

// 考虑右端点的最大移动次数
// 向右移动 n，然后向左移动 n - B，再向右移动 n - 2 × B ……
// 则移动次数为等差数列，首项 0，末项 n，公差 B
// 项数 = (n - 0) / B + 1 = n / B + 1
// 求和 = n × (n / B + 1) / 2 = n^2 / 2B + n / 2

// 左端点每个查询最多移动 B，最多 qB

// 总时间为 qB + n × n / 2B + n / 2，n 为常数，则问题变成求 qB + n × n / 2B 的最小值
// 由对勾函数可知，B = n / sqrt(2q) 时最小，最小值为 (n × sqrt(2q))

// 实际操作中，B = n / sqrt(2q) 可能会因为太小被卡时间或内存
// 建议令 B = n / sqrt(q)，然后两种都尝试一下

// 根号算法的时间还是比较玄学的，因为根号算法大约会比 log 时间大 10^2 左右
// 实际会到 10^7 或 10^8 数量级，所以与具体数据还是有比较大的关系的

// 时间复杂度：O(n / sqrt(n) + n / 2) 本题 B = n / sqrt(2q) 会被卡
// 空间复杂度：O(n)

package main

import(
	"slices"
	"math"
	"cmp"
)

func solve() {
	n, q, k := II(), II(), II()
	a := make([]int, n)
	pre := make([]int, n + 1)
	for i := range a {
		a[i] = II()
		pre[i + 1] = pre[i] ^ a[i]
	}

	B := int(math.Ceil(float64(n) / math.Sqrt(float64(q))))

	cnt := map[int]int{}

	// 双开区间莫队
	type query struct {
		id int
		l, r int // (l, r)
		qIdx int
	}
	qs := []query{}

	
	ans := make([]int, q)
	for Q := 0; Q < q; Q++ {
		l, r := II() - 2, II()
		qs = append(qs, query{l / B, l, r, Q})
	}

	// 奇偶排序优化
	slices.SortFunc(qs, func(a, b query) int {
		if a.id != b.id {
			return cmp.Compare(a.id, b.id)
		}
		// 奇数块 r 降序，偶数块 r 升序
		if a.id & 1 == 1 {
			return cmp.Compare(b.r, a.r)
		}
		return cmp.Compare(a.r, b.r)
	})

	l, r, res := -1, 0, 0
	cnt[0] = 1
	for _, b := range qs {		
		// 右端点右移
		for ; r < b.r; r++ {
			res += cnt[pre[r + 1] ^ k]
			cnt[pre[r + 1]]++
		}

		// 左端点左移
		for ; l > b.l; l-- {
			res += cnt[pre[l] ^ k]
			cnt[pre[l]]++
		}

		// 右端点左移
		// 开区间，先左移再删除
		for r > b.r {
			r--
			cnt[pre[r + 1]]--
			res -= cnt[pre[r + 1] ^ k]
		}

		// 左端点右移
		// 开区间，先右移再删除
		for l < b.l {
			l++
			cnt[pre[l]]--
			res -= cnt[pre[l] ^ k]
		}

		ans[b.qIdx] = res
	}

	for _, x := range ans { Println(x) }
}