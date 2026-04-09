// https://codeforces.com/contest/1083/problem/E
// 2400
// 凸包优化DP

// 将矩形按 x 升序排序后，由于没有包含矩形，y 递减
// 要在矩形数组中选择一个子序列，然后求子序列的最大价值
// 假设最后选择了下标为 i 的矩形，上一个矩形下标为 j
// 新矩形的贡献为 (x[i] - x[j]) * y[i] - a[i]
// 问题变成了最后一个选择的是下标为 j 的矩形，子序列最大价值
// 这是一个重复的子问题，可以用动态规划解决

// 定义 f[i] 为以 i 结尾的子序列最大价值
// f[i] = max{j = [0, i - 1]}(f[j] + (x[i] - x[j]) * y[i] - a[i]) = x[i] * y[i] - a[i] + max{j = [0, i - 1]}(f[j] - x[j] * y[i])
// 从递归角度考虑，每个 i 都可能作为子序列起点，此时 f[i] = x[i] * y[i] - a[i]
// 因此最终的状态转移方程为 f[i] = x[i] * y[i] - a[i] + max(0, max{j = [0, i - 1]}(f[j] - x[j] * y[i]))
// 你还可以将 (0, 0) 点当作初始点，这样就不用额外的 max 开销

// 由于 max 函数中 i 与 j 无法分离，将转移方程写成向量点积的形式，可以获得来自几何意义上的观察
// (-y[i], 1) * (x[j], f[j])，* 表示向量点积
// 这样我们可以将 j 单独分离
// 由于 y 排序后单调递减，因此我们可以用双端队列维护凸包，不需要二分凸包

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

    // 单调查询，这里假设 p 的 x 单调递增，最大值点单调右移
    // 复杂度 O(n)
    long long query_monotonic(const Vec& p) {
        while (hull.size() > 1 && p.dot(hull[0]) <= p.dot(hull[1]))
            hull.pop_front();
        return p.dot(hull.front());
    }

    // 二分查询最大值，复杂度 O(nlogn)
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

void solve() {
    int n; cin >> n;
    vector mat(n, vector<ll>(3));
    for (int i = 0; i < n; ++i) {
        cin >> mat[i][0] >> mat[i][1] >> mat[i][2];
    }
    ranges::sort(mat, {}, [](auto& a) { return a[0]; });
    
    vector<ll> f(n);
    f[0] = 0;
    UpperHull q;
    q.add(Vec(0, 0));
    ll ans = 0;

    for (int i = 0; i < n; ++i) {
        Vec v0(-mat[i][1], 1);
        ll best = q.query_monotonic(v0);

        f[i] = best + mat[i][0] * mat[i][1] - mat[i][2];
        Vec v1(mat[i][0], f[i]);

        q.add(v1);
        ans = max(ans, f[i]);
    }

    cout << ans << endl;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);
    
    solve();
    
    return 0;
}