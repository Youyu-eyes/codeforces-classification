// https://codeforces.com/problemset/problem/1829/F
// 1300
// 图论  BFS 思想  数学
// 找到 x, y, m, n 之间的关系
// 一共分为 3 种节点，中心节点，第二层节点，第三层节点
// 观察发现，第三层节点度永远为 1
// 第二层节点度为 y + 1
// 中心节点度为 x
// 而 x > 1, y > 1
// 因此度为 1 的第三层节点具有特殊性，即度为 1 的节点一定是第三层节点
// 第三层节点唯一一个连接的就是第二层节点
// 找到第二层节点就能计算出节点的度，也就求出了 y
// 而 x, y, m, n 之间有数学关系，x + x * y = m
// 可以直接求出 x
// 本题实际上不需要 BFS 整张图，但是用到了 BFS 的思想

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;

void solve() {
    int n, m;
    cin >> n >> m;
    vector<vector<int>> g(n);
    for (int i = 0; i < m; ++i) {
        int u, v;
        cin >> u >> v;
        g[u - 1].push_back(v - 1);
        g[v - 1].push_back(u - 1);
    }

    int y;
    for (int i = 0; i < n; ++i) {
        if (g[i].size() == 1) {
            int j = g[i][0];
            y = g[j].size() - 1;
            break;
        }
    }
    int x = m / (y + 1);
    cout << x << ' ' << y << endl;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);
    
    int T;
    cin >> T;
    
    while (T--) {
        solve();
    }
    
    return 0;
}