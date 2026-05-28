// https://codeforces.com/contest/455/problem/D
// 2700
// 分块难题

// 码量极大，涉及到一些代码方面的小 trick

// 更新操作，可以等价于一次单点更新和一次区间更新
// 查询操作，需要维护区间数字的出现次数
// 考虑到线段树难以维护区间数字出现次数，考虑分块

// 将数组 a 分块后，先看更新操作，假设 b 为当前的块
// 更新：
// 先看完整块，我们需要在 b.l 插入上一个块的最后一个数，然后将 b.r 保存并删除，用于更新下一个块
// 再看左端点，假设至少经过一个块，左端点的块需要将 l 处插入右端点的数，然后将 b.r 的数保存并删除，用于更新下一个块
// 再看右端点，右端点需要在 b.l 处插入上一个块的最后一个数，r - 1 处的数删除
// 特殊情况，左右端点都在一个块，暴力更新即可
// 在插入和删除的时候，维护每个数字的出现次数

// 由更新操作可知：
// 需要一个数据结构维护数字出现次数，定义为 cnt
// 由于本题的 1 < a[i] < n，我们可以直接用 n 长度数组，否则用 哈希表 或 离散化 + 数组

// 需要一个变量维护更新时，上一个块的最后一个数，定义为 cur

// 需要一个数据结构维护块内的数字
// 对于分块的题目，完整块至多有 n / B 个，实际上这个是非常大的，我们需要 O(1) 时间处理完整块，这也是分块的意义
// 有两种方法

// 1. 数组 + 指针
// 我们考虑用 b.head 维护左端点
// 完整块：左移 b.head，此时的 head 指向需要删除的数，用 tmp 保存，然后将 b.arr[b.head] 赋值为 cur，最后用 tmp 更新 cur
// 左端点：你可以用循环数组的思想，用指针暴力进行删除和插入操作，但是时间复杂度还是 O(B)
// 更直接也更好写的思路是，直接重建 b.arr，将 b.head 归零，这样整个数组没有偏移，直接用 l - b.l 写，非常清楚
// 右端点同理
// 在插入和删除时，不要忘记维护 cnt

// 时间复杂度：完整块最多 n / B 块，更新需要 O(1) 时间，左边最多一个不完整块，需要 O(B) 时间，右边同理
// 总体时间为 n / B + B

// 查询：
// 完整块，直接查询 cnt[k]，并累加
// 不完整块：同理，可以用指针慢慢算，但是更清楚的写法是直接重建，然后跑一个简单的循环即可
// 完整块最多 n / B，不完整块最多 B
// 复杂度仍然是 n / B + B

// 总体复杂度为 O(n / B + B) 根据均值不等式，B = sqrt(n) 时取最小值

// 时间复杂度：O(q×sqrt(n))
// 空间复杂度：O(n)

// 总体来说，数组的写法还是比较麻烦的，双端队列的写法见 solve2，会更简单

package main

