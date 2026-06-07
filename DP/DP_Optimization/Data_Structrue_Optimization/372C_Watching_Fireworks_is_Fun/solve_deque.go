// 方法二：单调队列优化
// 由于区间左端点和右端点持续单调右移，是典型的滑动窗口最大值问题
// 适用单调队列
// 由于是双闭区间，可以采取先更新窗口在计算 nf[p] 的方式

// 时间复杂度：O(n·m)
// 空间复杂度：O(n)

package main

import(
	"fmt"
	"bufio"
	"slices"
)

func solve2(in *bufio.Reader, out *bufio.Writer) {
	var n, m, d int
	fmt.Fscan(in, &n, &m, &d)

	type Firework struct{ a, b, t int }
	fw := make([]Firework, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &fw[i].a, &fw[i].b, &fw[i].t)
	}
    
    f := make([]int, n + 1)
    for p := 1; p <= n; p++ {
        f[p] = fw[0].b - abs(fw[0].a - p)
    }
    
    for i := 1; i < m; i++ {
        dist := (fw[i].t - fw[i - 1].t) * d
        q := deque[int]{}
        nf := make([]int, n + 1)

        for j := 1; j <= min(n, 1 + dist); j++ {
            for !q.empty() && f[q.back()] <= f[j] {
                q.popBack()
            }
            q.pushBack(j)
        }
        
        for p := 1; p <= n; p++ {
            left := max(1, p - dist)
            right := min(n, p + dist)

            if p > 1 {
                if right <= n {
                    for !q.empty() && f[q.back()] <= f[right] {
                        q.popBack()
                    }
                    q.pushBack(right)
                }
            }

            for !q.empty() && q.front() < left {
                q.popFront()
            }

            best := q.front()
            nf[p] = f[best] + fw[i].b - abs(fw[i].a - p)
        }
        copy(f, nf)
    }
    fmt.Fprintln(out, slices.Max(f[1:]))
}

// 双端队列
type deque[T any] struct{ l, r []T }

func (q deque[T]) empty() bool {
	return len(q.l) == 0 && len(q.r) == 0
}

func (q deque[T]) size() int {
	return len(q.l) + len(q.r)
}

func (q *deque[T]) pushFront(v T) {
	q.l = append(q.l, v)
}

func (q *deque[T]) pushBack(v T) {
	q.r = append(q.r, v)
}

func (q *deque[T]) popFront() (v T) {
	if len(q.l) > 0 {
		q.l, v = q.l[:len(q.l)-1], q.l[len(q.l)-1]
	} else {
		v, q.r = q.r[0], q.r[1:]
	}
	return
}

func (q *deque[T]) popBack() (v T) {
	if len(q.r) > 0 {
		q.r, v = q.r[:len(q.r)-1], q.r[len(q.r)-1]
	} else {
		v, q.l = q.l[0], q.l[1:]
	}
	return
}

func (q deque[T]) front() T {
	if len(q.l) > 0 {
		return q.l[len(q.l)-1]
	}
	return q.r[0]
}

func (q deque[T]) back() T {
	if len(q.r) > 0 {
		return q.r[len(q.r)-1]
	}
	return q.l[0]
}

// 0 <= i < q.size()
func (q deque[T]) get(i int) T {
	if i < len(q.l) {
		return q.l[len(q.l)-1-i]
	}
	return q.r[i-len(q.l)]
}
