// https://codeforces.com/contest/1817/problem/A
// 1500
// 贪心 + 前缀和

// 本题最重要的是想清楚贪心的思路
// 将数组按照非递增分段
// 例如 4 5 6 7 2 2 1 3 5
// 得到 4 | 5 | 6 | 7 2 2 1 | 3 | 5
// 分段后，如果一段的长度小于 3，则块内的数全部能选
// 如果段长度 > 2，则这一段只能选两个，并且一定能选头尾两个数，具体证明见官解

// 我们将能选的数标为 1，不能选的数标为 0，问题就变成了区间中 1 的个数
// 值得注意到是，有两个例外：如果开头或者结尾的数是 0，但是实际上这两个数一定是可选的
// 因此在计算前缀和的时候，要将最后一个数和第一个数排除
// 这样会出现 l == r 时区间长度为负的情况，因此需要特判

// 时间复杂度：O(n + q)
// 空间复杂度：O(n)

// 另有莫队解法，见数据结构题单

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


func solve() {
	n, q := II(), II()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = II()
	}

	for i := 0; i < n; {
		start := i
		i++
		for i < n && a[i - 1] >= a[i] {
			i++
		}
		a[start], a[i - 1] = 0, 0
		for j := start + 1; j < i - 1; j++ {
			a[j] = 1
		}
	}

	s := make([]int, n + 1)
	for i, x := range a {
		s[i + 1] = s[i] + x
	}

	for Q := 0; Q < q; Q++ {
		l, r := II() - 1, II() - 1
		if l == r {
			Println(1)
		} else {
			Println(r - l + 1 - s[r] + s[l + 1])
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
