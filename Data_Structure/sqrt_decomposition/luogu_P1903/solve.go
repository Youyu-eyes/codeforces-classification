// https://www.luogu.com.cn/problem/P1903
// 带修莫队模板


package main

import (
	"math"
	"slices"
	"cmp"
)

func solve() {
	n, m := II(), II()
	nums := make([]int, n)
	for i := range n {
		nums[i] = II()
	}

	B := int(max(1, math.Pow(float64(n), 2.0/3.0)))

	type query struct {
		lid, rid int
		l, r, t  int
		qIdx     int
	}
	type version struct {
		pos, color int
	}

	qs := []query{}
	vs := []version{}
	curQ := 0
	for qi := range m {
		op := IB()
		if op == 'Q' {
			l, r := II() - 1, II() - 1
			qs = append(qs, query{l / B, r / B, l - 1, r + 1, qi - curQ, curQ})
			curQ++
		} else {
			idx, val := II() - 1, II()
			vs = append(vs, version{idx, val})
		}
	}

	slices.SortFunc(qs, func(a, b query) int {
		if a.lid != b.lid {
			return cmp.Compare(a.lid, b.lid)
		}
		if a.rid != b.rid {
			return cmp.Compare(a.rid, b.rid)
		}
		return cmp.Compare(a.t, b.t)
	})

	cnt := make([]int, MX)
	res := 0
	
	add := func(x int) {
		if cnt[x] == 0 {
			res++
		}
		cnt[x]++
	}

	del := func(x int) {
		cnt[x]--
		if cnt[x] == 0 {
			res--
		}
	}

	upd := func(q query, t int) {
		pos, color := vs[t].pos, vs[t].color
		if q.l < pos && pos < q.r {
			del(nums[pos])
			add(color)
		}
		nums[pos], vs[t].color = vs[t].color, nums[pos]
	}

	l, r, t := -1, 0, 0
	ans := make([]int, curQ)
	for _, b := range qs {
		for ; r < b.r; r++ {
			add(nums[r])
		}

		for ; l > b.l; l-- {
			add(nums[l])
		}

		for r > b.r {
			r--
			del(nums[r])
		}

		for l < b.l {
			l++
			del(nums[l])
		}

		for ; t < b.t; t++ {
			upd(b, t)
		}

		for t > b.t {
			t--
			upd(b, t)
		}

		ans[b.qIdx] = res
	}
	for i := range ans { Println(ans[i]) }
}
