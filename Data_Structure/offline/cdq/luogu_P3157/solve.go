// https://www.luogu.com.cn/problem/P3157
// 动态逆序对
// cdq 分治
// 每个元素具有 (i, a[i], t[i]) 三个属性
// 分治后
// 对于右侧的每个数，左侧的数需要满足
// a[j] > a[i]
// t[j] > t[i]

// 对于左侧的每个数，右侧的数需要满足
// a[j] > a[i]
// t[j] < t[i]

// 注意枚举左侧时，由于我们采用 a[i] 降序排序
// 当 a[j] 减少，i 越大越满足，因此需要倒序枚举 j

// 归并排序时，由于 tmp 是复制了 nums 数组的
// 因此跑两个双指针时，必须在后一个双指针中进行归并

package main

import (
	"slices"
)

type tuple struct {
	v, t, inv int
}

func solve() {
	n, m := II(), II()
	bit := newFenwickTree(n)
	ans := 0
	nums := make([]tuple, n)
	mp := make([]int, n)
	for i := range n {
		nums[i].v = II() - 1
		ans += bit.query(nums[i].v + 1, n - 1)
		bit.update(nums[i].v, 1)
		mp[nums[i].v] = i
	}

	for t := range m {
		i := mp[II() - 1]
		nums[i].t = t + 1
	}
	for i := range n {
		if nums[i].t == 0 {
			nums[i].t = m
		} else {
			nums[i].t--
		}
	}

	bit = newFenwickTree(m + 1)
	var cdq func(int, int)
	tmp := make([]tuple, n)
	cdq = func(l, r int) {
		if l == r {
			return
		}

		mid := (l + r) >> 1
		cdq(l, mid)
		cdq(mid + 1, r)
		
		i := r
		for j := mid; j >= l; j-- {
			for ; i > mid && nums[j].v > nums[i].v; i-- {
				bit.update(nums[i].t, 1)
			}
			nums[j].inv += bit.query(nums[j].t + 1, m)
		}

		for p := r; p > i; p-- {
			bit.update(nums[p].t, -1)
		}

		j, k := l, l
		for i := mid + 1; i <= r; i++ {
			for ; j <= mid && nums[j].v > nums[i].v; j++ {
				bit.update(nums[j].t, 1)
				tmp[k] = nums[j]
				k++
			}
			nums[i].inv += bit.query(nums[i].t + 1, m)
			tmp[k] = nums[i]
			k++
		}

		for p := l; p < j; p++ {
			bit.update(nums[p].t, -1)
		}

		for ; j <= mid; j++ {
			tmp[k] = nums[j]
			k++
		}

		copy(nums[l:r+1], tmp[l:r+1])
	}
	cdq(0, n - 1)
	
	slices.SortFunc(nums, func(a, b tuple) int { return a.t - b.t })
	for i := range m {
		Println(ans)
		ans -= nums[i].inv
	}
}

type fenwick []int

func newFenwickTree(n int) fenwick {
	return make(fenwick, n+1)
}

func (f fenwick) update(i, val int) {
	for i++; i < len(f); i += i & -i {
		f[i] += val
	}
}

func (f fenwick) pre(i int) (res int) {
	for i++; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}

func (f fenwick) query(l, r int) int {
	if l > r {
		return 0
	}
	return f.pre(r) - f.pre(l-1)
}
