package main

func solve() {
	n := II()
	w := make([]int, n)
	for i := range w {
		w[i] = II()
	}
	g := make([][]int, n)
	for i := 0; i < n - 1; i++ {
		u, v := II() - 1, II() - 1
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	f0 := make([]int, n)
	f1 := make([]int, n)

	var dfs func(int, int)
	dfs = func(x, fa int) {
		if len(g[x]) < 1 {
			f0[x] = inf
			f1[x] = w[x]
			return
		}

		for _, y := range g[x] {
			if y != fa {
				dfs(y, x)
			}
		}

		diff := inf
		allWhite := true
		for _, y := range g[x] {
			if y != fa {
				diff = min(diff, abs(f0[y] - f1[y]))
				if f1[y] < f0[y] {
					allWhite = false
					f1[x] += f1[y]
				} else {
					f1[x] += f0[y]
				}
			}
		}
		f0[x] = f1[x]
		f1[x] += w[x]

		if allWhite {
			f0[x] += diff
		}
	}

	dfs(0, -1)
	Println(min(f0[0], f1[0]))
}