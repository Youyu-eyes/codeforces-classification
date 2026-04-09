#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
const int inf = 0x3f3f3f3f;
const ll ll_inf = 1e18;

struct Vec {
    ll x, y;
    Vec(ll x = 0, ll y = 0) : x(x), y(y) {}
    Vec operator-(const Vec& other) const {
        return Vec(x - other.x, y - other.y);
    }
    ll dot(const Vec& other) const {
        return x * other.x + y * other.y;
    }
    ll det(const Vec& other) const {
        return x * other.y - y * other.x;
    }
};

// 上凸包（求最大值）
struct UpperHull {
    deque<Vec> hull;
    void add(const Vec& p) {
        while (hull.size() > 1) {
            Vec& back1 = hull.back();
            Vec& back2 = hull[hull.size()-2];
            // 维护上凸包：新点使前一个点成为凹点则弹出
            if ((p - back1).det(back1 - back2) >= 0) hull.pop_back();
            else break;
        }
        hull.push_back(p);
    }
    // 单调查询（假设查询向量 p 的斜率单调递增，即 p.x 递增？）
    // 这里假设 p 的 x 单调递增，对于上凸包，最大值点单调右移
    long long query_monotonic(const Vec& p) {
        while (hull.size() > 1 && p.dot(hull[0]) <= p.dot(hull[1]))
            hull.pop_front();
        return p.dot(hull.front());
    }
    // 二分查询（p 不单调）
    long long query_binary(const Vec& p) const {
        int l = 0, r = hull.size()-1;
        while (l < r) {
            int mid = (l+r)>>1;
            if (p.dot(hull[mid]) <= p.dot(hull[mid+1])) l = mid+1;
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
        while (hull.size() > 1) {
            Vec& back1 = hull.back();
            Vec& back2 = hull[hull.size()-2];
            // 维护下凸包：新点使前一个点成为凸点？条件相反
            if ((p - back1).det(back1 - back2) <= 0) hull.pop_back();
            else break;
        }
        hull.push_back(p);
    }
    // 单调查询（假设 p 的 x 单调递增，最小值点单调右移）
    long long query_monotonic(const Vec& p) {
        while (hull.size() > 1 && p.dot(hull[0]) >= p.dot(hull[1]))
            hull.pop_front();
        return p.dot(hull.front());
    }
    // 二分查询最小值
    long long query_binary(const Vec& p) const {
        int l = 0, r = hull.size()-1;
        while (l < r) {
            int mid = (l+r)>>1;
            if (p.dot(hull[mid]) >= p.dot(hull[mid+1])) l = mid+1;
            else r = mid;
        }
        return p.dot(hull[l]);
    }
    bool empty() const { return hull.empty(); }
    void clear() { hull.clear(); }
};