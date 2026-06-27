// 矩阵快速幂优化 DP

// 时间复杂度：O(k^3 · log x)
// 空间复杂度：O(k^2)

package main

func solveMatrix() {
	x, p1, p2, p3, p4 := II(), II(), II(), II(), II()
	p := [4]int{p4, p3, p2, p1}
	inv := qpow(100 - p[0], -1, MOD)

	M := matrix{{inv * p[1] % MOD, inv * p[2] % MOD, inv * p[3] % MOD, inv * 100 % MOD},
			    {1, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 0, 1}}
	
	fx := matrix{{0},
				 {0},
				 {0},
				 {1}}
	
	f0 := matMul(matQpow(M, x, MOD), fx, MOD)

	Println(f0[0][0])
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
