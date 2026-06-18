// https://codeforces.com/problemset/problem/1117/D
// 2100
// 矩阵快速幂优化 DP

// 本题是标准的线性 DP，由于初始宝石数量无限，因此每个 n 至少有 1 个方案
// 注意，本题不是在已经选择的宝石中选择一个进行分裂，而是选择不在已选的方案里的宝石分裂
// 也就是，要么选 1 个魔法宝石，要么选 m 个普通宝石

// 定义 f[i] 为选择 i 个宝石的方案数总和
// f[i] = f[i - 1] + f[i - m]
// 由于 n 非常大，我们需要优化

// 求和优化 n，适用矩阵快速幂

// 本题是一维方程，我们需要构造列向量 V[i] 和 V[i - 1]
// V[i] = [[f[i]]
//         [f[i - 1]]
//         [f[i - 2]]
//         ...
//         [f[i - m + 1]]]

// V[i - 1] = [[f[i - 1]]
//             [f[i - 2]]
//             [f[i - 3]]
//             ...
//             [f[i - m]]]

// 根据状态转移方程，系数矩阵 M 为
// [1 0 0 ... 0 1]
// [1 0 0 ... 0 0]
// [0 1 0 ... 0 0]
// ...
// [0 0 0 ... 1 0]

// M 为 m × m 的方阵
// 从 v[n] 到 v[m - 1]，共 n - m + 1 个系数矩阵相乘

// 时间复杂度：O(m^3 logn)
// 空间复杂度：O(m^2)

// 更快的做法见 solve_BM

package main

func solve() {
	n, m := II(), II()
	M := newMatrix(m, m)
	M[0][0], M[0][m - 1] = 1, 1
	for i := 1; i < m; i++ {
		M[i][i - 1] = 1
	}

	f0 := newMatrix(m, 1)
	for i := range f0 {
		f0[i][0] = 1
	}

	if n < m {
		Println(1)
		return
	}

	fn := matMul(matQpow(M, n - m + 1, MOD), f0, MOD)
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
