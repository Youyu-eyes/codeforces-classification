package main

import (
	"slices"	
)

func solve() {
	n := II()
	s := make([]int, n)
	a := make([]int, n)
	for i := range n {
		s[i] = II()
	}
	for i := range n {
		a[i] = II()
	}
	pst := newPst(a)

	ans := make([]int, n + 1)
	var op func(int, int, int, int)
	op = func(l, r, optL, optR int) {
		if l > r {
			return
		}
		mid := (l + r) >> 1
		bestV, best := -1, optL
		for i := max(optL, mid - 1); i <= optR; i++ {
			val := pst[i].kth(pst[0], mid-1) + 2 * s[i] + a[i]
			if val > bestV {
				bestV = val
				best = i
			}
		}

		ans[mid] = bestV
		op(l, mid - 1, optL, best)
		op(mid + 1, r, best, optR)
	}

	op(1, n, 0, n - 1)
	for i := 1; i <= n; i++ {
		Println(ans[i])
	}
}

type pInfo struct{
    cnt int
    sum int
}

type pNode struct {
    lo, ro *pNode
	l, r   int
	pInfo
}

func (pNode) mergeInfo(l, r pInfo) pInfo {
	return pInfo{l.cnt + r.cnt, l.sum + r.sum}
}

func (pNode) diffInfo(l, r pInfo) pInfo {
	return pInfo{r.cnt - l.cnt, r.sum - l.sum}
}

func (o *pNode) maintain() {
	o.pInfo = o.mergeInfo(o.lo.pInfo, o.ro.pInfo)
}

func buildPst(a []pInfo, l, r int) *pNode {
	o := &pNode{l: l, r: r}
	if l == r {
		if a != nil {
			o.pInfo = a[l]
		}
		return o
	}
	m := (l + r) >> 1
	o.lo = buildPst(a, l, m)
	o.ro = buildPst(a, m+1, r)
	o.maintain()
	return o
}

func (o pNode) update(i int, val int) *pNode {
	if o.l == o.r {
		o.cnt++
		o.sum += val
		return &o
	}
	m := (o.l + o.r) >> 1
	if i <= m {
		o.lo = o.lo.update(i, val)
	} else {
		o.ro = o.ro.update(i, val)
	}
	o.maintain()
	return &o
}

func (o *pNode) query(old *pNode, ql, qr int) pInfo {
	if ql <= o.l && o.r <= qr {
		return o.diffInfo(old.pInfo, o.pInfo)
	}
	m := (o.l + o.r) >> 1
	if qr <= m {
		return o.lo.query(old.lo, ql, qr)
	}
	if m < ql {
		return o.ro.query(old.ro, ql, qr)
	}
	lRes := o.lo.query(old.lo, ql, qr)
	rRes := o.ro.query(old.ro, ql, qr)
	return o.mergeInfo(lRes, rRes)
}

func (o *pNode) kth(old *pNode, k int) int {
	if o.l == o.r {
		return o.l * k
	}
	cntR := o.ro.cnt - old.ro.cnt
	if k <= cntR {
		return o.ro.kth(old.ro, k)
	}
	return o.ro.sum - old.ro.sum + o.lo.kth(old.lo, k-cntR)
}

func newPst(a []int) ([]*pNode) {
	mx := slices.Max(a)
	t := make([]*pNode, len(a)+1)
	t[0] = buildPst(nil, 0, mx+1)
	for i, v := range a {
		t[i+1] = t[i].update(v, v)
	}
	return t
}
