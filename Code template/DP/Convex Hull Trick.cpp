#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
const int inf = 0x3f3f3f3f;
const ll ll_inf = 1e18;

// 在使用模板时要注意：
// 首先，在维护凸包时要保证查询向量 v0 的 y 值 v0.y > 0，这样我们才能用 max or min 简单地判断用 上凸包 or 下凸包
// 其次，在维护凸包时要保证待维护向量 v1 的 x 值 v1.x 单调递增

struct Vec {
    long long x, y;
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

// 上凸包（求最大值）
struct UpperHull {
    deque<Vec> hull;
    void add(const Vec& p) {
        while (hull.size() > 1 && (hull.back() - hull[hull.size() - 2]).det(p - hull.back()) >= 0) {
            hull.pop_back();
        }

        hull.push_back(p);
    }

    // 保证 v0.x 单调
    // 如果 v0.x 单调递增，则 UpperHull.query_monotonic(v0,  1)
    // 如果 v0.x 单调递减，则 UpperHull.query_monotonic(v0, -1)
    // 复杂度 O(n)
    long long query_monotonic(const Vec& p, int dir) {
        long long n = hull.size();
        if (dir > 0) {
            while (n > 1 && p.dot(hull[0]) <= p.dot(hull[1])) {
                hull.pop_front();
            }
            return p.dot(hull.front());
        }
        
        else {
            while (n > 1 && p.dot(hull.back()) <= p.dot(hull[n - 2])) {
                hull.pop_front();
            }
            return p.dot(hull.back());
        }
    }

    // 二分查询最大值，复杂度 O(nlogn)
    long long query_binary(const Vec& p) const {
        int l = 0, r = hull.size() - 1;
        while (l < r) {
            int mid = (l + r) >> 1;
            if (p.dot(hull[mid]) <= p.dot(hull[mid + 1])) l = mid + 1;
            else r = mid;
        }
        return p.dot(hull[l]);
    }
    bool empty() const { return hull.empty(); }
    void clear() { hull.clear(); }
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