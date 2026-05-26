// https://codeforces.com/contest/372/problem/C
// 2100
// RMQ 优化 DP
// 乍一看很难看出转移，本题需要把状态定义出来之后才能很好的看到
// 定义 f[i][p] 为放第 i 次烟花，人在 p 位置的最大开心度
// 思考上一次放烟花，假设在 q 位置，那么 p 与 q 满足：
// 如果 q 在 p 左边：
// q = p - (t[i] - t[i - 1]) * d
// 如果 q 在 p 右边
// q = p + (t[i] - t[i - 1]) * d
// 左边与 1 取最大，右边与 n 取最小，保证下标不越界，则我们可以得到 q 的取值范围
// q in [max(1, p - (t[i] - t[i - 1]) * d), min(n, p + (t[i] - t[i - 1]) * d)]

// 状态转移方程
// f[i][p] = max(f[i - 1][q]) + b[i] - abs(a[i] - p)
// 滚动数组空间优化
// nf[p] = max(f[q]) + b[i] - abs(a[i] - p)
// q 的取值是一段连续区间，可以采用 ST 表优化 RMQ 的过程
// 在更新 nf 之前，对 f 数组建 ST 表

// 普通 ST 表也能过，但是 ST_fast 更快
// 本题采用普通 ST 表
// 普通 ST 表见数据结构模板，ST_fast 见 Go 语言算法模板库

// 时间复杂度 O(n·m·logn)
// 空间复杂度 O(n)

// 另有双端队列做法，见下

package main

import (
	"bufio"
	"fmt"
	"os"
	"math/bits"
	"slices"
)

const (
	MOD = 1_000_000_007
	inf = 0x3f3f3f3f
)

var (
	in  *bufio.Reader
	out *bufio.Writer
)

// n := II()
func II() int {
    n := 0
    neg := false
    b, _ := in.ReadByte()

    for b < '0' || b > '9' {
        if b == '-' {
            neg = true
            b, _ = in.ReadByte()
            break
        }
        b, _ = in.ReadByte()
    }

    for b >= '0' && b <= '9' {
        n = n*10 + int(b&15)
        b, _ = in.ReadByte()
    }
    if neg {
        n = -n
    }
    return n
}

// n := IF()
func IF() (x float64) {
	fmt.Fscan(in, &x)
	return
}

// s := IS()
func IS() (x string) {
	fmt.Fscan(in, &x)
	return
}

func Print(a ...any)            { fmt.Fprint(out, a...) }
func Println(a ...any)          { fmt.Fprintln(out, a...) }
func Printf(f string, a ...any) { fmt.Fprintf(out, f, a...) }


func solve() {
	n := II()
	m := II()
	d := II()

	type Firework struct{ a, b, t int }
	fw := make([]Firework, m)
	for i := 0; i < m; i++ {
		fw[i].a = II()
		fw[i].b = II()
		fw[i].t = II()
	}

	f := make([]int, n + 1)
	for p := 1; p <= n; p++ {
		f[p] = fw[0].b - abs(fw[0].a - p)
	}

	for i := 1; i < m; i++ {
		st := newSparseTable(f, func(a, b int) int { return max(a, b) })
		dist := (fw[i].t - fw[i - 1].t) * d
		nf := make([]int, n+1)
		for p := 1; p <= n; p++ {
			best := st.query(max(1, p - dist), min(n, p + dist) + 1)
			nf[p] = best + fw[i].b - abs(fw[i].a - p)
		}
		copy(f, nf)
	}

	Println(slices.Max(f[1:]))
}


func main() {
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	T := 1

	for t := 0; t < T; t++ {
		solve()
	}
}


func abs[T int | int64 | float64](x T) T {
	if x >= 0 { return x }
	return -x
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// ------- 数值 ST 表 ------- //
type SparseTable[T any] struct {
	st [][]T
	op func(T, T) T
}

// op 需满足结合律（如 min, max, gcd 等）
func newSparseTable[T any](nums []T, op func(T, T) T) *SparseTable[T] {
	n := len(nums)
	w := bits.Len(uint(n))
	st := make([][]T, w)
	for i := range st {
		st[i] = make([]T, n)
	}
	copy(st[0], nums)
	for i := 1; i < w; i++ {
		step := 1 << (i - 1)
		for j := 0; j+step < n && j+(1<<i) <= n; j++ {
			st[i][j] = op(st[i-1][j], st[i-1][j+step])
		}
	}
	return &SparseTable[T]{st: st, op: op}
}

func (st *SparseTable[T]) query(l, r int) T {
	k := bits.Len(uint(r-l)) - 1
	return st.op(st.st[k][l], st.st[k][r-1<<k])
}


// 方法二：单调队列优化
// 由于区间左端点和右端点持续单调右移，是典型的滑动窗口最大值问题
// 适用单调队列
// 由于是双闭区间，可以采取先更新窗口在计算 nf[p] 的方式

// 时间复杂度：O(n·m)
// 空间复杂度：O(n)

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
