#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
const int inf = INT_MAX;
const ll ll_inf = LLONG_MAX;

// 返回图的二染色
// 如果是二分图，返回每个节点的颜色，用 1 和 2 表示两种颜色
// 如果不是二分图，返回空列表
// 时间复杂度 O(n+m)，n 是点数，m 是边数
vector<int8_t> colorBipartite(int n, vector<vector<int>>& edges) {
    // 建图（节点编号从 0 到 n-1）
    vector<vector<int>> g(n);
    for (auto& e : edges) {
        int x = e[0], y = e[1];
        g[x].push_back(y);
        g[y].push_back(x);
    }

    // colors[i] = 0 表示未访问节点 i
    // colors[i] = 1 表示节点 i 为红色
    // colors[i] = 2 表示节点 i 为蓝色
    vector<int8_t> colors(n);

    // lambda 递归
    auto dfs = [&](this auto&& dfs, int x, int8_t c) -> bool {
        colors[x] = c; // 节点 x 染成颜色 c
        for (int y : g[x]) {
            // 邻居 y 的颜色与 x 的相同，说明不是二分图，返回 false
            // 或者继续递归，发现不是二分图，返回 false
            if (colors[y] == c || 
                colors[y] == 0 && !dfs(y, 3 - c)) { // 1 和 2 交替染色
                return false;
            }
        }
        return true;
    };

    // 可能有多个连通块
    for (int i = 0; i < n; i++) {
        if (colors[i] == 0 && !dfs(i, 1)) {
            // 从节点 i 开始递归，发现 i 所在连通块不是二分图
            return {};
        }
    }
    return colors;
}