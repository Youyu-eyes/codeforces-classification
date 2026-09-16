// https://www.luogu.com.cn/problem/P3810
// CDQ 分治
// 三维偏序/陌上花开

package main

import (
	"slices"
	"cmp"
	"sort"
)

type tuple struct {
	a, b, c, cnt, ans int // 新增 ans 字段，随元素一起移动
}

func solve() {
	N, _ := II(), II()
	A := make([]int, N)
	B := make([]int, N)
	C := make([]int, N)
	for i := range N {
		A[i] = II()
		B[i] = II()
		C[i] = II()
	}

	ra, _ := discretize(A)
	rb, _ := discretize(B)
	rc, _ := discretize(C)
	Nums := make([]tuple, N)
	for i := range Nums {
		// 初始化 ans 为 0
		Nums[i] = tuple{ra[i], rb[i], rc[i], 1, 0}
	}

	slices.SortFunc(Nums, func(x, y tuple) int {
		if x.a != y.a {
			return cmp.Compare(x.a, y.a)
		}
		if x.b != y.b {
			return cmp.Compare(x.b, y.b)
		}
		return cmp.Compare(x.c, y.c)
	})

	nums := []tuple{Nums[0]}
	for i := 1; i < N; i++ {
		last := nums[len(nums)-1]
		if last.a == Nums[i].a && last.b == Nums[i].b && last.c == Nums[i].c {
			nums[len(nums)-1].cnt++
		} else {
			nums = append(nums, Nums[i])
		}
	}

	// 提前计算每个元素内部相同项对自身的贡献
	for i := range nums {
		nums[i].ans = nums[i].cnt - 1
	}

	bit := newFenwickTree(N)
	n := len(nums)
	tmp := make([]tuple, n)
	var cdq func(int, int)
	cdq = func(l, r int) {
		if l == r {
			return
		}
		mid := l + (r - l) >> 1
		cdq(l, mid)
		cdq(mid + 1, r)

		j, k := l, l
		for i := mid + 1; i <= r; i++ {
			for ; j <= mid && nums[j].b <= nums[i].b; j++ {
				bit.update(nums[j].c, nums[j].cnt)
				tmp[k] = nums[j]
				k++
			}
			// 将查询结果累加到当前元素的 ans 字段中，不再乘以 cnt
			nums[i].ans += bit.query(0, nums[i].c)
			tmp[k] = nums[i]
			k++
		}

		// 撤销树状数组的修改
		for p := l; p < j; p++ {
			bit.update(nums[p].c, -nums[p].cnt)
		}

		for ; j <= mid; j++ {
			tmp[k] = nums[j]
			k++
		}

		copy(nums[l:r+1], tmp[l:r+1])
	}

	cdq(0, n-1)
	
	// 最后统计时，直接用压缩后的 nums 数组
	ans := make([]int, N)
	for _, x := range nums {
		ans[x.ans] += x.cnt
	}
	
	for i := range ans {
		Println(ans[i])
	}
}

// 返回离散化后的数组以及去重后的元素个数 m
func discretize(arr []int) ([]int, int) {
    unique := slices.Clone(arr)
	slices.Sort(unique)
	unique = slices.Compact(unique)

    m := len(unique)
    rank := make([]int, len(arr))
	for i, x := range arr {
		rank[i] = sort.SearchInts(unique, x)
	}

	return rank, m
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