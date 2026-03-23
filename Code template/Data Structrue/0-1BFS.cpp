#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
const int inf = INT_MAX;
const ll ll_inf = LLONG_MAX;

// 计算从 start 到各个节点的最短路长度
// 如果节点不可达，则最短路长度为 -1
// 节点编号从 0 到 n-1，边权均为 1
vector<int> bfs(int n, vector<vector<int>>& edges, int start) {
    vector<vector<int>> g(n);
    for (auto& e : edges) {
        int x = e[0], y = e[1];
        g[x].push_back(y);
        g[y].push_back(x); // 无向图
    }

    vector<int> dis(n, -1); // -1 表示尚未访问到
    queue<int> q;
    dis[start] = 0;
    q.push(start);
    while (!q.empty()) {
        int x = q.front();
        q.pop();
        for (int y : g[x]) {
            if (dis[y] < 0) {
                dis[y] = dis[x] + 1;
                q.push(y);
            }
        }
    }
    return dis;
}