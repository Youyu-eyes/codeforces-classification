// https://codeforces.com/problemset/problem/713/C
// 2300

// slope trick
// DP 优化

package main

import (
	"container/heap"
)

func solve() {
	n := II()
	nums := make([]int, n)
	for i := range n {
		nums[i] = II() - i
	}
	f := make([]int, n)
	g := &hp{nums[0]}
	for i := 1; i < n; i++ {
		a := nums[i]
		heap.Push(g, a)
		heap.Push(g, a)
		x := heap.Pop(g).(int)
		f[i] = f[i-1] + x - a
	}
	Println(f[n-1])
}

type hp []int

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i] > h[j] }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(int)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }