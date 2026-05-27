// https://codeforces.com/problemset/problem/13/E
// 2700
// 分块，根号分解算法

// 对于本题，暴力地想，操作 0 更新 power，需要 O(1) 时间
// 操作 1 从 0 开始跳到 n - 1，最多跳 O(n) 次
// 假设没有更新只有查询，总体复杂度为 O(nq) 这是我们不能接受的

// 本题的跳跃是一个链表，非区间，我们需要单点修改和单点查询操作，很难用线段树维护
// 需要均衡更新和查询操作，考虑分块

// 假设将 holes 分成若干块，每块的长度为 B，共 n / B 块
// 我们希望跳跃操作能直接从一个块跳跃到另一个块，这样我们只需要最多跳 n / B 次
// 因此，我们维护 j = next[i] 表示从 i 跳到 j，j 在下一个块中
// 同时，从 i 跳到 j 可能需要在自己的块中跳跃几次，如果我们不模拟这个过程，则需要 O(1) 的快速查询在本块中的跳跃次数，再进行累加
// 维护 count[i] 表示从 i 跳到 j 所需的步数

// 此时跳跃操作就变成了，查询 next[i]，并将 count[i] 累计进答案
// 注意边界，我们需要找到最后一次跳跃出 holes 的下标，因此在跳到最后一块时，无法保证跳到的 end 下标就是最终的下标
// 但是最终的下标一定与 end 在同一块
// 我们模拟 end 在自己的块内跳跃，最终抵达边界。由于最终下标与 end 同块，最多跳 B 次
// 操作 1 的总体时间复杂度为 O(n / B + B)

// 更新操作，我们需要 O(1) 更新 power 数组的同时，维护 next 和 count 数组
// 假设已经更新 power 数组，如果 i + power[i] 与 i 在同一块
// 则 next[i] = next[i + power[i]]，count[i] = count[i + power[i]] + 1
// 如果不在同一块，则
// next[i] = i + power，count[i] = 1
// 同时，我们需要维护同一块中，在 i 之前的下标的信息
// 根据更新公式，类比状态转移方程，由于转移来源在 i 后面，我们需要倒序枚举
// 更新同一个块，需要 O(B) 时间，总体复杂度 O(B)

// 现在我们可以得到单次操作的时间复杂度
// O(B + n / B + B) = O(n / B + 2 * B)，当 B = sqrt(n / 2) 时，取最小值
// 为了保证 B > 0，B = isqrt(ceil(n / 2)) = isqrt((n + 1) / 2)

// 时间复杂度：O(2sqrt(2n))
// 空间复杂度：O(n)

package main

import (
	"bufio"
	"fmt"
	"os"
	"math"
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
	n, m := II(), II()
	power := make([]int, n)
	for i := 0; i < n; i++ {
		power[i] = II()
	}

	B := isqrt((n + 1) / 2)
	next := make([]int, n)
	count := make([]int, n)

	update := func(i int) {
		p := power[i]
		if i + p < n && (i + p)/B == i/B {
			next[i] = next[i + p]
			count[i] = count[i + p] + 1
		} else {
			next[i] = i + p
			count[i] = 1
		}
	}

	for i := n - 1; i >= 0; i-- {
		update(i)
	}

	for M := 0; M < m; M++ {
		number := II()
		if number == 0 {
			a, b := II() - 1, II()
			power[a] = b
			for i := a; i >= a / B * B; i-- {
				update(i)
			}
		} else {
			start := II() - 1
			end := -1
			ans := 0
			for i := start; i < n; {
				ans += count[i]
				end = i
				i = next[i]
			}
			for ;end + power[end] < n; end += power[end] {}
			Println(end + 1, ans)
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