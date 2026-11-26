// https://atcoder.jp/contests/arc070/tasks/arc070_c
// 洛谷紫 - 省选/NOI-

// slope trick 优化
// 见 leetcode LCP59

package main

import (
    "container/heap"
)

func solve() {
	n := II()
	wood := make([][2]int, n)
	for i := range n {
		wood[i][0] = II()
		wood[i][1] = II()
	}
    lens := make([]int, n)
    for i, p := range wood {
        lens[i] = p[1] - p[0] + 1
    }

    offsetL, offsetR := 0, 0
    gl := &hp{-wood[0][0]}
    gr := &hp{wood[0][0]}
    f := 0

    for i := 1; i < n; i++ {
        offsetL -= lens[i] - 1
        offsetR += lens[i-1] - 1

        w := wood[i][0]
        l := -(*gl)[0] + offsetL

        var x int
        if w < l {
            heap.Push(gl, -(w - offsetL))
            heap.Push(gl, -(w - offsetL))
            x = -heap.Pop(gl).(int) + offsetL
            heap.Push(gr, x-offsetR)
        } else {
            heap.Push(gr, w-offsetR)
            heap.Push(gr, w-offsetR)
            x = heap.Pop(gr).(int) + offsetR
            heap.Push(gl, -(x - offsetL))
        }
        f += abs(w - x)
    }
	Println(f)
}

type hp []int

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i] < h[j] }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(int)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }