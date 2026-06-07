// https://acm.timus.ru/problem.aspx?space=1&num=1162

// Bellman-Ford 算法
// 最长路 + 判断正权环

// Ford 算法证明
// 这里证明最长路，最短路是同理的
// 首先有定理：一个没有正权环的图，最长路一定是一条简单路径
// 如果有重复走过一个节点，说明经过了一个环，如果这个环的总边权是正的，那么最长路将是无穷大
// 反之，走过这个环回到一个节点 一定比 不走这个环的路程短
// Frod 算法的思想是，循环松弛 n - 1 次节点，一定能得到最长路（无正权环）

// 首先，从起点到起点，如果走回到起点就一定走一次环，根据上述定理，不走一定最佳
// 因此起点到起点的最长路就是 0
// 假设松弛 k 次能让最长路边数为 k 的节点的最长路被正确计算
// 那么松弛 k - 1 次能让 k - 1 的节点最长路被计算
// 假设节点 v 的最长路是 start -> a -> b -> …… -> u -> v，长度为 k
// 现在经过 k - 1 次松弛后，假设 start -> u 已经被正确计算
// 那么再经历一轮松弛，dist[v] = max(dist[v], dist[u] + w(u, v))
// dist[v] >= dist[u] + w(u, v)
// 且所有经过 k 条边到达 dist[v] 的路径都被确定，因此经过 k 次松弛后，dist[v] 就是最多经过 k 条边的最长路径

// 时间复杂度: O(n·m)
// 空间复杂度: O(n + m)

package main

import (
	"bufio"
	"fmt"
	"os"
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
	a, b int
	rab, cab float64
}

func solve() {
	n, m, s, v := II(), II(), II(), IF()
	edges := make([]edge, 0)
	for i := 0; i < m; i++ {
		a, b, rab, cab, rba, cba := II(), II(), IF(), IF(), IF(), IF()
		edges = append(edges, edge{a - 1, b - 1, rab, cab})
		edges = append(edges, edge{b - 1, a - 1, rba, cba})
	}
	_, has := BellmanFord(n, edges, s - 1, v)
	if has {
		Println("YES")
	} else {
		Println("NO")
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

func BellmanFord(n int, edges []edge, start int, start_val float64) ([]float64, bool) {
	dist := make([]float64, n)
	for i := range dist { dist[i] = 0.0 }
	dist[start] = start_val

	for i := 0; i < n - 1; i++ {
		updated := false
		for _, e := range edges {
			u, v, r, c := e.a, e.b, e.rab, e.cab
			if dist[u] > c && (dist[u] - c) * r > dist[v] {
				dist[v] = (dist[u] - c) * r
				updated = true
			}
		}
		if !updated {
			break
		}
	}

	hasCycle := false
	for _, e := range edges {
		u, v, r, c := e.a, e.b, e.rab, e.cab
		if dist[u] > c && (dist[u] - c) * r > dist[v] {
			hasCycle = true
			break
		}
	}

	return dist, hasCycle
}