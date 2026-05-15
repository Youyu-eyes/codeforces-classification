// https://acm.timus.ru/problem.aspx?space=1&num=1160

// 最小生成树
// 假定最小生成树的边权为 mst 数组，假设最大边权为 maxwt

// 反证法
// 假设有另一条边更小 maxwt1
// maxwt1 < maxwt
// 如果用 maxwt1 替换 mst 中的 maxwt1
// 此时，sum(mst1) < sum(mst)
// 但是 mst 数组的定义是最小生成树，sum(mst) 最小，与定义不符

// 如果用 maxwt1 替换 mst 中的另一条边
// 由于 maxwt1 < maxwt，此时生成树的最大边权仍旧是 maxwt 而非 maxwt1

// 综上，最小生成树的最大边权是所有生成树最大边权的最小值

// 以上可以当结论记住

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

const (
	MOD = 1_000_000_007
	inf = 0x3f3f3f3f
)

var (
	in  *bufio.Reader
	out *bufio.Writer
)

// n := II()
func II() (x int) {
	fmt.Fscan(in, &x)
	return
}

// n := IF()
func IF() (x float64) {
	fmt.Fscan(in, &x)
	return
}

// s := IS()
func IS() (x string) {
	fmt.Fscan(in, &x)
	return
}

func Print(a ...any)            { fmt.Fprint(out, a...) }
func Println(a ...any)          { fmt.Fprintln(out, a...) }
func Printf(f string, a ...any) { fmt.Fprintf(out, f, a...) }

type edge struct {
	u, v int
}

func solve() {
	n := II()
	m := II()
	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		u, v, w := II(), II(), II()
		edges[i] = []int{u - 1, v - 1, w}
	}
    maxWt, MST := mstKruskal(n, edges)
	Println(maxWt)
	Println(n - 1)
	for i := 0; i < n - 1; i++ {
		Println(MST[i].u + 1, MST[i].v + 1)
	}
}


func main() {
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	T := 1

	for t := 0; t < T; t++ {
		solve()
	}
}


func abs[T int | int64 | float64](x T) T {
	if x >= 0 { return x }
	return -x
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// 数组要手动切片 sum(nums[:])
func sum[T int | int64 | float64](nums []T) T {
	var total T
	for _, v := range nums {
		total += v
	}
	return total
}

func mstKruskal(n int, edges [][]int) (int, []edge) {
	slices.SortFunc(edges, func(a, b []int) int { return a[2] - b[2] })

	
	uf := newUnionFind(n)
	maxWt := 0
	tree := make([]edge, 0)
	for _, e := range edges {
		x, y, wt := e[0], e[1], e[2]
		if uf.merge(x, y) {
			maxWt = wt
			tree = append(tree, edge{x, y})
			if len(tree) == n - 1 { break }
		}
	}

	if uf.cc > 1 {
		return inf, make([]edge, 0)
	}
	return maxWt, tree
}

// ------- 并查集 ------- //
// 初始化：uf := newUnionFind(n)

type unionFind struct {
	fa []int // 代表元
	sz []int // 集合大小
	cc int   // 连通块个数
}

func newUnionFind(n int) unionFind {
	fa := make([]int, n)
	sz := make([]int, n)
	// 一开始有 n 个集合 {0}, {1}, ..., {n-1}
	// 集合 i 的代表元是自己，大小为 1
	for i := range fa {
		fa[i] = i
		sz[i] = 1
	}
	return unionFind{fa, sz, n}
}

// 返回 x 所在集合的代表元
// 同时做路径压缩，也就是把 x 所在集合中的所有元素的 fa 都改成代表元
func (u unionFind) find(x int) int {
	// 如果 fa[x] == x，则表示 x 是代表元
	if u.fa[x] != x {
		u.fa[x] = u.find(u.fa[x]) // fa 改成代表元
	}
	return u.fa[x]
}

// 判断 x 和 y 是否在同一个集合
func (u unionFind) same(x, y int) bool {
	// 如果 x 的代表元和 y 的代表元相同，那么 x 和 y 就在同一个集合
	// 这就是代表元的作用：用来快速判断两个元素是否在同一个集合
	return u.find(x) == u.find(y)
}

// 把 from 所在集合合并到 to 所在集合中
// 返回是否合并成功
func (u *unionFind) merge(from, to int) bool {
	x, y := u.find(from), u.find(to)
	if x == y { // from 和 to 在同一个集合，不做合并
		return false
	}

	// 按秩合并
	if u.sz[x] > u.sz[y] {
		x, y = y, x
	}

	u.fa[x] = y        // 合并集合。修改后就可以认为 from 和 to 在同一个集合了
	u.sz[y] += u.sz[x] // 更新集合大小（注意集合大小保存在代表元上）
	// 无需更新 sz[x]，因为我们不用 sz[x] 而是用 sz[find(x)] 获取集合大小，但 find(x) == y，我们不会再访问 sz[x]
	u.cc-- // 成功合并，连通块个数减一
	return true
}

// 返回 x 所在集合的大小
func (u unionFind) size(x int) int {
	return u.sz[u.find(x)] // 集合大小保存在代表元上
}