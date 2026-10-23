// https://codeforces.com/problemset/problem/660/F
// 2500
// 前缀和 + 枚举右维护左 + 凸包优化 DP

// 官解是折半枚举 + 凸包优化
// 这里采用前缀和 + 数形结合的思想

// 假设 a = [0, 1, 2, 3, 4, 5, 6, 7, 8 ……] 直接对应下标

//             -
//           - -
//         - - -
//       - - - -
//     - - - - -
//   - - - - - -
// - - - - - - -
// 0 1 2 3 4 5 6 7 8 ……

// 如果左端点是 0，得分为 ws
// 对于右端点 i，定义得分为 ws[i + 1]
// 如图，我们画出来的就是 ws[7]
// 我们假设左端点是 3，则在图形中显示为

//             -
//           - -
//         - - -
//       - - - -
// 
// 
// 
// 0 1 2 3 4 5 6 7 8 ……

// 注意我们要减掉的部分

//     - - - - -
//   - - - - - -
// - - - - - - -
// 0 1 2 3 4 5 6 7 8 ……

// 可以发现，需要减掉的地方分为两部分
// 从 2, 3 中间切一刀，左边就是 ws[3]，右边是一个矩形，面积为 长 * 宽
// 长是 a[3] + a[4] + a[5] + a[6]
// 显然可以用前缀和优化成 s[7] - s[3]
// 而左半部分是一个等腰三角形，长度为 3，因此右半部分宽度也为 3

// 现在不失一般性的抽象化上述过程
// 假设右端点为 i，左端点为 j
// 在固定右端时，计算左端点的最值
// 我们需要计算的部分可以抽象为
// max{j = [0, i - 1]}(ws[i] - ws[j] - (s[i] - s[j]) * j) = ws[i] + max{j = [0, j - 1]}(-s[i] * j + s[j] * j - ws[j])
// 然后枚举 i，计算所有右端点的大小即可
// 但是注意到我们无法将 i 和 j 分离，如果强行枚举 j，复杂度 O(n^2) 无法接受

// 将 max 内部的式子写成向量点积的形式，可以带来几何意义上的观察
// (-s[i], 1).dot(j, s[j] * j - ws[j]), dot 代表点积
// 维护 (j, s[j] * j - ws[j]) 形成的上凸包
// 用 (-s[i], 1) 向量查询即可
// 由于 a 中有负数，因此 s[i] 不单调，需要在凸包上二分

// 时间复杂度 O(nlogn)
// 空间复杂度 O(n)

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
        if (dir > 0) {
            while (hull.size() > 1 && p.dot(hull[0]) <= p.dot(hull[1])) {
                hull.pop_front();
            }
            return p.dot(hull.front());
        } else {
            while (hull.size() > 1 && p.dot(hull.back()) <= p.dot(hull[hull.size() - 2])) {
                hull.pop_back();
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

void solve() {
    int n; cin >> n;
    vector<ll> a(n);
    for (int i = 0; i < n; ++i) {
        cin >> a[i];
    }
    vector<ll> s(n + 1);
    vector<ll> ws(n + 1);
    for (int i = 0; i < n; ++i) {
        s[i + 1] = s[i] + a[i];
        ws[i + 1] = ws[i] + a[i] * (i + 1);
    }
    UpperHull q;
    q.add(Vec(0, 0));
    ll ans = 0;
    for (int i = 1; i <= n; ++i) {
        Vec v0(-s[i], 1);
        ans = max(ans, q.query_binary(v0) + ws[i]);
        Vec v1(i, s[i] * i - ws[i]);
        q.add(v1);
    }
    cout << ans << endl;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    solve();
    
    return 0;
}