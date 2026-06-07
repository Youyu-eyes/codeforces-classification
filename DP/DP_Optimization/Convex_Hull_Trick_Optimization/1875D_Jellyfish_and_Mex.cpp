// https://codeforces.com/contest/1875/problem/D
// 原题 1600
// 凸包/斜率优化DP
// 斜率优化DP 做法大概 1850

// 状态转移方程具体推导见 子序列DP

// 重点：斜率优化（凸包优化）DP
// 由 f[j] + cnt[i] * j 可知
// 写成点积形式，可以带来几何意义上的观察
// => cnt[i] * j + 1 * f[j]
// => dot((cnt[i], 1), (j, f[j]))
// 转化出一个与 j 无关的向量
// cnt[i] 不具单调性，无法用单调队列优化，需要在凸包上二分（类似通过导数与 0 的大小，判断下降还是上升）
// 每次更新完 f[i] 后，用新的 f[i] 构造向量，然后用 类似单调栈（不是单调栈，因为不是根据单调栈来的）维护下凸包

// 由于本题的 i 是倒序枚举的，所以待维护向量 v1.x 单调递减，则两边加负号变成单调递增
// dot((-cnt[i], 1), (-j, f[j]))
// 由于没有单调性，需要在凸包上二分，复杂度为 O(mlogm)

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
ll inf = LLONG_MAX;

struct Vec {
    ll x, y;
    Vec(ll x = 0, ll y = 0) : x(x), y(y) {}

    Vec operator-(const Vec& other) const {
        return Vec(x - other.x, y - other.y);
    }

    ll dot(const Vec& other) const {
        return x * other.x + y * other.y;
    }

    // a.det(b) > 0 => a 到 b 逆时针
    // a.det(b) < 0 => a 到 b 顺时针
    __int128 det(const Vec& other) const {
        return (__int128) x * other.y - (__int128) y * other.x;
    }
};

// 下凸包（求最小值）
struct LowerHull {
    deque<Vec> hull;
    void add(const Vec& p) {
        while (hull.size() > 1 && (hull.back() - hull[hull.size() - 2]).det(p - hull.back()) <= 0) {
            hull.pop_back();
        }

        hull.push_back(p);
    }

    // 保证 v0.x 单调
    // 如果 v0.x 单调递增，则 LowerHull.query_monotonic(v0,  1)
    // 如果 v0.x 单调递减，则 LowerHull.query_monotonic(v0, -1)
    // 复杂度 O(n)
    long long query_monotonic(const Vec& p, int dir) {
        if (dir < 0) {
            while (hull.size() > 1 && p.dot(hull[0]) >= p.dot(hull[1])) {
                hull.pop_front();
            }
            return p.dot(hull.front());
        } else {
            while (hull.size() > 1 && p.dot(hull.back()) >= p.dot(hull[hull.size() - 2])) {
                hull.pop_back();
            }
            return p.dot(hull.back());
        }
    }

    // 二分查询最小值，复杂度 O(nlogn)
    long long query_binary(const Vec& p) const {
        int l = 0, r = hull.size() - 1;
        while (l < r) {
            int mid = (l + r) >> 1;
            if (p.dot(hull[mid]) >= p.dot(hull[mid + 1])) l = mid + 1;
            else r = mid;
        }
        return p.dot(hull[l]);
    }
    bool empty() const { return hull.empty(); }
    void clear() { hull.clear(); }
};

// 斜率优化DP
void solve() {
    int n; cin >> n;
    unordered_map<ll, int> cnt;
    for (int i = 0; i < n; ++i) {
        ll x; cin >> x;
        cnt[x]++;
    }
    ll m = 0;
    while (cnt[m]) m++;
    vector<ll> f(m + 1, inf);
    f[m] = 0;

    LowerHull q;
    q.add(Vec(-m, f[m]));
    for (ll i = m - 1; i >= 0; --i) {
        Vec v0(-cnt[i], 1);
        
        // 更新答案
        f[i] = q.query_binary(v0);

        // 维护下凸包
        Vec v(-i, f[i]);
        
        q.add(v);
    }
    cout << f[0] - m << endl;
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