// https://codeforces.com/contest/86/problem/D
// 2200
// 普通莫队模板题

// 这个题是莫队模板题，读者可以当作母题

// 本题没有更新，只有区间查询，而线段树难以维护数字频率
// 分块主要针对的是区间更新操作的优化，这样的查询仍然是暴力
// 因此采用莫队

// 莫队的主要思想就是滑动窗口，通过将查询排序，使得我们可以用增量的思想计算多余的部分或减少的部分
// 将数组分块，将左端点按照分块的区间分组，在同一个块内的数据一起处理
// 左端点在同一个块内的询问，将询问的右端点按照奇偶分离排序，编号为奇数的倒序，偶数的正序
// 这样的好处是，不单单同一个块内我们有重复数据可利用，对于不同块之间，我们也可以通过重复的数据加速计算

// 但是请注意，在代码实现的时候，我们不能将右端点的左右移动情况用 if-else 简单判断
// 假设上一块中，查询的最大右端点为 6，假设上一块是偶数块
// 这一块为奇数块，右端点左移，假设最大右端点为 8
// 此时我们发现，从 6 - 8 还需要向右移动两次，才能将本块的查询初始化

// 因此，上一次是偶数，这一次是奇数的情况下，第一次右端点可能向右移动，然后向左
// 同理，上一次是奇数，这一次是偶数的情况下，第一次右端点可能向左移动，然后向右

// 我们不能简单地分开，反而用循环条件进行判断更简洁也最正确

// 剩下的无非就是增量的计算
// (k + 1) * (k + 1) * x - k * k * x
// = (2 * k + 1) * x
// 区间扩张时加，缩小时减

// 时间复杂度：O(n × sqrt(q))
// 空间复杂度：O(n)

package main

import (
	"bufio"
	"cmp"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

const (
	MOD = 1_000_000_007
	inf = 0x3f3f3f3f
	MX = 1000005
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
	n, q := II(), II()
	nums := make([]int, n)
	for i := range nums { nums[i] = II() }

	B := int(math.Ceil(float64(n) / math.Sqrt(float64(q))))
	cnt := [MX]int{}

	// 双开区间莫队
	type query struct {
		id int
		l, r int // (l, r)
		qIdx int
	}
	qs := []query{}

	var res int
	ans := make([]int, q)
	for Q := 0; Q < q; Q++ {
		l, r := II() - 1, II()

		qs = append(qs, query{l / B, l - 1, r, Q})
	}

	slices.SortFunc(qs, func(a, b query) int {
		if a.id != b.id {
			return cmp.Compare(a.id, b.id)
		}
		// 奇数块 r 降序，偶数块 r 升序
		if a.id & 1 == 1 {
			return cmp.Compare(b.r, a.r)
		}
		return cmp.Compare(a.r, b.r)
	})

	l, r := -1, 0
	for _, b := range qs {		
		// 偶数块，右端点右移
		for ; r < b.r; r++ {
			x := nums[r]
			res += 2 * cnt[x] * x + x
			cnt[x]++
		}

		// 左端点左移
		for ; l > b.l; l-- {
			x := nums[l]
			res += 2 * cnt[x] * x + x
			cnt[x]++
		}
		
		// 奇数块，右端点左移
		// 开区间，先左移再删除
		for r > b.r {
			r--
			x := nums[r]
			res -= 2 * cnt[x] * x - x
			cnt[x]--
		}

		// 左端点右移
		// 开区间，先右移再删除
		for l < b.l {
			l++
			x := nums[l]
			res -= 2 * cnt[x] * x - x
			cnt[x]--
		}

		ans[b.qIdx] = res
	}

	for _, x := range ans {
		Println(x)
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