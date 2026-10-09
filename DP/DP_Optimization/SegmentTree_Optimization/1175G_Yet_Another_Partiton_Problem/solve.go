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

	for K := 1; K <= k; K++ {
		st := []pair{}
		pLc := newPlc(1, n)
		nf := make([]int, n + 1)
		for i := range nf { nf[i] = inf }

		lh := LowerHull{}
		lh.add(vec{K - 1, f[K - 1]})
		st = append(st, pair{K - 1, lh})

		b := lh.queryBinary(vec{-a[K - 1], 1})
		pLc = append(pLc, pLc[len(pLc)-1].insert(line{float64(a[K - 1]), float64(b), K}, 1, n))
		
		for i := K; i <= n - (k - K); i++ {
			nf[i] = int(pLc[len(pLc)-1].query(i))

			if i == n {
				continue
			}

			newlh := LowerHull{}
			newlh.add(vec{i, f[i]})

			for len(st) > 0 && a[st[len(st)-1].idx] < a[i] {
				pLc = pLc[:len(pLc)-1]
				oldlh := st[len(st)-1].lh
				st = st[:len(st)-1]
				newlh = merge(oldlh, newlh)
			}
			st = append(st, pair{i, newlh})

			b := newlh.queryBinary(vec{-a[i], 1})
			newLine := line{float64(a[i]), float64(b), i}
			pLc = append(pLc, pLc[len(pLc)-1].insert(newLine, 1, n))	
		}

		copy(f, nf)
	}
	
	Println(f[n])
}
