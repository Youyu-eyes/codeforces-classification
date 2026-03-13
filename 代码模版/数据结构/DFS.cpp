#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
const int inf = INT_MAX;
const ll ll_inf = LLONG_MAX;

vector<int> solve(int n, vector<vector<int>>& edges) {
    // 节点编号从 0 到 n-1
    vector<vector<int>> g(n);
    for (auto& e : edges) {
        int x = e[0], y = e[1];
        g[x].push_back(y);
        g[y].push_back(x); // 无向图
    }

    vector<int8_t> vis(n);

    // lambda 递归
    auto dfs = [&](this auto&& dfs, int x) -> int {
        vis[x] = true; // 避免重复访问节点
        int size = 1;
        for (int y : g[x]) {
            if (!vis[y]) {
                size += dfs(y);
            }
        }
        return size;
    };

    // 计算每个连通块的大小
    vector<int> ans;
    for (int i = 0; i < n; i++) {
        if (!vis[i]) { // i 没有访问过
            int size = dfs(i);
            ans.push_back(size);
        }
    }
    return ans;
}