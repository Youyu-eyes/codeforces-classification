// https://codeforces.com/problemset/problem/321/E
// 2600
// WQS 二分的做法估计 2800

// 具体分析见 分治优化 DP 题单
// 分治优化枚举 k，时间复杂度为 O(k × n × logn)
// 如果 n 与 k 达到 10^5 就超时了，怎么优化

// 恰好/最多/最少 选择 k 个子数组的最优解，采用 WQS 二分

// WQS 二分的教程见
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-v/solutions/3981282/xiao-bai-si-lu-bi-ji-xi-lie-qian-xi-wqse-nhw1/

// 由于内部转移是一维的，采用二分栈优化 DP

package main

type Matrix [][]int
type pair struct{ b, x int }
type opt struct{ j, l, r int }

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
	n, limit := II(), II()
	u := make(Matrix, n)
	for i := range u {
		u[i] = make([]int, n)
		for j := range u[i] {
			u[i][j] = II()
		}
	}

	pre := Build(u)

	w := func(j, i int) int {
        return pre[i][i] + pre[j][j] - 2 * pre[i][j]
    }

	dp := func(k int) pair {
		f := make([]pair, n + 1)
		for i := range f { f[i].b = inf }
		f[0].b = 0
		q := deque[opt]{}
		q.pushBack(opt{0, 1, n})

		better := func(i, j, x int) bool {
			resI := pair{f[i].b + w(i, x), f[i].x + 1}
			resJ := pair{f[j].b + w(j, x), f[j].x + 1}
			return resI.b < resJ.b || resI.b == resJ.b && resI.x < resJ.x
		}

		firstFind := func(i int, tail opt) int {
			left, right := tail.l - 1, tail.r + 1
			for left + 1 < right {
				mid := left + (right - left) >> 1
				if better(i, tail.j, mid) {
					right = mid
				} else {
					left = mid
				}
			}
			return right
		}

		for i := 1; i <= n; i++ {
			for q.front().r < i {
				q.popFront()
			}

			best := q.front().j
			f[i].b = f[best].b + w(best, i) - k
			f[i].x = f[best].x + 1

			for !q.empty() {
				tail := q.back()
				
				if better(i, tail.j, tail.l) {
					q.popBack()
					continue
				}

				if !better(i, tail.j, tail.r) {
					break
				}
				
				pos := firstFind(i, tail)
				q.popBack()
				q.pushBack(opt{tail.j, tail.l, pos - 1})
				break
			}

			start := 0
			if q.empty() {
				start = i + 1
			} else {
				start = q.back().r + 1
			}
			if start <= n {
				q.pushBack(opt{i, start, n})
			}
		}
		return f[n]
	}

	ans := 0
	var res pair
	left, right := -pre[n][n] - 1, pre[n][n] + 1
	for left + 1 < right {
		mid := left + (right - left) >> 1
		res = dp(mid)
		if res.x <= limit {
			ans = mid * limit + res.b
			left = mid
		} else {
			right = mid
		}
	}

	Println(ans / 2)
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
