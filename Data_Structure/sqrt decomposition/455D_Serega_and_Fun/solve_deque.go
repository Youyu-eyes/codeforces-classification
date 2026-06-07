// 2. 双端队列
// 对于本题而言，左端点插入一个数，右端点删除一个数，可适用双端队列
// 完整的块，左端入队，右端出队，都是 O(1) 时间
// 左端点，在 b.q 的 l - b.l 插入一个数，最多 O(B) 时间
// 右端点，在 b.q 的 r - b.l 删除一个数，最多 O(B) 时间

// 总体时间和空间复杂度与数组写法相同

package main

func solve_deque() {
	n := II()
	a := make([]int, n)
	for i := range a { a[i] = II() }

	B := 1000

	type block struct {
		l, r int // [l, r)
		cnt  []int
		q    deque[int]
	}

	blocks := make([]block, (n-1)/B+1)
	calcBlock := func(l, r int) ([]int, deque[int]) {
		cnt := make([]int, n + 1)
		q := deque[int]{}
		for j := l; j < r; j++ {
			v := a[j]
			cnt[v]++
			q.pushBack(v)
		}
		return cnt, q
	}
	for i := 0; i < n; i += B {
		r := min(i+B, n)
		cnt, q := calcBlock(i, r)
		blocks[i/B] = block{i, r, cnt, q}
	}

	// [l, r), 从 0 开始
	sqrtUpdate := func(l, r int) {
		last := &blocks[(r - 1) / B]
		cur := last.q.get(r - 1 - last.l)

		if l / B == (r - 1) / B {
			last.q.erase(r - 1 - last.l)
			last.q.insert(cur, l - last.l)
			return
		}

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
				tmp := b.q.back()
				b.q.popBack()
				b.cnt[tmp]--
				b.q.pushFront(cur)
				b.cnt[cur]++
				cur = tmp
			} else if b.l < l {
				tmp := b.q.back()
				b.q.popBack()
				b.cnt[tmp]--
				b.q.insert(cur, l - b.l)
				b.cnt[cur]++
				cur = tmp
			} else {
				tmp := b.q.get(r - 1 - b.l)
				b.q.erase(r - 1 - b.l)
				b.cnt[tmp]--
				b.q.pushFront(cur)
				b.cnt[cur]++
				return
			}
		}
	}

	// [l, r), 从 0 开始
	sqrtQuery := func(l, r int, v int) (sum int) {
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
				sum += b.cnt[v]
			} else {
				// 部分块
				bl := max(b.l, l)
				br := min(b.r, r)
				for j := bl; j < br; j++ {
					if b.q.get(j - b.l) == v {
						sum++
					}
				}
			}
		}
		return
	}

	lastans := 0
	q := II()
	for Q := 0; Q < q; Q++ {
		queryNumber := II()
		if queryNumber == 1 {
			lPrime, rPrime := II(), II()

			l := ((lPrime + lastans - 1) % n) + 1
			r := ((rPrime + lastans - 1) % n) + 1

			if l > r {
				l, r = r, l
			}

			sqrtUpdate(l-1, r)
		} else {
			lPrime, rPrime, kPrime := II(), II(), II()

			l := ((lPrime + lastans - 1) % n) + 1
			r := ((rPrime + lastans - 1) % n) + 1
			if l > r {
				l, r = r, l
			}
			k := ((kPrime + lastans - 1) % n) + 1

			ans := sqrtQuery(l-1, r, k)
			Println(ans)

			lastans = ans
		}
	}
}

type deque[T any] struct{ l, r []T }

func (q deque[T]) empty() bool {
	return len(q.l) == 0 && len(q.r) == 0
}

func (q deque[T]) size() int {
	return len(q.l) + len(q.r)
}

func (q *deque[T]) pushFront(v T) {
	q.l = append(q.l, v)
}

func (q *deque[T]) pushBack(v T) {
	q.r = append(q.r, v)
}

func (q *deque[T]) popFront() (v T) {
	if len(q.l) > 0 {
		q.l, v = q.l[:len(q.l)-1], q.l[len(q.l)-1]
	} else {
		v, q.r = q.r[0], q.r[1:]
	}
	return
}

func (q *deque[T]) popBack() (v T) {
	if len(q.r) > 0 {
		q.r, v = q.r[:len(q.r)-1], q.r[len(q.r)-1]
	} else {
		v, q.l = q.l[0], q.l[1:]
	}
	return
}

func (q deque[T]) front() T {
	if len(q.l) > 0 {
		return q.l[len(q.l)-1]
	}
	return q.r[0]
}

func (q deque[T]) back() T {
	if len(q.r) > 0 {
		return q.r[len(q.r)-1]
	}
	return q.l[0]
}

// 0 <= i < q.size()
func (q deque[T]) get(i int) T {
	if i < len(q.l) {
		return q.l[len(q.l)-1-i]
	}
	return q.r[i-len(q.l)]
}

// insert 在逻辑索引 idx 处插入 val (0 <= idx <= q.size())
func (q *deque[T]) insert(val T, idx int) {
	if idx == 0 {
		q.pushFront(val)
		return
	}
	if idx == q.size() {
		q.pushBack(val)
		return
	}

	if idx <= len(q.l) {
		pos := len(q.l) - idx
		var zero T
		q.l = append(q.l, zero)
		copy(q.l[pos+1:], q.l[pos:len(q.l)-1])
		q.l[pos] = val
	} else {
		ridx := idx - len(q.l)
		if ridx == 0 {
			q.r = append([]T{val}, q.r...)
		} else {
			var zero T
			q.r = append(q.r, zero)
			copy(q.r[ridx+1:], q.r[ridx:len(q.r)-1])
			q.r[ridx] = val
		}
	}
}

// erase 删除逻辑索引 idx 处的元素 (0 <= idx < q.size())
func (q *deque[T]) erase(idx int) {
	if idx < len(q.l) {
		pos := len(q.l) - 1 - idx
		if pos == len(q.l)-1 {
			q.l = q.l[:len(q.l)-1]
		} else {
			copy(q.l[pos:], q.l[pos+1:])
			q.l = q.l[:len(q.l)-1]
		}
	} else {
		ridx := idx - len(q.l)
		if ridx == 0 {
			q.r = q.r[1:]
		} else {
			copy(q.r[ridx:], q.r[ridx+1:])
			q.r = q.r[:len(q.r)-1]
		}
	}
}
