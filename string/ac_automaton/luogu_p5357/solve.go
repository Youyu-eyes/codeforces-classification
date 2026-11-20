// https://www.luogu.com.cn/problem/P5357
// 洛谷蓝 - 提高+/省选-

// AC 自动机

package main

func solve() {
	n := II()
	patterns := make([]string, n)
	for i := range n {
		patterns[i] = IS()
	}
	ac := newAcam(patterns)
	text := IS()
	cnt := ac.acCnt(text)
	for _, x := range cnt {
		Println(x)
	}
}

const acamNodeSize = 26

type acamNode struct {
	son [acamNodeSize]*acamNode

	fail *acamNode
	last *acamNode

	indices []int
	cnt int
	nodeId int
}

type acam struct {
	patterns []string

	root *acamNode
	nodeCnt int

	g [][]int
	inDeg map[*acamNode]int
}

func newAcam(patterns []string) *acam {
	ac := &acam{
		patterns: patterns,
		root: &acamNode{},
		nodeCnt: 1,
		inDeg: map[*acamNode]int{},
	}
	for i, s := range patterns {
		ac.put(s, i+1)
	}
	ac.buildFail()
	return ac
}

func (ac *acam) ord(c rune) rune { return c - 'a' }

func (ac *acam) put(s string, idx int) {
	o := ac.root
	for _, b := range s {
		b = ac.ord(b)
		if o.son[b] == nil {
			newNode := &acamNode{nodeId: ac.nodeCnt}
			o.son[b] = newNode
			ac.inDeg[newNode] = 0
			ac.nodeCnt++
		}
		o = o.son[b]
	}
	o.indices = append(o.indices, idx)
}

func (ac *acam) buildFail() {
	ac.root.fail = ac.root
	ac.root.last = ac.root
	q := make([]*acamNode, 0, ac.nodeCnt)
	ac.g = make([][]int, ac.nodeCnt)
	for i, son := range ac.root.son[:] {
		if son == nil {
			ac.root.son[i] = ac.root
		} else {
			son.fail = ac.root
			son.last = ac.root
			ac.g[0] = append(ac.g[0], son.nodeId)
			ac.inDeg[ac.root]++
			q = append(q, son)
		}
	}
	for len(q) > 0 {
		o := q[0]
		q = q[1:]
		for i, son := range o.son[:] {
			if son == nil {
				o.son[i] = o.fail.son[i]
				continue
			}
			son.fail = o.fail.son[i]
			ac.g[son.fail.nodeId] = append(ac.g[son.fail.nodeId], son.nodeId)
			ac.inDeg[son.fail]++
			if len(son.fail.indices) > 0 {
				son.last = son.fail
			} else {
				son.last = son.fail.last
			}
			q = append(q, son)
		}
	}
}

func (ac *acam) acCnt(text string) []int {
	o := ac.root
	for _, b := range text {
		o = o.son[ac.ord(b)]
		o.cnt++
	}
	cnt := make([]int, len(ac.patterns))
	inDeg := ac.inDeg
	q := make([]*acamNode, 0, ac.nodeCnt)
	for o, d := range inDeg {
		if d == 0 {
			q = append(q, o)
		}
	}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		for _, i := range v.indices {
			cnt[i-1] = v.cnt
		}
		u := v.fail
		u.cnt += v.cnt
		if inDeg[u]--; inDeg[u] == 0 {
			q = append(q, u)
		}
	}
	return cnt
}