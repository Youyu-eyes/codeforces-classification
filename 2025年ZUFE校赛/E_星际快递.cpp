#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
const int inf = 0x3f3f3f3f;
const ll ll_inf = 1e18;

vector<long long> shortestPathDijkstra(int n, vector<vector<pair<int, int>>>& g, int start) {
    vector<long long> dis(n, ll_inf);
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


void solve() {
    int n, m; cin >> n >> m;
    int start, end, k; cin >> start >> end >> k;
    start -= 1;
    end -= 1;
    vector<int> target(k);
    for (int i = 0; i < k; ++i) {
        cin >> target[i];
        target[i] -= 1;
    }
    vector<vector<pair<int, int>>> g1(n);
    vector<vector<pair<int, int>>> g2(n);
    for (int i = 0; i < m; ++i) {
        int u, v, wt; cin >> u >> v >> wt;
        g1[u - 1].push_back({v - 1, wt});
        g2[v - 1].push_back({u - 1, wt});
    }
    vector<ll> come = shortestPathDijkstra(n, g1, start);
    vector<ll> back = shortestPathDijkstra(n, g2, end);
    ll ans = ll_inf;
    for (int t : target) {
        ans = min(ans, come[t] + back[t]);
    }
    cout << (ans < ll_inf ? ans : -1) << endl;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);
    
    solve();
    
    return 0;
}