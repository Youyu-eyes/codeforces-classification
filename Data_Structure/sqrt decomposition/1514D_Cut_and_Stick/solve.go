// https://codeforces.com/contest/1514/problem/D
// 2000
// 普通莫队 / 回滚莫队

// 首先本题基于绝对众数的相关结论，一个数组中枚举 2 个数删除，最多操作 min(n / 2, n - m) 其中 n 表示数组长度，m 表示众数个数
// 对于本题，我们可以选择一个众数加一个非众数，然后组成一个序列

// 此时，如果非众数被选完了，众数还有剩余，则我们将剩余的众数每个组成一个子序列，剩余的众数为 m - (n - m) = 2 * m - n 个
// 然后将一个众数添加到刚开始的序列中，这样最初的序列恰好有 ceil(m / 2) 个众数，满足题意
// 共有 2 * m - n - 1 + 1 = 2 * m - n 个子序列

// 如果没有剩余，则说明没有绝对众数，那么子数组就可以作为合法子序列，个数为 1

// 二者取 max 即为答案

// 先讲容易理解的回滚莫队
// 由于我们需要求众数的个数，众数的定义是所有元素出现次数最大的元素值
// 众数的个数就是出现的最大次数
// 由于存在删除操作，max 难以撤销，考虑回滚莫队直接优化掉删除操作

// 维护 cnt 表示元素出现次数，本题保证 a[i] < n 无需离散化
// 右端点右移，左端点左移，统计元素出现次数的同时，维护 cnt 的最大值
// 然后计算答案，再回滚即可

// 时间复杂度：O(n × sqrt(q))，其中 n 表示数组长度，q 表示询问次数
// 空间复杂度：O(n)

// 另有普通莫队做法，见下 solveMo

package main

import (
	"math"
	"slices"
	"cmp"
)


func solve() {
	n, q := II(), II()
	nums := make([]int, n)
	for i := range nums { nums[i] = II() }

	B := int(math.Ceil(float64(n) / math.Sqrt(float64(q))))

	// 双开区间莫队
	type query struct {
		id int
		l, r int // (l, r)
		qIdx int
	}
	qs := []query{}

	cnt := make([]int, n + 1)
	major := 0

	ans := make([]int, q)
	for Q := 0; Q < q; Q++ {
		l, r := II() - 1, II()

		// 大区间离线
		if r - l > B {
			qs = append(qs, query{l / B, l - 1, r, Q})
			continue
		}

		// 小区间暴力
		for i := l; i < r; i++ {
			cnt[nums[i]]++
			major = max(major, cnt[nums[i]])
		}
		ans[Q] = max(1, 2 * major - (r - l))

		for i := l; i < r; i++ {
			cnt[nums[i]]--
		}
		major = 0
	}

	slices.SortFunc(qs, func(a, b query) int {
		return cmp.Or(
			cmp.Compare(a.id, b.id),
			cmp.Compare(a.r, b.r),
		)
	})

	var l, r int
	for i, b := range qs {
		start := (b.id + 1) * B
		if i == 0 || b.id > qs[i - 1].id {
			l = start - 1
			r = start
			clear(cnt)
			major = 0
		}
		
		// 右端点右移
		for ; r < b.r; r++ {
			cnt[nums[r]]++
			major = max(major, cnt[nums[r]])
		}

		// 保留状态
		tmp := major

		// 左端点左移
		for ; l > b.l; l-- {
			cnt[nums[l]]++
			major = max(major, cnt[nums[l]])
		}

		ans[b.qIdx] = max(1, 2 * major - (b.r - b.l - 1))

		// 回滚
		major = tmp
		l = start - 1
		for j := b.l + 1; j <= l; j++ {
			cnt[nums[j]]--
		}
	}

	for _, x := range ans { Println(x) }
}
