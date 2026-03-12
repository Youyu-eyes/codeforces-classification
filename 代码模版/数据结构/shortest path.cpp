#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
const int inf = INT_MAX;
const ll ll_inf = LLONG_MAX;


// dijkstra：单源最短路
// 返回从起点 start 到每个点的最短路长度 dis，如果节点 x 不可达，则 dis[x] = LLONG_MAX
// 要求：没有负数边权
// 时间复杂度 O(n + mlogm)，注意堆中有 O(m) 个元素
vector<long long> shortestPathDijkstra(int n, vector<vector<int>>& edges, int start) {
    // 注：如果节点编号从 1 开始（而不是从 0 开始），可以把 n 加一
    vector<vector<pair<int, int>>> g(n); // 邻接表
    for (auto& e : edges) {
        int x = e[0], y = e[1], wt = e[2];
        g[x].emplace_back(y, wt);
        // g[y].emplace_back(x, wt); // 无向图加上这行
    }

    vector<long long> dis(n, LLONG_MAX);
    // 堆中保存 (起点到节点 x 的最短路长度，节点 x)
    priority_queue<pair<long long, int>, vector<pair<long long, int>>, greater<>> pq;
    dis[start] = 0; // 起点到自己的距离是 0
    pq.emplace(0, start);

    while (!pq.empty()) {
        auto [dis_x, x] = pq.top();
        pq.pop();
        if (dis_x > dis[x]) { // x 之前出堆过
            continue;
        }
        for (auto& [y, wt] : g[x]) {
            auto new_dis_y = dis_x + wt;
            if (new_dis_y < dis[y]) {
                dis[y] = new_dis_y; // 更新 x 的邻居的最短路
                // 懒更新堆：只插入数据，不更新堆中数据
                // 相同节点可能有多个不同的 new_dis_y，除了最小的 new_dis_y，其余值都会触发上面的 continue
                pq.emplace(new_dis_y, y);
            }
        }
    }

    return dis;
}


// Floyd：全源最短路

// 返回一个二维列表，其中 (i,j) 这一项表示从 i 到 j 的最短路长度
// 如果无法从 i 到 j，则最短路长度为 LLONG_MAX / 2
// 允许负数边权
// 如果计算完毕后，存在 i，使得从 i 到 i 的最短路长度小于 0，说明图中有负环
// 节点编号从 0 到 n-1
// 时间复杂度 O(n^3 + m)，其中 m 是 edges 的长度
vector<vector<long long>> shortestPathFloyd(int n, vector<vector<int>>& edges) {
    const long long INF = LLONG_MAX / 2; // 防止加法溢出
    vector f(n, vector<long long>(n, INF));
    for (int i = 0; i < n; i++) {
        f[i][i] = 0;
    }

    for (auto& e : edges) {
        int x = e[0], y = e[1];
        long long wt = e[2];
        f[x][y] = min(f[x][y], wt); // 如果有重边，取边权最小值
        f[y][x] = min(f[y][x], wt); // 无向图
    }

    for (int k = 0; k < n; k++) {
        for (int i = 0; i < n; i++) {
            if (f[i][k] == INF) { // 针对稀疏图的优化
                continue;
            }
            for (int j = 0; j < n; j++) {
                f[i][j] = min(f[i][j], f[i][k] + f[k][j]);
            }
        }
    }
    return f;
}