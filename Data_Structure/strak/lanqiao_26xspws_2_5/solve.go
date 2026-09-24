package main

func solve() {
	n := II()
	nums := make([]int, n)
	for i := range n {
		nums[i] = II()
	}

	cnt := make([]int, n)
	st := []int{}
	ans := n * (n + 1) / 2
	for i, x := range nums {
		tmp := []int{}
		for len(st) > 0 && nums[st[len(st) - 1]] <= x {
			j := st[len(st) - 1]
			st = st[:len(st) - 1]

			if cnt[j] == 0 {
				cnt[j] = 1
				tmp = append(tmp, j)
			} else {
				ans -= n - i
			}
		}
		st = append(st, i)

		for t := len(tmp) - 1; t >= 0; t-- {
			st = append(st, tmp[t])
		}
	}
	Println(ans)
}
