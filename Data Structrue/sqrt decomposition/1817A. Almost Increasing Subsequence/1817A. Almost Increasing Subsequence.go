// https://codeforces.com/contest/1817/problem/A
// 1500
// 贪心 + 莫队

// 本题最重要的是想清楚贪心的思路
// 将数组按照非递增分段
// 例如 4 5 6 7 2 2 1 3 5
// 得到 4 | 5 | 6 | 7 2 2 1 | 3 | 5
// 分段后，如果一段的长度小于 3，则块内的数全部能选
// 如果段长度 > 2，则这一段只能选两个，并且一定能选头尾两个数，具体证明见官解

// 莫队算法
// 核心思路依旧是贪心，右端点右移，左端点抖动
// 分块后将左端点分组，以右端点排序，则每次右端点从下一个组的左端出发
// 小区间暴力，大区间离线
// 注意贡献，从右往左看，a[l] >= a[l + 1] >= a[l + 2] 表示 a[l + 1] 没有为答案做出贡献，a[l] 做出了贡献，二者抵消
// 因此可以等价于 a[l] 没有做出贡献

// 从左往右看同理，计算贡献

// 时间复杂度：O(nsqrt(q))
// 空间复杂度：O(n)

// 另有贪心 + 前缀和解法，见贪心题单

package main

import (
	"bufio"
	"fmt"
	"os"
	"math"
	"strings"
	"slices"
	"cmp"
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


func solveMo() {
	n, q := II(), II()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = II()
	}

	B := int(math.Ceil(float64(n) / math.Sqrt(float64(q))))
	B = max(B, 7)

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

		// 大区间离线
		if r - l > B {
			qs = append(qs, query{l / B, l - 1, r, Q})
			continue
		}

		// 小区间暴力
		if r - l < 3 {
			ans[Q] = r - l
		} else {
			res = 2
			for i := l + 2; i < r; i++ {
				if !(a[i] <= a[i - 1] && a[i - 1] <= a[i - 2]) {
					res++
				}
			}
			ans[Q] = res
		}
	}

	slices.SortFunc(qs, func(a, b query) int {
		return cmp.Or(
			cmp.Compare(a.id, b.id),
			cmp.Compare(a.r, b.r),
		)
	})

	var l, r int
	for i, b := range qs {
		start := (b.id + 1) * B
		if i == 0 || b.id > qs[i - 1].id {
			l = start - 1
			r = start
			res = 0
		}

		for ; r < b.r; r++ {
			if l + 3 > r || !(a[r] <= a[r - 1] && a[r - 1] <= a[r - 2]) {
				res++
			}
		}

		for ; l > b.l; l-- {
			if l + 3 > r || !(a[l] >= a[l + 1] && a[l + 1] >= a[l + 2]) {
				res++
			}
		}

		for l < b.l {
			l++
			if l + 3 > r || !(a[l] >= a[l + 1] && a[l + 1] >= a[l + 2]) {
				res--
			}
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
		solveMo()
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
