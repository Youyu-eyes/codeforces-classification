#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
ll inf = LLONG_MAX;

struct Vec {
    ll x, y;

    Vec operator-(const Vec& b) {
        return Vec(x - b.x, y - b.y);
    }

    __int128 det(const Vec& b) {
        return (__int128) x * b.y - (__int128) y * b.x;
    }

    ll dot(const Vec& b) {
        return x * b.x + y * b.y;
    }

};


// 向量 v0 单调
void solve1() {
    int n; cin >> n;
    vector<ll> a(n);
    vector<ll> b(n);
    for (int i = 0; i < n; ++i) {
        cin >> a[i];
    }
    for (int i = 0; i < n; ++i) {
        cin >> b[i];
    }
    
    vector<ll> f(n);
    deque<Vec> q;
    q.push_back(Vec(b[0], 0));
    for (int i = 1; i < n; ++i) {
        Vec v0(a[i], 1);
        while (q.size() > 1 && v0.dot(q[0]) >= v0.dot(q[1])) {
            q.pop_front();
        }

        f[i] = v0.dot(q[0]);
        Vec v1(b[i], f[i]);

        while (q.size() > 1 && (v1 - q[q.size() - 2]).det(q.back() - q[q.size() - 2]) <= 0) {
            q.pop_back();
        }
        q.push_back(v1);
    }

    cout << f[n - 1] << endl;
}


// 向量 v0 不单调，二分下凸包
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
            if (v0.dot(q[mid]) <= v0.dot(q[mid + 1])) {
                r = mid;
            } else {
                l = mid + 1;
            }
        }
        
        // 更新答案
        f[i] = v0.dot(q[l]);

        // 维护下凸包
        Vec v(f[i], i);
        while (q.size() > 1 && (q[q.size() - 1] - q[q.size() - 2]).det(v - q[q.size() - 1]) <= 0) {
            q.pop_back();
        }
        q.push_back(v);
    }
    cout << f[0] - m << endl;
}