// https://codeforces.com/contest/1628/problem/A
// 1400
// 贪心 + 分组循环

// https://leetcode.cn/problems/lexicographically-maximum-mex-array/description/
// leetcode 504 周赛

package main

func solve() {
    n := II()
	nums := make([]int, n)
	for i := range nums { nums[i] = II() }

    freq := [MX]int{}
    for _, x := range nums {
        freq[x]++
    }
    
    mex := 0
    for x := 0; freq[x] > 0; x++ { mex = x + 1 }

    vis := make([]int, MX)
    tag := 0
    l := 0
    var tmp_mex, x, r, need int

    ans := []int{}
    for l < n {
        tag++
        tmp_mex = mex
        need = mex
        if mex == 0 {
            for i := l; i < n; i++ {
                ans = append(ans, 0)
            }
            break
        }
        ans = append(ans, mex)

        for r = l; r < n && need > 0; r++ {
            x = nums[r]
            if vis[x] != tag && x < tmp_mex {
                need--
                vis[x] = tag
            }

            freq[x]--
            if freq[x] == 0 && x < mex {
                mex = x
            }
        }
        l = r
    }
    Println(len(ans))
	for _, x := range ans {
		Print(" ")
		Print(x)
	}
	Println()
}
