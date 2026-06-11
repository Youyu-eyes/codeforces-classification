// https://codeforces.com/problemset/problem/321/E
// 2600
// 决策单调性优化 DP
// 具体分析过程见 分治优化 DP 题单

// 这里写一下二维情况下的二分栈
// 通常情况下，二维使用分治，一维使用二分栈
// 但是二维也可以使用二分栈，但是需要存储上一层的快照

// 实现细节见代码
// 复杂度分析见 分治优化 DP 题单

package main

type Matrix [][]int

func Build(matrix [][]int) Matrix {
    m, n := len(matrix), len(matrix[0])
    sum := make([][]int, m+1)
    sum[0] = make([]int, n+1)
    for i, row := range matrix {
        sum[i+1] = make([]int, n+1)
        for j, x := range row {
            sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j] + x
        }
    }
    return sum
}

func solve() {
    n, K := II(), II()
    u := make(Matrix, n)
    for i := range u {
        u[i] = make([]int, n)
        for j := range u[i] {
            u[i][j] = II()
        }
    }

    pre := Build(u)

    f := make([]int, n + 1)
    for i := range f { f[i] = inf }
    f[0] = 0

    w := func(j, i int) int {
        return pre[i][i] + pre[j][j] - 2 * pre[i][j]
    }

    better := func(i, j, old, x int) bool {
        resJ := old + w(j, x)
        resI := f[i] + w(i, x)
        return resI < resJ
    }

    findFirst := func(i int, tail opt) int {
        left, right := tail.l - 1, tail.r + 1
        for left + 1 < right {
            mid := left + (right - left) >> 1
            if better(i, tail.j, tail.val, mid) {
                right = mid
            } else {
                left = mid
            }
        }
        return right
    }

    for k := 1; k <= K; k++ {
        q := deque[opt]{}
        q.pushBack(opt{k - 1, k, n, f[k - 1]})
        for i := k; i <= n - (K - k); i++ {
            for q.front().r < i {
                q.popFront()
            }
            best := q.front()

            for !q.empty() {
                tail := q.back()
                if better(i, tail.j, tail.val, tail.l) {
                    q.popBack()
                    continue
                }

                if !better(i, tail.j, tail.val, tail.r) {
                    break
                }

                pos := findFirst(i, tail)

                q.popBack()
                q.pushBack(opt{tail.j, tail.l, pos - 1, tail.val})
                break
            }

			start := 0
            if q.empty() {
                start = i + 1
            } else {
				start = q.back().r + 1
			}
			if start <= n {
				q.pushBack(opt{i, start, n, f[i]})
			}

            // 注意这里，先用老的 f[i] 更新队列
            // 再更新 f[i]
            f[i] = best.val + w(best.j, i)
        }
    }

    Println(f[n] / 2)
}

type opt struct{
    j int
    l int
    r int
	val int
}

// 泛型双端队列，T 可以是任意类型
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
