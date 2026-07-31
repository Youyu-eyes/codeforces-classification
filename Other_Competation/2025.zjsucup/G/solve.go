// 调和级数枚举

package main

func solve() {
	n := II()
	cnt := make([]int, MX + 1)
	for i := 0; i < n; i++ {
		x := II()
		cnt[x]++
	}

	mutiCnt := make([]int, MX + 1)
	for i := 1; i <= MX; i++ {
		for j := i; j <= MX; j += i {
			mutiCnt[i] += cnt[j]
		}
	}

	cur := 0
	ans := 0
	for i := MX; i >= 0; i-- {
		x := mutiCnt[i]
		if x > cur {
			ans += i * (x - cur)
			cur = x
		}
		if cur == n {
			break
		}
	}

	Println(ans)
}