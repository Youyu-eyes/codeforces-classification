// https://codeforces.com/contest/1791/problem/F
// 1500
// 分块

// 能分块的核心是，证明这个题最多需要 3 次操作就能将一个数变成个位数，变成个位数后，这个数就不需要操作了

// 先证明两位数最多需要两次操作

// 我们希望两位数多操作一次，需要 x = 10*a + b  =>  a + b > 9
// 则最小的 x 为 19，穷举法可以证明 < 19 的数都可以一次操作变成个位数
// 但是，二位数最大为 99，x 最大为 18，因此从 99 到 18 到 9 需要两步
// 因此，二位数到个位数最多需要两步操作

// 再证明，a[i] 到二位数最多需要 1 次操作

// a[i] < 10^9，最大 9 位数
// 9 * 9 = 81，9 位数相加最大 81，不可能出现更大的情况，因此最多需要 1 步就能变成二位数

// 综上，最多需要更新 3 次

// 我们将数组分块，块内维护参数 stable 表示块是否稳定
// 如果稳定就直接跳过，更新下一个块
// 不稳定则枚举块内元素，更新
// 更新完毕后，再遍历一遍块内元素，判断是否全部是个位数，如果是，则标记 stable 为 true，下一次直接跳过

// 枚举块需要 O(B) 时间，遍历每一个块需要 O(n / B) 时间
// 时间复杂度 O(n / B + B) B 取 sqrt(n) 时，复杂度最小

// 时间复杂度：O(n + q × sqrt(n))
// 空间复杂度：O(n)

package main

func digitSum(x int) (res int) {
	for x > 0 {
		res += x % 10
		x /= 10
	}
	return
}

func solve() {
	n, q := II(), II()
	
	a := make([]int, n)

	for i := range a {
		a[i] = II()
	}

	B := 500

	type block struct {
		l, r int // [l, r)
		stable bool
	}
	blocks := make([]block, (n-1)/B+1)

	for i := 0; i < n; i += B {
		r := min(i+B, n)
		blocks[i/B] = block{i, r, false}
	}

	update := func(ql, qr int) {
		for i := ql; i < qr; i++ {
			if a[i] >= 10 {
				a[i] = digitSum(a[i])
			}
		}
	}

	rejudge := func(b *block) {
		for i := b.l; i < b.r; i++ {
			if a[i] >= 10 {
				return
			}
		}
		b.stable = true
	}

	// [l, r), 从 0 开始
	sqrtUpdate := func(l, r int) {
		for i := range blocks {
			b := &blocks[i]

			if b.stable {
				continue
			}

			if b.l >= r {
				break
			}
			if b.r <= l {
				continue
			}
			if l <= b.l && b.r <= r {
				// 完整块
				
				update(b.l, b.r)
				rejudge(b)

			} else {
				// 部分块
				bl := max(b.l, l)
				br := min(b.r, r)
				update(bl, br)
				rejudge(b)
			}
		}
	}

	for Q := 0; Q < q; Q++ {
		queryNumber := II()
		if queryNumber == 1 {
			l, r := II() - 1, II()
			sqrtUpdate(l, r)
		} else {
			x := II() - 1
			Println(a[x])
		}
	}
}
