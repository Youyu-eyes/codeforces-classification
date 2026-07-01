package main

func solve() {
	n, L, s := II(), II(), II()
	P := newMatrix(n, n)
	Q := newMatrix(n, n)
	for i := range P {
		for j := range P[i] {
			P[i][j] = II()
		}
	}
	for i := range Q {
		for j := range Q[i] {
			Q[i][j] = II()
		}
	}

	M := newMatrix(2 * n, 2 * n)
	for x := 0; x < n; x++ {
		M[2 * x + 1][2 * x] = P[x][x]
		M[2 * x + 1][2 * x + 1] = Q[x][x]
		for y := 0; y < n; y++ {
			if y != x {
				M[2 * x][2 * y] = P[y][x]
				M[2 * x][2 * y + 1] = Q[y][x]
			}
		}
	}

	f1 := newMatrix(2 * n, 1)
	f1[2 * (s - 1)][0] = 1

	fL := matMul(matQpow(M, L - 1, MOD), f1, MOD)
	for i := 0; i < n; i++ {
		Print((fL[2 * i][0] + fL[2 * i + 1][0]) % MOD)
		Print(" ")
	}
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
				c[i][j] = int((uint64(c[i][j]) + uint64(aik)*uint64(b[k][j])) % uint64(mod))
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
