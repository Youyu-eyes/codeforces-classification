// https://codeforces.com/problemset/problem/148/D
// 1800
// 概率DP

// 定义 f[i][j] 为袋子里有 i 只白色老鼠，j 只黑色老鼠，下一次是公主抽，公主赢的概率
// 初始状态：如果袋子里全是白老鼠，公主一定赢
// 			如果袋子里没有白老鼠，公主一定输

// 现在，我们分类讨论

// 1. 首先累加从 i+j 只老鼠中抽出 i 只白色老鼠的概率
// 2. 这一次双方都没有赢，则分两种情况：

// 	  1. 公主抽黑色，龙抽黑色，跑了一只黑色（需要至少 3 只黑色）
//    2. 公主抽黑色，龙抽黑色，跑了一只白色（需要至少 1 只白色，2 只黑色）

// 时间复杂度：O(w × b)
// 空间复杂度：O(w × b)

package main

func solve() {
	w, b := II(), II()
	f := make([][]float64, w + 1)
	for i := range f {
		f[i] = make([]float64, b + 1)
		f[i][0] = 1.0
	}
	f[0][0] = 0.0

	for i := 1; i <= w; i++ {
		for j := 1; j <= b; j++ {
			f[i][j] += float64(i) / float64(i + j)
			if j >= 3 {
				f[i][j] += float64(j) / float64(i + j) * float64(j - 1) / float64(i + j - 1) * float64(j - 2) / float64(i + j - 2) * f[i][j - 3]
			}
			if i >= 1 && j >= 2 {
				f[i][j] += float64(j) / float64(i + j) * float64(j - 1) / float64(i + j - 1) * float64(i) / float64(i + j - 2) * f[i - 1][j - 2]
			}
		}
	}

	Println(f[w][b])
}