// https://www.lanqiao.cn/problems/21744/learning/?contest_id=304
// 树状数组求逆序对

// 问题：给你一个数组，排成有序，每次只允许交换相邻两个数字，求最小交换次数
// 结论：最小交换次数为数组中逆序对个数

package main

func solve() {
	n := II()
	odd := make([]int, n / 2)
	even := make([]int, (n + 1) / 2)
	for i := range n {
		x := II() - 1
		if i & 1 != x & 1 {
			Println(-1)
			return
		}
		if i & 1 > 0 {
			odd[i >> 1] = x
		} else {
			even[i >> 1] = x
		}
	}
	Println(inversion(odd) + inversion(even))
}

func inversion(nums []int) (res int) {
	n := len(nums)
	bit := newFenwickTree(2 * n)
	for _, x := range nums {
		res += bit.query(x, 2 * n - 1)
		bit.update(x, 1)
	}
	return
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
	if l > r {
		return 0
	}
	return f.pre(r) - f.pre(l-1)
}
