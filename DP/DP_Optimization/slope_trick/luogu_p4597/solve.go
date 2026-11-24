// https://www.luogu.com.cn/problem/P4597
// 洛谷蓝 - 提高+/省选-

// slope trick 优化 DP
// 经典模板题，可当结论

package main

import (
	"container/heap"
)

func solve() {
	n := II()
	nums := make([]int, n)
	for i := range n {
		nums[i] = II()
	}
	f := make([]int, n)
	g := &hp{nums[0]}
	for i := 1; i < n; i++ {
		heap.Push(g, nums[i])
		heap.Push(g, nums[i])
		x := heap.Pop(g).(int)
		f[i] = f[i-1] + abs(x - nums[i])
	}
	Println(f[n-1])
}

type hp []int

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i] > h[j] }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(int)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }