// 预处理后缀 + 划分型DP
// 复杂度：O(kLn)，k 表示合法分段数，L 表示最长后缀长度，n 表示字符串总长度

package main

func solve() {
	n := II()
	s := []byte(IS())

	f0 := make([]int, n + 1)
	f1 := make([]int, n + 1)
	f0[0] = 1
	f1[0] = 1

	divide := make([]bool, n + 1)
	divide[n] = true
	cur := 0
	base := 1
	for m := n - 1; m >= max(0, n - 3); m-- {
		if !(s[m] >= '0' && s[m] <= '9') {
			break
		}
		cur = int(s[m] - '0') * base + cur
		base *= 10
		if !(m < n - 1 && s[m] == '0') && cur <= 128 {
			divide[m] = true
		}
	}

	ans := 0
	if divide[0] {
		ans += f1[0]
	}

	for k := 1; k <= 8; k++ {
		nf0 := make([]int, n + 1)
		nf1 := make([]int, n + 1)
		for i := k; i <= n; i++ {
			for j := max(0, i - 4); j < i; j++ {
				if s[j] != '0' || j == i - 1 {
 					nf0[i] += f0[j]
				}
			}
			nf1[i] = nf0[i]
			for j := max(0, i - 4); j < i; j++ {
				if s[j] != '0' || j == i - 1 {
 					nf1[i] += f1[j]
				}
			}
		}
		copy(f0, nf0)
		copy(f1, nf1)

		if k < 8 {
			for m := max(0, n - 3); m <= n; m++ {
				if divide[m] {
					ans += f1[m]
				}
			}
		}
	}

	for m := max(0, n - 3); m <= n; m++ {
		if divide[m] {
			ans += f0[m]
		}
	}

	Println(ans)
}