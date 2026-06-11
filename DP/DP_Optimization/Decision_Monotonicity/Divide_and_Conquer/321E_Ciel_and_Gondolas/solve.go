// https://codeforces.com/problemset/problem/321/E
// 2600
// 分治优化 DP

// 常规的划分型 DP，定义 f[k][i] 前 i 个数选 k 个数的最小代价
// 假设上一段划分的左端点 j
// f[k][i] = min{j = [k, i - 1]}(f[k - 1][j] + w(j, i))

// u[j][i] 表示 i 与 j 的不熟悉程度，抽象为选择 i 和 j 的代价
// 由于 j 到 i 的所有人都坐一辆车，总代价为

//   u[j][j]     + u[j][j + 1]     + u[j][j + 2]     + …… + u[j][i - 1]
// + u[j + 1][j] + u[j + 1][j + 1] + u[j + 1][j + 2] + …… + u[j + 1][i - 1]
// + u[j + 2][j] + u[j + 2][j + 1] + u[j + 2][j + 2] + …… + u[j + 2][i - 1]
// ……
// + u[i - 1][j] + u[i - 1][j + 1] + u[i - 1][j + 2] + …… + u[i - 1][i - 1]

// 由于我们 i 和 j 都枚举了两次，有一半的数据，如 u[i][j] 与 u[j][i]
// 由于矩阵是对称的，所以上述总代价为实际总代价的一半

// 但是为了方便计算，(每个代价 ÷ 2) 之和 = 代价之和 ÷ 2
// 因此我们定义 f[k][i] 为实际代价的两倍，最后返回答案 f[K][n] / 2

// 现在让我们观察代价函数 w
// 显然，总代价 w(j, i) 是 u 矩形中一个子矩形的和
// 求子矩形的和，可用二维前缀和优化，关于二维前缀和模板，见数据结构部分
// 二维前缀和讲解，见 leetcode 304

// 由于区间左闭右开，代价函数为
// w(j, i) = pre[i][i] - pre[j][i] - pre[i][j] + pre[j][j]
// 由于矩阵对称
// w(j, i) = pre[i][i] + pre[j][j] - 2 × pre[i][j]

// f[k][i] = min{j = [k, i - 1]}(f[k - 1][j] + pre[i][i] + pre[j][j] - 2 × pre[i][j])

// 观察到，由于 j 是作为下标与 i 一同计算前缀和，因此完全不可分离
// 无法写成向量点积形式

// 观察代价函数，定义 0 <= j <= j + 1 <= i <= i + 1 <= n
// 我们需要证明 四边形不等式
// w(j, i) + w(j + 1, i + 1) <= w(j, i + 1) + w(j + 1, i) 这里注意，max 方程为 >=

// 几何方法就是直接绘图，会发现左边比右边少了两个小矩形，边长为 1

// 下面给出代数法

//    s[i][i] + s[j][j] - 2s[i][j] + s[i + 1][i + 1] + s[j + 1][j + 1] - 2s[i + 1][j + 1]
// <= s[i + 1][i + 1] + s[j][j] - 2s[i + 1][j] + s[i][i] + s[j + 1][j + 1] - 2s[i][j + 1]

//    s[i][j] + s[i + 1][j + 1]
// >= s[i + 1][j] + s[i][j + 1]

// s[i][j + 1] - s[i][j] <= s[i + 1][j + 1] - s[i + 1][j]
// u[0][j] + u[1][j] + …… + u[i][j]
// u[0][j] + u[1][j] + …… + u[i][j] + u[i + 1][j]

// 因此右边大于左边
// 证毕

// 现在，由于平行四边形不等式成立
// f[x] 满足决策单调性，最优决策点单调不减
// 因此考虑 分治优化 DP

// 细节见代码

// 时间复杂度：O(k × n × logn)
// 空间复杂度：O(n^2)

package main

type Matrix [][]int

func Build(matrix [][]int) Matrix {
    m, n := len(matrix), len(matrix[0])
    sum := make([][]int, m+1)
    sum[0] = make([]int, n+1)
    for i, row := range matrix {
        sum[i+1] = make([]int, n+1)
        for j, x := range row {
            sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j] + x
        }
    }
    return sum
}

func solve() {
	n, K := II(), II()
	u := make(Matrix, n)
	for i := range u {
		u[i] = make([]int, n)
		for j := range u[i] {
			u[i][j] = II()
		}
	}

	pre := Build(u)
	f := make([]int, n + 1)
	for i := range f { f[i] = inf }
	f[0] = 0
	nf := make([]int, n + 1)

	w := func(j, i int) int {
		return pre[i][i] + pre[j][j] - 2 * pre[i][j]
	}

	var op func(int, int, int, int)
	op = func(l, r, optL, optR int) {
		if l > r {
			return
		}

		mid := l + (r - l) >> 1
		best := -1
		for j := optL; j <= min(optR, mid - 1); j++ {
			if best < 0 || f[j] + w(j, mid) < f[best] + w(best, mid) {
				best = j
			}
		}
		nf[mid] = f[best] + w(best, mid)
		op(l, mid - 1, optL, best)
		op(mid + 1, r, best, optR)
	}

	for k := 1; k <= K; k++ {
		nf = make([]int, n + 1)
		op(k, n, k - 1, n - 1)
		copy(f, nf)
	}

	Println(f[n] / 2)
}