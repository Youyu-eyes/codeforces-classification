// https://codeforces.com/contest/1175/problem/G
// 3000
// 单调栈分段 + 可持久化李超线段树 + 凸包优化

// 见 leetcode 1959 我的题解
// https://leetcode.cn/problems/minimum-total-space-wasted-with-k-resizing-operations/solutions/3985960/xiao-bai-si-lu-oknlogn-fu-za-zuo-fa-go-b-w523/

package main

type pair struct {
	idx int
	lh LowerHull
}

func solve() {
	n, k := II(), II()
	a := make([]int, n)
	for i := range a { a[i] = II() }

	f := make([]int, n + 1)
	for i := range f { f[i] = inf }
	f[0] = 0
	nf := make([]int, n + 1)

	PlcSegtree := NewPersistentLiChaoTree(1, n)

	for K := 1; K <= k; K++ {
		st := deque[pair]{}
		PlcSegtree.clear()
		clear(nf)
		for i := range nf { nf[i] = inf }

		lh := LowerHull{}
		lh.add(vec{K - 1, f[K - 1]})
		st.pushBack(pair{K - 1, lh})

		b := lh.queryBinary(vec{-a[K - 1], 1})
		PlcSegtree.insert(Line{float64(a[K - 1]), float64(b), K})
		
		for i := K; i <= n - (k - K); i++ {
			nf[i] = int(PlcSegtree.query(i))

			if i == n {
				continue
			}

			newlh := LowerHull{}
			newlh.add(vec{i, f[i]})

			for !st.empty() && a[st.back().idx] < a[i] {
				PlcSegtree.rollBack()
				oldlh := st.back().lh
				st.popBack()
				newlh = merge(oldlh, newlh)
			}
			st.pushBack(pair{i, newlh})

			b := newlh.queryBinary(vec{-a[i], 1})
			newLine := Line{float64(a[i]), float64(b), i}
			PlcSegtree.insert(newLine)	
		}

		copy(f, nf)
	}
	
	Println(f[n])
}