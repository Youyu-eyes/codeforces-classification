// https://codeforces.com/contest/1628/problem/A
// 1400
// 贪心 + 分组循环

// https://leetcode.cn/problems/lexicographically-maximum-mex-array/description/
// leetcode 504 周赛

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
	MX = 200_005
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
	nums := make([]int, n)
	for i := range nums { nums[i] = II() }

    freq := [MX]int{}
    for _, x := range nums {
        freq[x]++
    }
    
    mex := 0
    for x := 0; freq[x] > 0; x++ { mex = x + 1 }

    vis := make([]int, MX)
    tag := 0
    l := 0
    var tmp_mex, x, r, need int

    ans := []int{}
    for l < n {
        tag++
        tmp_mex = mex
        need = mex
        if mex == 0 {
            for i := l; i < n; i++ {
                ans = append(ans, 0)
            }
            break
        }
        ans = append(ans, mex)

        for r = l; r < n && need > 0; r++ {
            x = nums[r]
            if vis[x] != tag && x < tmp_mex {
                need--
                vis[x] = tag
            }

            freq[x]--
            if freq[x] == 0 && x < mex {
                mex = x
            }
        }
        l = r
    }
    Println(len(ans))
	for _, x := range ans {
		Print(" ")
		Print(x)
	}
	Println()
}


func main() {
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	T := II() // 1

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
