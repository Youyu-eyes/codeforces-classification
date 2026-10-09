// https://www.luogu.com.cn/problem/P3919

// 本题是可持久化数据结构的思想起源
// 但是对于可持久化线段树通过非常不友好，极度卡常

// 用到的优化方式都是只限于本题的
// 因此下面的可持久化线段树对于别的题目没有任何参考价值
// 可能这也是没进 2026 省选赛前模板题合集的原因
// 模板题还是要看 P3834


package main

import (
	"math/bits"
)

func solve() {
	n, m := II(), II()
	nums := make([]int, n)
	for i := range n {
		nums[i] = II()
	}

	maxNodes := m*(bits.Len(uint(n-1))+2) + 5
	sgt := &pst{
		nodes: make([]pNode, 1, maxNodes),
		a:     nums,
	}
	root0 := 0

	roots := make([]int, m+1)
	roots[0] = root0

	for qi := range m {
		v, q, idx := II(), II(), II()-1
		if q == 1 {
			val := II()
			roots[qi+1] = sgt.update(roots[v], idx, val, 0, n-1)
		} else {
			Println(sgt.get(roots[v], idx, 0, n-1))
			roots[qi+1] = roots[v]
		}
	}
}

type pInfo = int

type pNode struct {
	lo, ro int
	pInfo
}

type pst struct {
	nodes []pNode
	a     []int
}

func (t *pst) update(o, i, val int, l, r int) int {
	t.nodes = append(t.nodes, pNode{})
	new := len(t.nodes) - 1
	t.nodes[new] = t.nodes[o]

	if l == r {
		t.nodes[new].pInfo = val
		return new
	}
	m := (l + r) >> 1
	if i <= m {
		t.nodes[new].lo = t.update(t.nodes[o].lo, i, val, l, m)
	} else {
		t.nodes[new].ro = t.update(t.nodes[o].ro, i, val, m+1, r)
	}
	return new
}

func (t *pst) get(o, i, l, r int) pInfo {
	if o == 0 {
		return t.a[i]
	}
	if l == r {
		return t.nodes[o].pInfo
	}
	m := (l + r) >> 1
	if i <= m {
		return t.get(t.nodes[o].lo, i, l, m)
	}
	return t.get(t.nodes[o].ro, i, m+1, r)
}
