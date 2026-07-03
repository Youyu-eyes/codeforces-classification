package main

import (
	"math"
)

func solve() {
	n := II()
	a := make([]int, n)
	for i := range a { a[i] = II() }

	L := make([]int, n)
	cnt := make([]int, n + 1)
	left, diff := 0, 0
	for i, x := range a {
		if cnt[x] == 0 {
			diff++
		}
		cnt[x]++
		for diff < i - left + 1 {
			cnt[a[left]]--
			if cnt[a[left]] == 0 {
				diff--
			}
			left++
		}
		L[i] = left
	}
	
	D := make([]int, n)
	for i := range D {
		D[i] = i
	}
	B := int(math.Sqrt(float64(n)))
	//B := int(math.Sqrt(float64(n * bits.Len(uint(n)))))

	type block struct {
		l, r int // [l, r)
		lazy int
		min  int
		cnt int
	}
	blocks := make([]block, (n-1)/B+1)

	for i := 0; i < n; i += B {
		r := min(i+B, n)
		blocks[i/B] = block{i, r, 0, i, 1}
	}

	// [l, r), 从 0 开始
	add := func(l, r int, v int) {
		for i := range blocks {
			b := &blocks[i]
			if b.l >= r {
				break
			}
			if b.r <= l {
				continue
			}
			if l <= b.l && b.r <= r {
				// 完整块
				b.lazy += v
			} else {
				// 部分块
				bl := max(b.l, l)
				br := min(b.r, r)
				for j := bl; j < br; j++ {
					D[j] += v
				}
				b.min = inf
				for j := b.l; j < b.r; j++ {
					if D[j] < b.min {
						b.min = D[j]
						b.cnt = 1
					} else if D[j] == b.min {
						b.cnt++
					}
				}
			}
		}
	}

	// [l, r), 从 0 开始
	query := func(l, r int) int {
		res := 0
		for i := range blocks {
			b := &blocks[i]
			if b.l >= r {
				break
			}
			if b.r <= l {
				continue
			}
			if l <= b.l && b.r <= r {
				// 完整块
				if b.min + b.lazy == r {
					res += b.cnt
				}
			} else {
				// 部分块
				bl := max(b.l, l)
				br := min(b.r, r)
				for j := bl; j < br; j++ {
					if D[j] + b.lazy == r {
						res++
					}
				}
			}
		}
		return res
	}

	st := []int{}
	ans := 0
	for i, x := range a {
		add(i, i + 1, x)
		for len(st) > 0 && x >= a[st[len(st) - 1]] {
			top := st[len(st) - 1]
			st = st[:len(st) - 1]
			last := -1
			if len(st) > 0 {
				last = st[len(st) - 1]
			}
			add(last + 1, top + 1, x - a[top])
		}
		st = append(st, i)
		ans += query(L[i], i + 1)
	}
	Println(ans)
}
