package main

func solve() {
	n := II()
	S := IS()
	s := []byte(S)

	judge := make([]int, n / 2)
	sum := 0
	for i := range judge {
		if s[i] == s[n - 1 - i] {
			judge[i] = 1
			sum++
		}
	}

	out := func(sum int) {
		if sum == n / 2 {
			Println("YES")
		} else {
			Println("NO")
		}
	}

	q := II()
	for Q := 0; Q < q; Q++ {
		idx, char := II() - 1, IB()
		if n & 1 == 1 && idx == n / 2 {
			out(sum)
			continue
		}
		s[idx] = char
		if idx >= n / 2 {
			idx = n - 1 - idx
		}
		if s[idx] == s[n - 1 - idx] && judge[idx] == 0 {
			judge[idx] = 1
			sum++
		} else if s[idx] != s[n - 1 - idx] && judge[idx] == 1 {
			judge[idx] = 0
			sum--
		}
		out(sum)
	}
}