import (
	"bufio"
	"fmt"
	"os"
	"math"
	"strings"
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

type block struct {
	l, r int
	cnt  []int
	head int
	arr  []int
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func (b *block) rebuild() {
	if b.head == 0{
		return
	}
	reverse(b.arr[:b.head])
	reverse(b.arr[b.head:])
	reverse(b.arr)
	b.head = 0
}

func (b *block) move(idxL, idxR int) {
	for j := idxR - 1; j > idxL; j-- {
		b.arr[j] = b.arr[j - 1]
	}
}

func solve() {
	n := II()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = II()
	}

	B := 1000 // isqrt(n) 是理论最优解

	blocks := make([]block, (n - 1) / B + 1)
	calcBlock := func(l, r int) []int {
		cnt := make([]int, n + 1) 
		for j := l; j < r; j++ {
			v := a[j]
			cnt[v]++
		}
		return cnt
	}
	for i := 0; i < n; i += B {
		r := min(i+B, n)
		cnt := calcBlock(i, r)
		blocks[i/B] = block{i, r, cnt, 0, a[i: r]}
	}

	// [l, r), 从 0 开始
	sqrtUpdate := func(l, r int) {
		last := &blocks[(r - 1) / B]
		cur := last.arr[(last.head + (r - 1 - last.l)) % (last.r - last.l)]

		if l / B == (r - 1) / B {
			last.rebuild()
			last.move(l - last.l, r - last.l)
			last.arr[l - last.l] = cur
			return
		}

		for i := range blocks {
			b := &blocks[i]
			L := b.r - b.l
			if b.l >= r {
				break
			}
			if b.r <= l {
				continue
			}
			if l <= b.l && b.r <= r {
				// 完整块
				b.head = (b.head - 1 + L) % L
				tmp := b.arr[b.head]
				b.cnt[tmp]--
				b.arr[b.head] = cur
				b.cnt[cur]++
				cur = tmp
			} else if b.l < l {
				b.rebuild()
				idxL := l - b.l
				tmp := b.arr[L - 1]
				b.cnt[tmp]--
				b.move(idxL, L)
				b.arr[idxL] = cur
				b.cnt[cur]++
				cur = tmp
			} else {
				b.rebuild()
				idxR := r - b.l
				b.cnt[b.arr[idxR - 1]]--
				b.move(0, idxR)
				b.arr[0] = cur
				b.cnt[cur]++
				return
			}
		}
	}

	// [l, r)，0-base
	sqrtQuery := func(l, r int, v int) (sum int) {
		for i := range blocks {
			b := &blocks[i]
			if b.l >= r {
				break
			}
			if b.r <= l {
				continue
			}
			if l <= b.l && b.r <= r {
				// 完整块
				sum += b.cnt[v]
			} else {
				// 部分块
				b.rebuild()
				idxL := max(b.l, l) - b.l
				idxR := min(b.r, r) - b.l
				for j := idxL; j < idxR; j++ {
					if b.arr[j] == v {
						sum++
					}
				}
			}
		}
		return
	}

	lastans := 0
	q := II()
	for Q := 0; Q < q; Q++ {
		queryNumber := II()
		if queryNumber == 1 {
			lPrime, rPrime := II(), II()

			l := ((lPrime + lastans - 1) % n) + 1
			r := ((rPrime + lastans - 1) % n) + 1

			if l > r {
				l, r = r, l
			}

			sqrtUpdate(l-1, r)
		} else {
			lPrime, rPrime, kPrime := II(), II(), II()

			l := ((lPrime + lastans - 1) % n) + 1
			r := ((rPrime + lastans - 1) % n) + 1
			if l > r {
				l, r = r, l
			}
			k := ((kPrime + lastans - 1) % n) + 1

			ans := sqrtQuery(l-1, r, k)
			Println(ans)

			lastans = ans
		}
	}
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

// 数组要手动切片 sum(nums[:])
func sum[T int | int64 | float64](nums []T) T {
	var total T
	for _, v := range nums {
		total += v
	}
	return total
}

func isqrt(x int) int {
	rt := int(math.Sqrt(float64(x))) // 可能会算多一点点
	if rt*rt > x {
		rt--
	}
	return rt
}

func ReverseString(s string) string {
    var b strings.Builder
    b.Grow(len(s))
    for i := len(s) - 1; i >= 0; i-- {
        b.WriteByte(s[i])
    }
    return b.String()
}

// 2. 双端队列
// 对于本题而言，左端点插入一个数，右端点删除一个数，可适用双端队列
// 完整的块，左端入队，右端出队，都是 O(1) 时间
// 左端点，在 b.q 的 l - b.l 插入一个数，最多 O(B) 时间
// 右端点，在 b.q 的 r - b.l 删除一个数，最多 O(B) 时间

// 总体时间和空间复杂度与数组写法相同

func solve2() {
	n := II()
	a := make([]int, n)
	for i := range a { a[i] = II() }

	B := 1000

	type block struct {
		l, r int // [l, r)
		cnt  []int
		q    deque[int]
	}

	blocks := make([]block, (n-1)/B+1)
	calcBlock := func(l, r int) ([]int, deque[int]) {
		cnt := make([]int, n + 1)
		q := deque[int]{}
		for j := l; j < r; j++ {
			v := a[j]
			cnt[v]++
			q.pushBack(v)
		}
		return cnt, q
	}
	for i := 0; i < n; i += B {
		r := min(i+B, n)
		cnt, q := calcBlock(i, r)
		blocks[i/B] = block{i, r, cnt, q}
	}

	// [l, r), 从 0 开始
	sqrtUpdate := func(l, r int) {
		last := &blocks[(r - 1) / B]
		cur := last.q.get(r - 1 - last.l)

		if l / B == (r - 1) / B {
			last.q.erase(r - 1 - last.l)
			last.q.insert(cur, l - last.l)
			return
		}

		for i := range blocks {
			b := &blocks[i]
			if b.l >= r {
				break
			}
			if b.r <= l {
				continue
			}
			if l <= b.l && b.r <= r {
				// 完整块
				tmp := b.q.back()
				b.q.popBack()
				b.cnt[tmp]--
				b.q.pushFront(cur)
				b.cnt[cur]++
				cur = tmp
			} else if b.l < l {
				tmp := b.q.back()
				b.q.popBack()
				b.cnt[tmp]--
				b.q.insert(cur, l - b.l)
				b.cnt[cur]++
				cur = tmp
			} else {
				tmp := b.q.get(r - 1 - b.l)
				b.q.erase(r - 1 - b.l)
				b.cnt[tmp]--
				b.q.pushFront(cur)
				b.cnt[cur]++
				return
			}
		}
	}

	// [l, r), 从 0 开始
	sqrtQuery := func(l, r int, v int) (sum int) {
		for i := range blocks {
			b := &blocks[i]
			if b.l >= r {
				break
			}
			if b.r <= l {
				continue
			}
			if l <= b.l && b.r <= r {
				// 完整块
				sum += b.cnt[v]
			} else {
				// 部分块
				bl := max(b.l, l)
				br := min(b.r, r)
				for j := bl; j < br; j++ {
					if b.q.get(j - b.l) == v {
						sum++
					}
				}
			}
		}
		return
	}

	lastans := 0
	q := II()
	for Q := 0; Q < q; Q++ {
		queryNumber := II()
		if queryNumber == 1 {
			lPrime, rPrime := II(), II()

			l := ((lPrime + lastans - 1) % n) + 1
			r := ((rPrime + lastans - 1) % n) + 1

			if l > r {
				l, r = r, l
			}

			sqrtUpdate(l-1, r)
		} else {
			lPrime, rPrime, kPrime := II(), II(), II()

			l := ((lPrime + lastans - 1) % n) + 1
			r := ((rPrime + lastans - 1) % n) + 1
			if l > r {
				l, r = r, l
			}
			k := ((kPrime + lastans - 1) % n) + 1

			ans := sqrtQuery(l-1, r, k)
			Println(ans)

			lastans = ans
		}
	}
}

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

// insert 在逻辑索引 idx 处插入 val (0 <= idx <= q.size())
func (q *deque[T]) insert(val T, idx int) {
	if idx == 0 {
		q.pushFront(val)
		return
	}
	if idx == q.size() {
		q.pushBack(val)
		return
	}

	if idx <= len(q.l) {
		pos := len(q.l) - idx
		var zero T
		q.l = append(q.l, zero)
		copy(q.l[pos+1:], q.l[pos:len(q.l)-1])
		q.l[pos] = val
	} else {
		ridx := idx - len(q.l)
		if ridx == 0 {
			q.r = append([]T{val}, q.r...)
		} else {
			var zero T
			q.r = append(q.r, zero)
			copy(q.r[ridx+1:], q.r[ridx:len(q.r)-1])
			q.r[ridx] = val
		}
	}
}

// erase 删除逻辑索引 idx 处的元素 (0 <= idx < q.size())
func (q *deque[T]) erase(idx int) {
	if idx < len(q.l) {
		pos := len(q.l) - 1 - idx
		if pos == len(q.l)-1 {
			q.l = q.l[:len(q.l)-1]
		} else {
			copy(q.l[pos:], q.l[pos+1:])
			q.l = q.l[:len(q.l)-1]
		}
	} else {
		ridx := idx - len(q.l)
		if ridx == 0 {
			q.r = q.r[1:]
		} else {
			copy(q.r[ridx:], q.r[ridx+1:])
			q.r = q.r[:len(q.r)-1]
		}
	}
}
