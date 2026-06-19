// https://codeforces.com/problemset/problem/621/E
// 2000
// 刷表法 + 矩阵快速幂

// 假设拼接上 j 后的余数为 cur，拼接前为 pre
// cur = (pre * 10 + j) % x

// 显然，如果知道 cur 求 pre 很难，但是知道 pre 求 cur 很容易
// 因此，我们枚举的是状态转移来源，适用刷表法

// f[i][cur] 为选择 i 个元素后余数为 cur 的操作次数
// f[i][cur] = sum{j in set}(cnt[j] * f[i - 1][pre])
// cur = (pre * 10 + j) % x
// 每个 pre 会给 cur 贡献 cnt[j]
// 注意给系数矩阵 M 赋值的时候要累加

// 时间复杂度：O(x^3 logn)
// 空间复杂度：O(x^2)

package main

func solve() {
	n, b, k, x := II(), II(), II(), II()
	cnt := [10]int{}
	for i := 0; i < n; i++ {
		cnt[II()]++
	}
	M := newMatrix(x, x)
	for pre := 0; pre < x; pre++ {
		for j, v := range cnt {
			cur := (pre * 10 + j) % x
			M[cur][pre] += v
		}
	}

	f0 := newMatrix(x, 1)
	f0[0][0] = 1

	fb := matMul(matQpow(M, b, MOD), f0, MOD)
	Println(fb[k][0])
}

// ------- 矩阵快速幂 ------- //

type matrix [][]int

// 生成 m × n 的矩阵 a
func newMatrix(m, n int) matrix {
	a := make(matrix, m)
	for i := range a {
		a[i] = make([]int, n)
	}
	return a
}

// 返回矩阵 a 和矩阵 b 相乘的结果，若 mod > 0 则取模
func matMul(a, b matrix, mod int) matrix {
	n, m := len(a), len(b[0])
	c := make(matrix, n)
	for i := 0; i < n; i++ {
		c[i] = make([]int, m)
		for k := 0; k < len(a[0]); k++ {
			if a[i][k] == 0 {
				continue
			}
			aik := a[i][k]
			for j := 0; j < m; j++ {
				c[i][j] += aik * b[k][j]
				if mod > 0 {
					c[i][j] %= mod
				}
				// 如果溢出
				// c[i][j] = int((uint64(c[i][j]) + uint64(aik)*uint64(b[k][j])) % uint64(mod))
			}
		}
	}
	return c
}

// 计算 A^n，若 mod > 0 则所有乘法取模
func matQpow(a matrix, n int, mod int) matrix {
	size := len(a)
	res := make(matrix, size)
	for i := 0; i < size; i++ {
		res[i] = make([]int, size)
		res[i][i] = 1
	}

	for n > 0 {
		if n&1 == 1 {
			res = matMul(a, res, mod)
		}
		a = matMul(a, a, mod)
		n >>= 1
	}
	return res
}
