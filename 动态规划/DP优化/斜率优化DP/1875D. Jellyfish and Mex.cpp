// https://codeforces.com/contest/1875/problem/D
// 原题 1600
// 凸包/斜率优化DP
// 斜率优化DP 做法大概 1850

// 状态转移方程具体推导见 子序列DP

// 重点：斜率优化（凸包优化）DP
// 由 f[j] + cnt[i] * j 可知
// 写成点积形式，可以带来几何意义上的观察
// => 1 * f[j] + cnt[i] * j
// => dot((1, cnt[i]), (f[j], j))
// 转化出一个与 j 无关的向量
// cnt[i] 不具单调性，无法用单调队列优化，需要在凸包上二分（类似通过导数与 0 的大小，判断下降还是上升）
// 每次更新完 f[i] 后，用新的 f[i] 构造向量，然后用 类似单调栈（不是单调栈，因为不是根据单调栈来的）维护下凸包
// 由于没有单调性，需要在凸包上二分，复杂度为 O(mlogm)

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
ll inf = LLONG_MAX;

struct Vec {
    ll x, y;
    Vec(ll x = 0, ll y = 0) : x(x), y(y) {}
};

Vec operator-(const Vec& a, const Vec& b) {
    return Vec(a.x - b.x, a.y - b.y);
}

__int128 det(const Vec& a, const Vec& b) {
    return (__int128) a.x * b.y - (__int128) a.y * b.x;
}

ll dot(const Vec& a, const Vec& b) {
    return a.x * b.x + a.y * b.y;
}

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

    deque<Vec> q;  // 不知道为什么，双端队列比 vector 开销小，时间快一点
    q.push_back(Vec(f[m], m));
    for (ll i = m - 1; i >= 0; --i) {
        Vec v0(1, cnt[i]);

        // 二分找最小
        ll l = 0, r = q.size() - 1;
        while (l < r) {
            ll mid = (l + r) / 2;
            if (dot(v0, q[mid]) <= dot(v0, q[mid + 1])) {
                r = mid;
            } else {
                l = mid + 1;
            }
        }
        
        // 更新答案
        f[i] = dot(v0, q[l]);

        // 维护下凸包
        Vec v(f[i], i);
        while (q.size() > 1 && det(q[q.size() - 1] - q[q.size() - 2], v - q[q.size() - 1]) <= 0) {
            q.pop_back();
        }
        q.push_back(v);
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