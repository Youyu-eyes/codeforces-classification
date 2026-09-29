> 本题解实际采用的是 Go 语言，您可以借助 AI 翻译，IO 模板见文末链接
# 思路
结论题：一个数组变为非降序序且每次仅允许交换相邻两个数，最少需要的交换次数为 **数组中逆序对个数**

## 证明
定义逆序对为 $i < j \land a_i > a_j$，一个非降序数组显然逆序对数为 $0$，因此我们需要消除所有逆序对，也就得到了非降序数组。

由于我们只能交换相邻两个数，因此考虑 $a_i$ 和 $a_{i+1}$，假定 $a_i > a_{i+1}$

首先，交换 $a_i$ 和 $a_{i+1}$ 会消除 $a_i$ 与 $a_{i+1}$ 产生的一个逆序对。

其次，考虑 $[0, i-1]$ 的数，它们与 $a_i$ 和 $a_{i+1}$ 的相对位置不变，如

$$
1\ 4\ 3\ 2 
$$

- 当你交换 $3$ 和 $2$ 时，$4$ 依旧在 $3$ 和 $2$ 的前面，形成的逆序对个数不会减少
- 同理交换后，$1$ 依旧在 $3$ 和 $2$ 前面，形成的逆序对个数不会增加

因此逆序对个数不变。

最后，考虑 $[i+2, n-1]$ 的数，它们与 $a_i$ 和 $a_{i+1}$ 的相对位置也不变，因此逆序对个数不变。

我们可以证明，交换一对 $a_i > a_{i+1}$ 的数只会减少一个逆序对，且一定会减少一个逆序对。

因此总交换次数就是逆序对个数。

## 回到本题
本题只能交换连续的 $3$ 个数，本质上就是交换 $a_{i}$ 与 $a_{i+2}$，可以发现下标奇偶性不变，即将数组按下标奇偶性拆分成两个数组，问题变成了：
给你两个数组，问使这两个数组升序（本题数组中无相同元素）所需的最小交换次数，每次只能交换相邻的两个数。

特殊情况，如果只含奇数的数组中有偶数，由于交换只能在下标奇偶性相同的数之间交换，因此不可能有序，直接返回 $-1$ 即可。

关于如何计算给定数组中的逆序对个数，本题的数据范围可以暴力 $\mathcal{O}(n^2)$ 计算，也可以用树状数组计算，本质是二维偏序问题。

关于树状数组计算逆序对个数，见 [我的题解](https://leetcode.cn/problems/count-subarrays-with-even-odd-ratio-ii/solutions/4005441/xiao-bai-si-lu-shu-zhuang-shu-zu-qiu-ni-d9vqk/)

## Code
```go [sol-Go]
func solve() {
	n := II()
	odd := make([]int, n / 2)
	even := make([]int, (n + 1) / 2)
	for i := range n {
		x := II() - 1
		if i & 1 != x & 1 {
			Println(-1)
			return
		}
		if i & 1 > 0 {
			odd[i >> 1] = x
		} else {
			even[i >> 1] = x
		}
	}
	Println(inversion(odd) + inversion(even))
}

func inversion(nums []int) (res int) {
	n := len(nums)
	bit := newFenwickTree(2 * n)
	for _, x := range nums {
		res += bit.query(x, 2 * n - 1)
		bit.update(x, 1)
	}
	return
} 

type fenwick []int

func newFenwickTree(n int) fenwick {
	return make(fenwick, n+1) // 使用下标 1 到 n
}

// a[i] 增加 val
// 时间复杂度 O(log n)
func (f fenwick) update(i, val int) {
	for i++; i < len(f); i += i & -i {
		f[i] += val
	}
}

// 求前缀和 a[1] + ... + a[i]
// 时间复杂度 O(log n)
func (f fenwick) pre(i int) (res int) {
	for i++; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}

// 求区间和 a[l] + ... + a[r]
// 时间复杂度 O(log n)
// 0 <= l <= r < n
func (f fenwick) query(l, r int) int {
	if l > r {
		return 0
	}
	return f.pre(r) - f.pre(l-1)
}

```

## 复杂度

- 时间复杂度: $\mathcal{O}(n\log n)$，其中 $n$ 表示数组长度。
- 空间复杂度: $\mathcal{O}(n)$

更多模板请见 [我的github仓库](https://github.com/Youyu-eyes/algorithm_go)，感谢关注