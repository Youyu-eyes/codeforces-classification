// https://www.luogu.com.cn/problem/P2893
// 洛谷青 - 提高

// 数据加强做法，slope trick

// 同 CF13C 将序列翻转后再计算一次

package main

import (
	"slices"
	"container/heap"
)

func solve() {
	n := II()
	a := make([]int, n)
	for i := range n {
		a[i] = II()
	}
	ans := squence(a)
	slices.Reverse(a)
	ans = min(ans, squence(a))
	Println(ans)
}

func squence(a []int) int {
	n := len(a)
	f := 0
	g := &hp{a[0]}
	for i := 1; i < n; i++ {
		heap.Push(g, a[i])
		heap.Push(g, a[i])
		x := heap.Pop(g).(int)
		f += x - a[i]
	}
	return f
}

type hp []int

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i] > h[j] }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(int)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }