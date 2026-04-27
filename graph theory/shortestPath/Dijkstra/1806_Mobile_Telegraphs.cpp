// https://acm.timus.ru/problem.aspx?space=1&num=1806
// Dijkstra 变形题

// 题目很容易想到，从一个起点走到一个终点，求最短路径，可以抽象成图论问题，用 Dijkstra 解决
// 边权为两个点的 LCP，最多可以 O(10) 地计算出来
// 比较麻烦的是建图，如果暴力遍历两个点判断是否满足条件，需要 O(n^2) 的时间，是不被接受的
// 更好的方法是从一个点出发，暴力遍历更改一个数，或者交换两个数，生成的新数再和所有点的集合判断
// 枚举 10 位，暴力枚举 [0, 9] 10 个数字，需要 100，交换两个数需要 C(10, 2) = 45，哈希集合查询 O(1)
// 枚举共需要 O(145 * n) 大约 10^7 数量级，可以接受
// 显式建图的话每个节点最多关联 135 个节点（枚举更改时，每一位数字有一位是自己，比如 1 改成 1，没有更改，但是图中每个点都不一样）
// 空间复杂度为 O(135 * n)
// 如果想节省空间可以在 Dijkstra 计算的中途枚举，就可以节省空间到 10^5 级别
// 本题解采用显式建图，这样可以直接调用模板，比较方便

// 建完图之后，跑一个 Dijkstra，顺便返回路径，再跑一个 dfs 输出路径即可

// 现在的问题变成了，如何建图，思路已经在上面了，代码要怎么实现
// 如果依赖字符串或者数组，显然不是我们的风格，纯数字的处理要转换成字符串还是太丑陋了

// 操作一：更改一个数

// 9123493342
//      i

// 假设我们需要更改 9 这个数字改成 d，那么 9 会将这个数字分为三部分
// 91234, d, 3342
// 更改后的数字 = 91234 * 10^5 + d * 10^4 + 3342
// 问题就变成了怎么提取前后两部分
// 前半段用整除，后半段用取模
// 最后的公式即为 u / pow10[i + 1] * pow10[i + 1] + d * pow10[i] + u % pow10[i], u 表示当前的点

// 操作二是将数分为 3 段，其余处理逻辑一样，请同学们自行画图

// 时间复杂度: O(k * n + n + m log m)，其中 m 表示 Dijkstra 堆中的元素数量，k = 145
// 时间复杂度: O((k - 10) * n + m)

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
const int inf = 0x3f3f3f3f;
const ll ll_inf = 1e18;

vector<ll> pow10(11);
auto init = [] {
    pow10[0] = 1;
    for (int i = 1; i <= 10; ++i) {
        pow10[i] = pow10[i - 1] * 10;
    }
    return 0;
}();

// dijkstra：单源最短路
// 返回从起点 start 到每个点的最短路长度 dis，如果节点 x 不可达，则 dis[x] = ll_inf
// 要求：没有负数边权
// 时间复杂度 O(n + mlogm)，注意堆中有 O(m) 个元素
pair<vector<long long>, vector<int>> shortestPathDijkstra(int n, vector<vector<pair<int, int>>>& g, int start) {
    vector<long long> dis(n, ll_inf);
    // 堆中保存 (起点到节点 x 的最短路长度，节点 x)
    priority_queue<pair<long long, int>, vector<pair<long long, int>>, greater<>> pq;
    dis[start] = 0; // 起点到自己的距离是 0
    pq.emplace(0, start);

    vector<int> pre(n, -1);

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

                // 找到更短的路径，重置前驱节点
                pre[y] = x;
            }
        }
    }

    return {dis, pre};
}

void solve() {
    int n; cin >> n;
    vector<int> weights(10);
    for (int i = 0; i < 10; ++i) {
        cin >> weights[i];
    }
    unordered_map<ll, int> p2i;
    vector<ll> points(n);
    for (int i = 0; i < n; ++i) {
        ll p; cin >> p;
        p2i[p] = i;
        points[i] = p;
    }

    vector<vector<pair<int, int>>> g(n);
    auto update = [&](int x, ll v, int i) -> void {
        if (!p2i.contains(v)) {
            return;
        }
        int y = p2i[v];
        ll wt = weights[9 - i];
        g[x].push_back({y, wt});
        g[y].push_back({x, wt});
    };

    for (int x = 0; x < n; ++x) {
        ll u = points[x];
        for (int i = 0; i < 10; ++i) {
            for (ll d = 0; d < 10; ++d) {
                ll v = u / pow10[i + 1] * pow10[i + 1] + d * pow10[i] + u % pow10[i];
                update(x, v, i);
            }
            for (int j = 0; j < i; ++j) {
                ll v = u / pow10[i + 1] * pow10[i + 1] + u / pow10[j] % 10 * pow10[i] + u % pow10[i] / pow10[j + 1] * pow10[j + 1] + u / pow10[i] % 10 * pow10[j] + u % pow10[j];
                update(x, v, i);
            }
        }
    }

    auto [dis, pre] = shortestPathDijkstra(n, g, 0);

    if (pre[n - 1] < 0) {
        cout << -1 << endl;
        return;
    }

    cout << dis[n - 1] << endl;

    vector<int> path;
    auto dfs = [&](this auto&& dfs, int x) -> void {
        if (x > 0) dfs(pre[x]);
        path.push_back(x);
    };
    dfs(n - 1);
    int ans = path.size();
    cout << ans << endl;
    for (int i = 0; i < ans; ++i) {
        cout << path[i] + 1 << ' ';
    }
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);
    
    solve();
    
    return 0;
}