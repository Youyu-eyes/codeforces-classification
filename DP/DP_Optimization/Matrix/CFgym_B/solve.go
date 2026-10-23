// https://codeforces.com/gym/102644/problem/B
// 预估 1900
// 矩阵快速幂优化 DP

// 定义 f[i][0] 为填了 i 个字符后状态为开心
// 定义 f[i][1] 为填了 i 个字符后状态为悲伤

// 定义集合 happy = {A E I O U H}，sad = {A E I O U S D}
// withoutHappy 和 withoutSad 分别为 happy 与 sad 的补集

// f[i][0] = sum{c in happy}(f[i - 1][1]) + sum{c in withoutSad}  (f[i - 1][0])
// f[i][1] = sum{c in sad}  (f[i - 1][1]) + sum{c in withoutHappy}(f[i - 1][1])

// 由于无论 c 为什么，f[i][k] 与 c 无关，所以
// sum{c in happy}(f[i - 1][1]) = 6 × f[i - 1][1]

// 同理，状态转移方程可以写成
// f[i][0] = 6 × f[i - 1][1] + 19 × f[i - 1][0]
// f[i][1] = 7 × f[i - 1][0] + 20 × f[i - 1][1]

// 写成矩阵乘法的形式，可以用矩阵快速幂优化

// 时间复杂度：O(logn)
// 空间复杂度：O(1)

package main

var M = matrix{
	{19, 6},
	{7, 20},
}

var f0 = matrix{
	{1},
	{0},
}

func solve() {
	n := II()
	fn := matMul(matQpow(M, n, MOD), f0, MOD)
	Println(fn[0][0])
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
