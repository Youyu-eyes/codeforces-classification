// https://codeforces.com/problemset/problem/2193/E
// 1300
// 调和级数枚举

// 由于本题可以重复选择一个相同的数，因此无需关心 nums 中数字出现的具体次数，只标记是否出现
// 由于预处理因子的第二重循环是调和级数枚举，最多枚举 log(x) 项

// 用相向双指针两两枚举 x 的因子，使得 x = a * b
// 由于我们从小到大处理询问，因此 a 和 b 的最少组成已经计算出来了
// a 需要 ans[a] 个数组成，b 需要 ans[b] 个数组成
// 根据加法原理，ans[x] = ans[a] + ans[b]
// 两两枚举 a, b，取所有 ans[a] + ans[b] 的最小值即为结果

// 有两种写法，预处理写法就是上述思路照搬，但是复杂度非严格 log(n) 级别

// 调和级数枚举写法复杂度严格为 log(n) 级别

// 关于调和级数枚举的证明见 proof.md

package main

// ======= 预处理写法 ======= //

// ------- 全局预处理变量 ------- //
var (
	divisors   [][]int
)

func init() {
	divisors = make([][]int, MX)

	// 4. 预处理因子
	for i := 1; i < MX; i++ {
		for j := i; j < MX; j += i { // 枚举 i 的倍数 j
			divisors[j] = append(divisors[j], i) // i 是 j 的因子
		}
	}
}

func solveInit() {
	n := II()
	cnt := make([]bool, n + 1)
	for range n {
		cnt[II()] = true
	}

	ans := make([]int, n + 1)
	ans[1] = 1
	for i := 1; i <= n; i++ {
		total := len(divisors[i])
		for j := range total / 2 + 1 {
			a, b := divisors[i][j], divisors[i][total - 1 - j]
			if a == 1 {
				if cnt[i] {
					ans[i] = 1
					continue
				} else {
					ans[i] = inf
				}
			} else {
				ans[i] = min(ans[i], ans[a] + ans[b])
			}
		}
	}
	if !cnt[1] {
		ans[1] = inf
	}

	for i := 1; i <= n; i++ {
		x := ans[i]
		if x < inf {
			Print(x)
		} else {
			Print(-1)
		}
		Print(" ")
	}
	Println()
}

// ======= 调和级数枚举 ======= //

func solve() {
	n := II()

	ans := make([]int, n + 1)
	for i := range ans {
		ans[i] = inf
	}
	for range n {
		ans[II()] = 1
	}

	for i := 1; i <= n; i++ {
		if ans[i] == inf {
			continue
		}

		for j := 2 * i; j <= n; j += i {
			ans[j] = min(ans[j], ans[i] + ans[j / i])
		}
	}

	for i := 1; i <= n; i++ {
		x := ans[i]
		if x < inf {
			Print(x)
		} else {
			Print(-1)
		}
		Print(" ")
	}
	Println()
}