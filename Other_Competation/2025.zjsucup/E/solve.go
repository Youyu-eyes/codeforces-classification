package main

// leetcode 84. 柱状图中最大的矩形
func largestRectangleArea(heights []int) (ans int) {
    left := leftGreater(heights)
    right := rightGreater(heights)
    for i, h := range heights {
        l := left[i] + 1
        r := right[i] - 1
        ans = max(ans, (r - l + 1) * h)
    }
    return
}

// 返回 left，其中 left[i] 是 nums[i] 左侧最近的严格大于 nums[i] 的数的下标，若不存在则为 -1
func leftGreater(nums []int) []int {
	n := len(nums)
	left := make([]int, n)
	for i := range left { left[i] = -1 }
	st := []int{-1} // 哨兵
	for i, x := range nums {
		for len(st) > 1 && nums[st[len(st)-1]] >= x {
			st = st[:len(st)-1]
		}
		left[i] = st[len(st)-1]
		st = append(st, i)
	}
	return left
}

// 返回 right，其中 right[i] 是 nums[i] 右侧最近的严格大于 nums[i] 的数的下标，若不存在则为 len(nums)
func rightGreater(nums []int) []int {
	n := len(nums)
	right := make([]int, n)
	for i := range right { right[i] = n }
	st := []int{n} // 哨兵
	for i := n - 1; i >= 0; i-- {
		x := nums[i]
		for len(st) > 1 && nums[st[len(st)-1]] >= x {
			st = st[:len(st)-1]
		}
		right[i] = st[len(st)-1]
		st = append(st, i)
	}
	return right
}

func solve() {
	m, n := II(), II()
	grid := make([][]int, m)
	for i := range m {
		grid[i] = make([]int, n)
		for j := range n {
			x := II()
			if i & 1 == j & 1 {
				grid[i][j] = x ^ 1
			} else {
				grid[i][j] = x
			}
		}
	}

	h0 := make([]int, n)
	h1 := make([]int, n)
	ans := 0
	for i := range m {
		for j := range n {
			if grid[i][j] == 1 {
				h0[j] = 0
				h1[j]++
			} else {
				h0[j]++
				h1[j] = 0
			}
		}
		area0 := largestRectangleArea(h0)
		area1 := largestRectangleArea(h1)
		ans = max(ans, area0, area1)
	}

	Println(ans)
}