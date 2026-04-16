#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
const int inf = 0x3f3f3f3f;
const ll ll_inf = 1e18;

// 计算从 start 到各个节点的最短路长度
// 如果节点不可达，则最短路长度为 inf
// 节点编号从 0 到 n-1，边权为 0 or 1
vector<int> zero_one_bfs(int n, vector<vector<pair<int, int>>>& g, int start) {
    vector<int> dis(n, inf);
    deque<int> q;
    q.push_back(start);

    while (!q.empty()) {
        int x = q.front();
        q.pop_front();

        for (auto &[y, w] : g[x]) {
            if (dis[x] + w < dis[y]) {
                dis[y] = dis[x] + w;
                if (w == 0) {
                    q.push_front(y);
                } else {
                    q.push_back(y);
                }
            }
        }
    }
    return dis;
}