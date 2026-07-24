package main

import (
	"sort"
)

var pow2 = make([]int, 61)

func init() {
	pow2[0] = 1
	for i := 1; i <= 60; i++ {
		pow2[i] = pow2[i - 1] * 2
	}
}

func solve() {
	n := II()
	ans := sort.SearchInts(pow2, n + 1) - 1
	n -= pow2[ans]
	ans++
	for n > 0 {
		idx := sort.SearchInts(pow2, n + 1) - 1
		n -= pow2[idx]
		ans++
	}
	Println(ans)
}