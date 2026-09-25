// https://codeforces.com/problemsets/acmsguru/problem/99999/512
// leetcode 3027 数据加强版，将 y 轴取反再算一遍，然后去重

// 题解
// https://leetcode.cn/problems/find-the-number-of-ways-to-place-people-ii/solutions/4034477/xiao-bai-si-lu-tu-jie-cdqfen-zhi-xian-du-ymad/

package main

import (
    "slices"
    "sort"
    "cmp"
)

func solve() {
    n := II()
    X := make([]int, n)
    Y := make([]int, n)
    for i := range n {
        X[i] = II()
        Y[i] = II()
    }
    rx, _ := discretize(X)
    ry, _ := discretize(Y)

	points := make([][]int, n)
    for i := range points {
		points[i] = make([]int, 2)
        points[i][0] = rx[i]
        points[i][1] = ry[i]
    }

	ans := 0

	ans += numberOfPairs(points)

	for i, p := range points {
		points[i][1] = -p[1]
	}

	ans += numberOfPairs(points)

	del := func(k int) {
		slices.SortFunc(points, func(a, b []int) int {
			return cmp.Compare(a[k], b[k])
		})

		for i := range n - 1 {
			if points[i][k] == points[i + 1][k] {
				ans--
			}
		}
	}

	del(0)
	del(1)

	Println(ans)
}

func numberOfPairs(points [][]int) (ans int) {
	n := len(points)

    slices.SortFunc(points, func(a, b []int) int {
		if a[0] != b[0] {
			return cmp.Compare(a[0], b[0])
		}
		return cmp.Compare(b[1], a[1])
	})

    var st = [][]int{}
    update := func(np []int) {
        for len(st) > 0 && st[len(st) - 1][0] <= np[0] {
            st = st[:len(st) - 1]
        }
        st = append(st, np)
    }

    lowerBound := func(y int) int {
        left, right := -1, len(st)
        for left + 1 < right {
            mid := left + (right - left) >> 1
            if st[mid][1] < y {
                right = mid
            } else {
                left = mid
            }
        }
        return right
    }

    bit := newFenwickTree(n)
    tmp := make([][]int, n)
    var cdq func(int, int)
    cdq = func(l, r int) {
        if l == r {
            return
        }
        mid := l + (r - l) >> 1
        cdq(l, mid)
        cdq(mid + 1, r)

        st = [][]int{}
        j, k := l, l
        for i := mid + 1; i <= r; i++ {
            for ; j <= mid && points[j][1] >= points[i][1]; j++ {
                update(points[j])
                tmp[k] = points[j]
                k++
            }

            minY := bit.pre(points[i][0])
            idx := lowerBound(minY)
            ans += len(st) - idx
            
            bit.update(points[i][0], points[i][1])
            tmp[k] = points[i]
            k++
        }

        for i := mid + 1; i <= r; i++ {
            bit.update(points[i][0], inf)
        }

        for ; j <= mid; j++ {
            tmp[k] = points[j]
            k++
        }

        copy(points[l:r+1], tmp[l:r+1])
    }
    cdq(0, n - 1)
    return
}

func discretize(arr []int) ([]int, int) {
    unique := slices.Clone(arr)
	slices.Sort(unique)
	unique = slices.Compact(unique)

    m := len(unique)
    rank := make([]int, len(arr))
	for i, x := range arr {
		rank[i] = sort.SearchInts(unique, x)
	}

	return rank, m
}

type fenwick []int

func newFenwickTree(n int) fenwick {
	f := make(fenwick, n+1) // 使用下标 1 到 n
    for i := range n + 1 {
        f[i] = inf
    }
    return f
}

// a[i] 增加 val
// 时间复杂度 O(log n)
func (f fenwick) update(i, val int) {
	for i++; i < len(f); i += i & -i {
		f[i] = val
	}
}

// 求前缀和 a[1] + ... + a[i]
// 时间复杂度 O(log n)
func (f fenwick) pre(i int) int {
    res := inf
	for i++; i > 0; i &= i - 1 {
		res = min(res, f[i])
	}
	return res
}
