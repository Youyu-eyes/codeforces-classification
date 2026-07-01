# 思路
> 首先声明，以下题解的标签全部改成 $0-base$ 下的数字

本题是概率 DP 问题，分为两类情况
1. 如果 $v_{t} = v_{t - 1}$ 则第 $t + 1$ 个视频为强化模式
2. 如果 $v_{t} \ne v_{t - 1}$ 则第 $t + 1$ 个视频为标准模式

因此我们定义：
$f[i][x][k]$ 为当看到第 $i$ 个视频，标签为 $x$ 时，当前视频的状态为 $k$ 时的概率。
定义 $k = 0$ 表示标准模式，$k = 1$ 表示强化模式。

则状态转移方程为：

$$
\begin{aligned}
& f[i][x][0] = \sum_{\substack{y \in [0, n - 1] \\ y \ne x}}(f[i - 1][y][0] * P[y][x] + f[i - 1][y][1] * Q[y][x]) \\
& f[i][x][1] = f[i - 1][x][0] * P[x][x] + f[i - 1][x][1] * Q[x][x]
\end{aligned}
$$

注意这里有一个误区：由于题面说明，第 $i$ 个视频的状态取决于前两个视频的状态，但是实际上不需要 $i - 2$ 作为转移来源。

因为我们需要将标准状态与强化状态区分开，需要定义 $k$ 这个参数额外维护，$f[i - 1][y][0]$ 就已经表示第 $i - 1$ 个视频与第 $i - 2$ 个视频的标签是不同的，不需要额外转移。

而 $f[i - 1][y][1]$ 表示第 $i - 1$ 个视频与第 $i - 2$ 个视频的标签是相同的。

---

## 矩阵快速幂优化
由于本题的 $n$ 不大，且 $L$ 很大，上述状态转移方程满足结合率，考虑矩阵快速幂优化。

首先进行状态压缩，将后两个维度压缩至一维，具体方法及证明可见 [题解](https://leetcode.cn/problems/number-of-zigzag-arrays-ii/solutions/3985024/xiao-bai-si-lu-dai-ni-ru-men-zhuang-tai-5a8qv/)

压缩后，状态转移方程为：
$$
\begin{aligned}
& f[i][2x] = \sum_{\substack{y \in [0, n - 1] \\ y \ne x}}(f[i - 1][2y] * P[y][x] + f[i - 1][2y+1] * Q[y][x]) \\
& f[i][2x+1] = f[i - 1][2x] * P[x][x] + f[i - 1][2x+1] * Q[x][x]
\end{aligned}
$$

枚举标签 $x$ 与标签 $y$
- $(2x, 2y)$ 填 $P[y][x]$
- $(2x, 2y+1)$ 填 $Q[y][x]$
- $(2x+1, 2x)$ 填 $P[x][x]$
- $(2x+1, 2x+1)$ 填 $Q[x][x]$

定义系数矩阵 $M$
将整个 $f$ 压缩成 $v$ 列向量后，可得
$$
\begin{aligned}
& v[i] = M\cdot v[i - 1]
\end{aligned}
$$

更特殊的
$$
\begin{aligned}
& v[L] = M\cdot v[L - 1] \\
& v[L] = M\cdot M\cdot v[L - 2] \\
& \vdots \\
& v[L] = M^{L - 1}\cdot v[1]
\end{aligned}
$$

注意本题是第一个视频开始转移，不是 $v[0]$.

第一个视频的标签为 $s$，因此压缩后 $f[1][2 * s] = 1$，其余元素均为 $0$.

每个标签的答案为标准模式概率 $+$ 强化模式概率，注意压缩。

## 复杂度

- 时间复杂度: $\mathcal{O}(n^3\times \log L)$
- 空间复杂度: $\mathcal{O}(n^2)$

更多模板请见 [我的github仓库](https://github.com/Youyu-eyes/algorithm_go)，感谢关注