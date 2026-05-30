// https://codeforces.com/contest/1514/problem/D
// 2000
// 普通莫队 / 回滚莫队

// 首先本题基于绝对众数的相关结论，一个数组中枚举 2 个数删除，最多操作 min(n / 2, n - m) 其中 n 表示数组长度，m 表示众数个数
// 对于本题，我们可以选择一个众数加一个非众数，然后组成一个序列

// 此时，如果非众数被选完了，众数还有剩余，则我们将剩余的众数每个组成一个子序列，剩余的众数为 m - (n - m) = 2 * m - n 个
// 然后将一个众数添加到刚开始的序列中，这样最初的序列恰好有 ceil(m / 2) 个众数，满足题意
// 共有 2 * m - n - 1 + 1 = 2 * m - n 个子序列

// 如果没有剩余，则说明没有绝对众数，那么子数组就可以作为合法子序列，个数为 1

// 二者取 max 即为答案

// 先讲容易理解的回滚莫队
// 由于我们需要求众数的个数，众数的定义是所有元素出现次数最大的元素值
// 众数的个数就是出现的最大次数
// 由于存在删除操作，max 难以撤销，考虑回滚莫队直接优化掉删除操作

// 维护 cnt 表示元素出现次数，本题保证 a[i] < n 无需离散化
// 右端点右移，左端点左移，统计元素出现次数的同时，维护 cnt 的最大值
// 然后计算答案，再回滚即可

// 时间复杂度：O(n × sqrt(q))，其中 n 表示数组长度，q 表示询问次数
// 空间复杂度：O(n)

// 另有普通莫队做法，见下 solveMo

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


func solve() {
	n, q := II(), II()
	nums := make([]int, n)
	for i := range nums { nums[i] = II() }

	B := int(math.Ceil(float64(n) / math.Sqrt(float64(q))))

	// 双开区间莫队
	type query struct {
		id int
		l, r int // (l, r)
		qIdx int
	}
	qs := []query{}

	cnt := make([]int, n + 1)
	major := 0

	ans := make([]int, q)
	for Q := 0; Q < q; Q++ {
		l, r := II() - 1, II()

		// 大区间离线
		if r - l > B {
			qs = append(qs, query{l / B, l - 1, r, Q})
			continue
		}

		// 小区间暴力
		for i := l; i < r; i++ {
			cnt[nums[i]]++
			major = max(major, cnt[nums[i]])
		}
		ans[Q] = max(1, 2 * major - (r - l))

		for i := l; i < r; i++ {
			cnt[nums[i]]--
		}
		major = 0
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
			clear(cnt)
			major = 0
		}
		
		// 右端点右移
		for ; r < b.r; r++ {
			cnt[nums[r]]++
			major = max(major, cnt[nums[r]])
		}

		// 保留状态
		tmp := major

		// 左端点左移
		for ; l > b.l; l-- {
			cnt[nums[l]]++
			major = max(major, cnt[nums[l]])
		}

		ans[b.qIdx] = max(1, 2 * major - (b.r - b.l - 1))

		// 回滚
		major = tmp
		l = start - 1
		for j := b.l + 1; j <= l; j++ {
			cnt[nums[j]]--
		}
	}

	for _, x := range ans { Println(x) }
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


// 普通莫队
// 思路是维护一个 freq 频率数组，表示 cnt[x] 出现的次数，用于删除操作中维护 major

// 这里重要的一个结论是，freq[cnt[x]]-- 后，如果 freq[major] == 0，则 cnt[x] 一定为 major
// 因为如果 cnt[x] != major，则 freq[major] 在删除操作中没有被更改，但是之前的操作一定保证 freq[major] > 0，这与 freq[major] = 0 矛盾
// 假设不成立，因此 cnt[x] = major

// 然后由于 freq[cnt[x]--]++ = freq[major--]++，因此 freq[major--] > 0，所以我们不需要枚举全部的 cnt[x] 来更新 major，major 一定是 cnt[x]--
// 这样的话，删除的时间复杂度就是 O(1) 的

// 总体时空复杂度与回滚莫队一致

func solveMo() {
	n, q := II(), II()
	nums := make([]int, n)
	for i := range nums { nums[i] = II() }

	B := int(math.Ceil(float64(n) / math.Sqrt(float64(q))))

	// 双开区间莫队
	type query struct {
		id int
		l, r int // (l, r)
		qIdx int
	}
	qs := []query{}

	
	ans := make([]int, q)
	for Q := 0; Q < q; Q++ {
		l, r := II() - 2, II()
		qs = append(qs, query{l / B, l, r, Q})
	}

	// 奇偶排序优化
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

	l, r, res := B - 1, B, 0
	cnt := make([]int, n + 1)
	freq := make([]int, n + 1)
	major := 0

	add := func(x int) {
		freq[cnt[x]]--

		cnt[x]++
		freq[cnt[x]]++

		if major < cnt[x] {
			major = cnt[x]
		}
	}

	remove := func(x int) {
		freq[cnt[x]]--

		cnt[x]--
		freq[cnt[x]]++

		if freq[major] == 0 {
			major--
		}
	}

	for _, b := range qs {		
		// 右端点右移
		for ; r < b.r; r++ {
			add(nums[r])
		}

		// 左端点左移
		for ; l > b.l; l-- {
			add(nums[l])
		}

		// 右端点左移
		// 开区间，先左移再删除
		for r > b.r {
			r--
			remove(nums[r])
		}

		// 左端点右移
		// 开区间，先右移再删除
		for l < b.l {
			l++
			remove(nums[l])
		}

		res = max(1, 2 * major - (b.r - b.l - 1))

		ans[b.qIdx] = res
	}

	for _, x := range ans {
		Println(x)
	}
}