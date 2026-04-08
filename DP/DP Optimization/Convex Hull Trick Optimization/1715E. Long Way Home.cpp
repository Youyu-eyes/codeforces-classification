// https://codeforces.com/problemset/problem/1715/E
// 2400
// 最短路 + 凸包优化DP

// 综合性很强的一题
// 为了方便与代码对应，本文的 u, v 默认从 0 开始，n - 1 结束
// 题目问最多 k 次飞行的最短路，如果没有飞行就是普通的 Dijkstra 问题
// 如果我们能计算单次飞行对 dis 数组的影响，问题就变成了最多 k - 1 次飞行的问题，这是一个规模更小的子问题
// 假设我们知道最多 k - 1 次飞行的 dis 数组，对于图上每一个点 v，我们都可以找到任意一个图上的点 u 飞到 v
// 那么到 v 的路径就被分成了两端：从 0 -> u，再从 u -> v
// 由于后一段是定值，要使总路程最小，则前一段最小，上一次的 dis[u] 计算的就是 0 -> u 的最多飞行 k - 1 次的最短路径
// 所以最后一次是飞到 v 的最短路径就是 min{u = [0, n - 1]}(dis[u] + (u - v)^2)

// 显然可以用凸包优化
// 将式子展开，写成点积形式，可以带来几何意义上的观察
// min{u = [0, n - 1]}(dis[u] + (u - v)^2) = min{u = [0, n - 1]}(dis[u] + u^2 - 2uv) + v^2
// 中间的方程写成点积形式：(-2v, 1) * (u, dis[u] + u^2)，其中 * 表示点积
// 这个转移是从整个 u = [0, n - 1] 转移的，而非 [0, v - 1] 的区间，所以我们需要预处理整个下凸包

// 现在我们得到了最后一次以飞行结尾的距离数组，我们需要以这个数组为源，再跑一遍 Dijkstra
// 为什么不能直接取 min
// 因为可能出现飞到 D 点再走到 E 点的情况，直接取 min 只会更新 D 点和 E 点，忽略上述情况
// 你也可以先取 min 再跑 Dijkstra，但是这样是多余的操作
// 你可以想象成从 0 到每个点 u 全部连一条 fly[u] 长度的路径，再用这张新的图求最短路

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
const int inf = 0x3f3f3f3f;
const ll ll_inf = 1e18;

struct Vec {
    ll x, y;

    Vec operator-(const Vec& b) {
        return Vec(x - b.x, y - b.y);
    }

    // a.det(b) > 0 => a 到 b 逆时针
    // a.det(b) < 0 => a 到 b 顺时针
    __int128 det(const Vec& b) {
        return (__int128) x * b.y - (__int128) y * b.x;
    }

    ll dot(const Vec& b) {
        return x * b.x + y * b.y;
    }

};

vector<long long> shortestPathDijkstra(int n, vector<long long> dis,
                                      vector<vector<pair<int, int>>>& g) {
    priority_queue<pair<long long, int>, vector<pair<long long, int>>, greater<>> pq;
    for (int i = 0; i < n; ++i) {
        if (dis[i] < ll_inf) {
            pq.emplace(dis[i], i);
        }
    }
    while (!pq.empty()) {
        auto [d, x] = pq.top(); pq.pop();
        if (d > dis[x]) continue;
        for (auto& [y, w] : g[x]) {
            if (dis[x] + w < dis[y]) {
                dis[y] = dis[x] + w;
                pq.emplace(dis[y], y);
            }
        }
    }
    return dis;
}

void solve() {
    int n, m, K; cin >> n >> m >> K;
    vector<vector<pair<int, int>>> g(n);
    for (int i = 0; i < m; ++i) {
        int u, v, w; cin >> u >> v >> w;
        u -= 1;
        v -= 1;
        g[u].push_back({v, w});
        g[v].push_back({u, w});
    }
    vector<ll> dis(n, ll_inf);
    dis[0] = 0;
    dis = shortestPathDijkstra(n, dis, g);
    for (int k = 1; k <= K; ++k) {
        vector<ll> fly = dis;
        deque<Vec> q;

        // 计算下凸包（从左到右）
        for (int v = 0; v < n; ++v) {
            Vec p(v, dis[v] + 1LL * v * v);
            while (q.size() > 1 && (p - q.back()).det(q.back() - q[q.size() - 2]) >= 0) {
                q.pop_back();
            }
            q.push_back(p);
        }

        for (int v = 0; v < n; ++v) {
            Vec v0(-2 * v, 1);
            while (q.size() > 1 && v0.dot(q[0]) >= v0.dot(q[1])) {
                q.pop_front();
            }
            fly[v] = v0.dot(q[0]) + 1LL * v * v;
        }
        dis = shortestPathDijkstra(n, fly, g);
    }
    for (int i = 0; i < n; ++i) {
        cout << dis[i] << ' ';
    }
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    solve();
    
    return 0;
}