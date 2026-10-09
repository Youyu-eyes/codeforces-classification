// https://www.luogu.com.cn/problem/P3834
// 可持久化线段树：静态区间第 k 小
// 见 lc3762 可持久化线段树

package main

import (
	"slices"
	"sort"
	"runtime/debug"
)

// 卡常，需要关闭 GC
func init() { debug.SetGCPercent(-1) } 

func solve() {
	n, m := II(), II()
	nums := make([]int, n)
	for i := range n {
		nums[i] = II()
	}
	pSt, unique := newPst(nums)
	
	for range m {
		l, r, k := II() - 1, II(), II() - 1
		i := pSt[r].kth(pSt[l], k)
		Println(unique[i])
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

func buildPst(l, r int) *pNode {
	o := &pNode{l: l, r: r}
	if l == r {
		return o
	}
	m := (l + r) >> 1
	o.lo = buildPst(l, m)
	o.ro = buildPst(m+1, r)
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
		return o.l
	}
	cntL := o.lo.cnt - old.lo.cnt
	if k < cntL {
		return o.lo.kth(old.lo, k)
	}
	return o.ro.kth(old.ro, k-cntL)
}

func newPst(a []int) ([]*pNode, []int) {
	unique := slices.Clone(a)
	slices.Sort(unique)
	unique = slices.Compact(unique)

	t := make([]*pNode, len(a)+1)
	t[0] = buildPst(0, len(unique)-1)
	for i, v := range a {
		j := sort.SearchInts(unique, v)
		t[i+1] = t[i].update(j, v)
	}
	return t, unique
}
