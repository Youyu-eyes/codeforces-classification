// https://www.luogu.com.cn/problem/P3808
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
	Println(ac.acCountOnce(text))
}

type acamNode struct {
	son [acamNodeSize]*acamNode

	fail *acamNode
	last *acamNode

	idx int
	cnt int
}

type acam struct {
	patterns []string
	
	root *acamNode
	nodeCnt int
}

func newAcam(patterns []string) *acam {
	ac := &acam{
		patterns: patterns,
		root: &acamNode{},
		nodeCnt: 1,
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
			newNode := &acamNode{}
			o.son[b] = newNode
			ac.nodeCnt++
		}
		o = o.son[b]
	}
	o.cnt++
	o.idx = idx
}

func (ac *acam) buildFail() {
	ac.root.fail = ac.root
	ac.root.last = ac.root
	q := make([]*acamNode, 0, ac.nodeCnt)
	for i, son := range ac.root.son[:] {
		if son == nil {
			ac.root.son[i] = ac.root
		} else {
			son.fail = ac.root
			son.last = ac.root
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
			if son.fail.idx > 0 {
				son.last = son.fail
			} else {
				son.last = son.fail.last
			}
			q = append(q, son)
		}
	}
}

func (ac *acam) acCountOnce(text string) (ans int) {
	o := ac.root
	for _, b := range text {
		o = o.son[ac.ord(b)]
		for match := o; match.cnt >= 0; match = match.last {
			ans += match.cnt
			match.cnt = -1
		}
	}
	return
}