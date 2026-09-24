// https://www.lanqiao.cn/problems/21746/learning/?contest_id=304
// 详细做法见栈部分的题解

// 题解链接
// https://github.com/Youyu-eyes/codeforces-classification/blob/main/Data_Structure/strak/lanqiao_26xspws_2_5/solution.md

package main

import (
	"slices"
	"cmp"
)

func solve() {
	n := II()
	nums := make([]int, n)
	for i := range n {
		nums[i] = II()
	}

	order := make([]int, n)
	for i := range order { order[i] = i }
	slices.SortFunc(order, func(a, b int) int { return cmp.Compare(nums[b], nums[a]) })

	bit := newFenwickTree(n)
	ans := n * (n + 1) / 2
	
	i := 0
	for i < n {
		start := i
		for i < n && nums[order[i]] == nums[order[start]] {
			bit.update(order[i], 1)
			i++
		}
		for t := start; t < i; t++ {
			j := order[t]
			le := bit.pre(j)
			if bit.query(j + 1, n - 1) >= 2 {
				pos := bit.kth(le + 2)
				ans -= n - pos
			}
		}
	}

	Println(ans)
}

type fenwick []int

func newFenwickTree(n int) fenwick {
	return make(fenwick, n+1) // 使用下标 1 到 n
}

// a[i] 增加 val
// 时间复杂度 O(log n)
func (f fenwick) update(i, val int) {
	for i++; i < len(f); i += i & -i {
		f[i] += val
	}
}

// 求前缀和 a[1] + ... + a[i]
// 时间复杂度 O(log n)
func (f fenwick) pre(i int) (res int) {
	for i++; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}

// 求区间和 a[l] + ... + a[r]
// 时间复杂度 O(log n)
// 0 <= l <= r < n
func (f fenwick) query(l, r int) int {
	return f.pre(r) - f.pre(l-1)
}

// 返回最小的 0-based 下标 i，满足 a[0] + ... + a[i] >= k
// 若总和 < k，返回 n
func (f fenwick) kth(k int) int {
	n := len(f) - 1
	bit := 1
	for bit<<1 <= n {
		bit <<= 1
	}
	idx := 0
	for ; bit > 0; bit >>= 1 {
		nxt := idx + bit
		if nxt <= n && f[nxt] < k {
			k -= f[nxt]
			idx = nxt
		}
	}
	return idx
}
