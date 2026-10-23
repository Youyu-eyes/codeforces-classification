// https://atcoder.jp/contests/aalc001/tasks/aalc001_h
// 1700
// 树状数组二分

package main

func solve() {
	n, m := II(), II()
	nums := make([]int, n)
	bit := newFenwickTree(n)
	bit0 := newFenwickTree(n)
	for i := range n {
		x := II()
		nums[i] = x
		bit.update(i, x)
		if x == 0 {
			bit0.update(i, 1)
		}
	}
	
	for range m {
		op := II()
		if op == 1 {
			idx, val := II() - 1, II()
			bit.update(idx, val - nums[idx])
			if nums[idx] == 0 && val != 0 {
				bit0.update(idx, -1)
			} else if nums[idx] != 0 && val == 0 {
				bit0.update(idx, 1)
			}
			nums[idx] = val
			continue
		}
		l, r, x := II() - 1, II() - 1, II()
		zero := bit0.query(l, r)
		if zero == 0 {
			x += bit.query(l, r)
		} else {
			zero += bit0.pre(l - 1)
			k := bit0.kth(zero)
			x = (x + bit.query(l, k)) % MOD
			x += bit.query(k, r)
		}
		Println(x)
	}
}

type fenwick []int

func newFenwickTree(n int) fenwick {
	return make(fenwick, n+1)
}

func (f fenwick) update(i, val int) {
	for i++; i < len(f); i += i & -i {
		f[i] += val
	}
}

func (f fenwick) pre(i int) (res int) {
	for i++; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}

func (f fenwick) query(l, r int) int {
	return f.pre(r) - f.pre(l-1)
}

// 返回最小的 0-based 下标 i，满足 a[0] + ... + a[i] >= k
// 若总和 < k，返回 n
func (f fenwick) kth(k int) int {
	n := len(f) - 1
	bit := 1
	for bit<<1 <= n {
		bit <<= 1
	}
	idx := 0
	for ; bit > 0; bit >>= 1 {
		nxt := idx + bit
		if nxt <= n && f[nxt] < k {
			k -= f[nxt]
			idx = nxt
		}
	}
	return idx
}
