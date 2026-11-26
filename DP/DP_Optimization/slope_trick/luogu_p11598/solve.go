// https://www.luogu.com.cn/problem/P11598
// 洛谷紫 - 省选/NOI-

// slope trick
// 

package main

import (
	"container/heap"
)

func solve() {
	n, h := II(), II()
	a := make([]int, n)
	for i := range n {
		a[i] = II()
	}

	gl := &hp{-a[0]}
	gr := &hp{a[0]}
	offsetL, offsetR := 0, 0
	f := 0
	for i := 1; i < n; i++ {
		w := a[i]
		offsetL -= h
		offsetR += h
		l := -(*gl)[0] + offsetL
		var x int
		if w < l {
			heap.Push(gl, -(w-offsetL))
			heap.Push(gl, -(w-offsetL))
			x = -heap.Pop(gl).(int) + offsetL
			heap.Push(gr, x-offsetR)
		} else {
			heap.Push(gr, w-offsetR)
			heap.Push(gr, w-offsetR)
			x = heap.Pop(gr).(int) + offsetR
			heap.Push(gl, -(x-offsetL))
		}
		f += abs(x - w)
	}
	Println(f)
}

type hp []int

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i] < h[j] }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(int)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }