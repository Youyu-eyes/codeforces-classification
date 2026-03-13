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