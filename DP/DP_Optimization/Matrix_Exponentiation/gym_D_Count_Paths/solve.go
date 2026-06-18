// https://codeforces.com/gym/102644/problem/E
// 矩阵快速幂优化 DP

// DAG 有向无环图

// 状态转移方程为
// f[k][i] = sum{j in fa[x]}(f[k - 1][j])

// 时间复杂度：O(n^3 logk)
// 空间复杂度：O(n^2)

package main

func solve() {
	n, m, k := II(), II(), II()
	fa := make([][]int, n)
	for i := 0; i < m; i++ {
		u, v := II() - 1, II() - 1
		fa[v] = append(fa[v], u)
	}

	M := newMatrix(n, n)
	for i := 0; i < n; i++ {
		for _, j := range fa[i] {
			M[i][j] = 1
		}
	}

	f0 := newMatrix(n, 1)
	for i := range f0 {
		f0[i][0] = 1
	}

	fn := matMul(matQpow(M, k, MOD), f0, MOD)
	ans := 0
	for _, f := range fn {
		ans = (ans + f[0]) % MOD
	}

	Println(ans)
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
