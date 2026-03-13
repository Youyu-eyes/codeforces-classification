#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
const int inf = INT_MAX;
const ll ll_inf = LLONG_MAX;

// 返回有向无环图（DAG）的其中一个拓扑序
// 如果图中有环，返回空列表
// 节点编号从 0 到 n-1
vector<int> topologicalSort(int n, vector<vector<int>>& edges) {
    vector<vector<int>> g(n);
    vector<int> in_deg(n);
    for (auto& e : edges) {
        int x = e[0], y = e[1];
        g[x].push_back(y);
        in_deg[y]++; // 统计 y 的先修课数量
    }

    queue<int> q;
    for (int i = 0; i < n; i++) {
        if (in_deg[i] == 0) { // 没有先修课，可以直接上
            q.push(i); // 加入学习队列
        }
    }

    vector<int> topo_order;
    while (!q.empty()) {
        int x = q.front();
        q.pop();
        topo_order.push_back(x);
        for (int y : g[x]) {
            in_deg[y]--; // 修完 x 后，y 的先修课数量减一
            if (in_deg[y] == 0) { // y 的先修课全部上完
                q.push(y); // 加入学习队列
            }
        }
    }

    if (topo_order.size() < n) { // 图中有环
        return {};
    }
    return topo_order;
}