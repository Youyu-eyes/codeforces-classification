// 普通莫队
// 思路是维护一个 freq 频率数组，表示 cnt[x] 出现的次数，用于删除操作中维护 major

// 这里重要的一个结论是，freq[cnt[x]]-- 后，如果 freq[major] == 0，则 cnt[x] 一定为 major
// 因为如果 cnt[x] != major，则 freq[major] 在删除操作中没有被更改，但是之前的操作一定保证 freq[major] > 0，这与 freq[major] = 0 矛盾
// 假设不成立，因此 cnt[x] = major

// 然后由于 freq[cnt[x]--]++ = freq[major--]++，因此 freq[major--] > 0，所以我们不需要枚举全部的 cnt[x] 来更新 major，major 一定是 cnt[x]--
// 这样的话，删除的时间复杂度就是 O(1) 的

// 总体时空复杂度与回滚莫队一致

package main

import(
	"math"
	"slices"
	"cmp"
)

func solveMo() {
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

	l, r, res := B - 1, B, 0
	cnt := make([]int, n + 1)
	freq := make([]int, n + 1)
	major := 0

	add := func(x int) {
		freq[cnt[x]]--

		cnt[x]++
		freq[cnt[x]]++

		if major < cnt[x] {
			major = cnt[x]
		}
	}

	remove := func(x int) {
		freq[cnt[x]]--

		cnt[x]--
		freq[cnt[x]]++

		if freq[major] == 0 {
			major--
		}
	}

	for _, b := range qs {		
		// 右端点右移
		for ; r < b.r; r++ {
			add(nums[r])
		}

		// 左端点左移
		for ; l > b.l; l-- {
			add(nums[l])
		}

		// 右端点左移
		// 开区间，先左移再删除
		for r > b.r {
			r--
			remove(nums[r])
		}

		// 左端点右移
		// 开区间，先右移再删除
		for l < b.l {
			l++
			remove(nums[l])
		}

		res = max(1, 2 * major - (b.r - b.l - 1))

		ans[b.qIdx] = res
	}

	for _, x := range ans {
		Println(x)
	}
}