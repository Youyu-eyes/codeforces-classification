// https://codeforces.com/gym/102644/problem/E
// 预估 1950

// 先思考如果恰好 k 步怎么做
// f[k][i][j] = f[k - 1][i - 1][j - 2]
//              f[k - 1][i + 1][j - 2]
//              ...

// 枚举一共 8 个格子可以转移，注意不能越界
// 将第 2 维与第 3 维合并为 8 × i + j
// 写成矩阵乘法形式即可

// 但是我们需要求最多 k 步，即 0 - k 的所有值求和
// 答案为 f0 + M × f0 + M^2 × f0 + ... + M^k × f0

// 现在，我们构造增广矩阵 A
// A = [M 0]
//     [M E]

// M 表示系数矩阵，E 表示单位阵，0 表示零矩阵

// v0 = [f0]
//      [f0]

// vn = [f0 × M]
//      [f0 + f0 × M]

// 因此，vn 的后半部分就存储了答案
// 对 vn 的下半部分求和即可

// 时间复杂度：O(k^3 logn)
// 空间复杂度：O(k^2)
// k = 128

package main

var M = newMatrix(64, 64)
var f0 = newMatrix(64, 1)

var A = newMatrix(128, 128)
var v0 = newMatrix(128, 1)

func init() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if j - 2 >= 0 {
				if i - 1 >= 0 {
					M[8 * i + j][8 * (i - 1) + j - 2] = 1
				}
				if i + 1 < 8 {
					M[8 * i + j][8 * (i + 1) + j - 2] = 1
				}
			}

			if j - 1 >= 0 {
				if i - 2 >= 0{
					M[8 * i + j][8 * (i - 2) + j - 1] = 1
				}
				if i + 2 < 8 {
					M[8 * i + j][8 * (i + 2) + j - 1] = 1
				}
			}

			if j + 1 < 8 {
				if i - 2 >= 0 {
					M[8 * i + j][8 * (i - 2) + j + 1] = 1
				}
				if i + 2 < 8 {
					M[8 * i + j][8 * (i + 2) + j + 1] = 1
				}
			}

			if j + 2 < 8 {
				if i - 1 >= 0 {
					M[8 * i + j][8 * (i - 1) + j + 2] = 1
				}
				if i + 1 < 8 {
					M[8 * i + j][8 * (i + 1) + j + 2] = 1
				}
			}
		}
	}

	f0[0][0] = 1

	for i := 0; i < 128; i++ {
		for j := 0; j < 64; j++ {
			if i < 64 {
				A[i][j] = M[i][j]
			} else {
				A[i][j] = M[i - 64][j]
			}
		}

		if i < 64 {
			v0[i][0] = f0[i][0]
		} else {
			v0[i][0] = f0[i - 64][0]
		}
	}

	for i := 64; i < 128; i++ {
		A[i][i] = 1
	}
}

func solve() {
	k := II()
	vn := matMul(matQpow(A, k, MOD), v0, MOD)
	ans := 0
	for i := 64; i < 128; i++ {
		ans = (ans + vn[i][0]) % MOD
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
