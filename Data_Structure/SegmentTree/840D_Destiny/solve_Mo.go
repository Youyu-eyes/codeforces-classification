package main

import (
	"math"
	"cmp"
	"container/heap"
	"slices"
	"sort"
)

const (
	MX = 20
	mask = (1 << MX) - 1
)

func solve_Mo() {
	n, q := II(), II()
	a := make([]int, n)
	for i := range a { a[i] = II() }

	lhq := newLazyHeap()
	cnt := make([]int, n + 1)

	B := int(math.Ceil(float64(n) / math.Sqrt(float64(q))))

	// 双开区间莫队
	type query struct {
		id int
		l, r int // (l, r)
		k int
		qIdx int
	}
	qs := []query{}

	
	ans := make([]int, q)
	for Q := 0; Q < q; Q++ {
		l, r, k := II() - 2, II(), II()
		qs = append(qs, query{l / B, l, r, k, Q})
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

	// 初始化左右端点
	l, r := -1, 0

	add := func(val int) {
		if cnt[val] > 0 {
			lhq.remove(val << MX | cnt[val])
		}
		cnt[val]++
		lhq.push(val << MX | cnt[val])
	}

	remove := func(val int) {
		lhq.remove(val << MX | cnt[val])
		cnt[val]--
		if cnt[val] > 0 {
			lhq.push(val << MX | cnt[val])
		}
	}

	for _, b := range qs {		
		// 右端点右移
		for ; r < b.r; r++ {
			add(a[r])
		}

		// 左端点左移
		for ; l > b.l; l-- {
			add(a[l])
		}

		// 右端点左移
		// 开区间，先左移再删除
		for r > b.r {
			r--
			remove(a[r])
		}

		// 左端点右移
		// 开区间，先右移再删除
		for l < b.l {
			l++
			remove(a[l])
		}
		
		mn := inf
		threshold := (b.r - b.l - 1) / b.k

		tmp := make([]int, 0, b.k)
		for i := 0; i < b.k && lhq.size > 0; i++ {
			top := lhq.pop()
			val := top >> MX
			freq := top & mask
			tmp = append(tmp, top)
			if freq > threshold {
				mn = min(mn, val)
			}
		}

		if mn < inf {
			ans[b.qIdx] = mn
		} else {
			ans[b.qIdx] = -1
		}

		for _, top := range tmp {
			lhq.push(top)
		}
	}

	for _, x := range ans { Println(x) }
}

// ------- Lazy Heap ------- //

func newLazyHeap() *lazyHeap {
	return &lazyHeap{removeCnt: map[int]int{}}
}

// 懒删除堆（默认是最小堆）
type lazyHeap struct {
	sort.IntSlice
	removeCnt map[int]int // 每个元素剩余需要删除的次数
	size      int         // 堆的实际大小
}

// 必须实现的两个接口
func (h *lazyHeap) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *lazyHeap) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

// 加上这行变成最大堆
func (h *lazyHeap) Less(i, j int) bool {
	fi, fj := h.IntSlice[i], h.IntSlice[j]
	if fi != fj {
		return fi > fj
	}
	return h.IntSlice[i] >> MX < h.IntSlice[j] >> MX
}

// 删除
func (h *lazyHeap) remove(v int) {
	h.removeCnt[v]++ // 懒删除
	h.size--
}

// 正式执行删除操作
func (h *lazyHeap) applyRemove() {
	for h.removeCnt[h.IntSlice[0]] > 0 {
		h.removeCnt[h.IntSlice[0]]--
		heap.Pop(h)
	}
}

// 查看堆顶
func (h *lazyHeap) top() int {
	h.applyRemove()
	return h.IntSlice[0] // 真正的堆顶
}

// 出堆
func (h *lazyHeap) pop() int {
	h.applyRemove()
	h.size--
	return heap.Pop(h).(int)
}

// 入堆
func (h *lazyHeap) push(v int) {
	if h.removeCnt[v] > 0 {
		h.removeCnt[v]-- // 抵消之前的删除
	} else {
		heap.Push(h, v)
	}
	h.size++
}