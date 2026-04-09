// https://codeforces.com/problemset/problem/319/C
// 2100 实际分数大约 1800
// 贪心 + 凸包/斜率优化DP
// 凸包优化模板题

// 首先我们需要贪心地证明两个结论
// 一. 先证明砍一棵树就一次性砍到底，而不是这棵树砍一刀另一个树砍一刀
// 定义操作一：砍一棵树砍完，然后再砍第二棵树
// 定义操作二：一棵树砍到剩一刀，然后另一棵树砍完，再砍第一棵树

// 假设第 i 棵树是目前砍掉的下标最大的树

// 如果现在要砍的两棵树是 i - 2, i - 1
// 操作一的开销：b[i] * a[i - 2] + b[i] * a[i - 1] = b[i] * (a[i - 2] + a[i - 1])
// 因为无论先砍 i - 2 还是 i - 1，最大下标仍是 i
// 操作二的开销：b[i] * (a[i - 2] - 1) + b[i] * a[i - 1] + b[i] = b[i] * (a[i - 2] + a[i - 1])
// 看到两者的开销相等，所以当要砍的树下标在最大下标左侧时，开销仅取决于 b[i] 的大小，贪心的想，i 越大，b[i] 越小，总开销越小

// 如果现在要砍的两棵树是 i + 1, i + 2
// 操作一的开销：min(b[i] * a[i + 1] + b[i + 1] * a[i + 2], b[i] * a[i + 2] + b[i + 2] * a[i + 1])
// 在这里我们能看到，对于 a[i + 1]，先砍 i + 1 的系数为 b[i]，先砍 i + 2 的系数为 b[i + 2]，b[i + 2] < b[i]，所以先砍 i + 2 优
// 而对于 a[i + 2]，先砍 i + 1 的系数为 b[i + 1]，先砍 i + 2 的系数为 b[i]，b[i + 1] < b[i]，所以先砍 i + 1 优
// 发现两者矛盾，证明此题没有最优贪心策略
// 操作二的开销：这里假设操作二先砍左边的树剩一刀，然后砍完右边的树，再砍左边的树
// b[i] * (a[i + 1] - 1) + b[i] * a[i + 2] + b[i + 2] = b[i] * (a[i + 1] + a[i + 2]) + b[i + 2] - b[i]
// 而操作一第二种情况：b[i] * a[i + 2] + b[i + 2] * a[i + 1] = b[i] * (a[i + 1] + a[i + 2]) - b[i] * a[i + 1] + b[i + 2] * a[i + 1] = b[i] * (a[i + 1] + a[i + 2]) + a[i + 1] * (b[i + 2] - b[i])
// 发现 a[i + 2] * (b[i + 2] - b[i]) < b[i + 2] - b[i] 因为 b[i + 2] < b[i] && a[i + 2] > 0
// 而反过来，先砍左边的树，证明方法是一样的，这里略过

// 因此证明操作一的开销一定小于操作二

// 二. 需要证明我们每次只会向右选择一棵树砍掉，不会选择左边的树
// 仍假设第 i 棵树是目前砍掉的下标最大的树
// 假设需要砍的树为 i - 1, i + 1
// 先砍左边：a[i - 1] * b[i] + a[i + 1] * b[i]
// 先砍右边：a[i + 1] * b[i] + a[i - 1] * b[i + 1] = a[i - 1] * b[i + 1] + a[i + 1] * b[i]
// 显然，由于 b[i] > b[i + 1]，针对两棵树分布 i 两端的情况下，先砍右边一定更优
// 所以我们再砍树时，不会破坏树的原有顺序，即我们会选择一个子序列然后从左往右砍
// 由于 b[-1] = 0，因此我们在砍完最右边的树后，用 0 的开销砍剩下的树，就能实现 0 的额外开销
// 所以总开销就是从 0 开始向右砍树，直到砍掉最后一棵树的开销

// 问题就变成了从 a 中选择一个子序列，求所有子序列的最小总开销
// 假设我们要求以 i 结尾的子序列的最小开销，上一次砍的树为 j，则这一次砍树的开销为 cost(i, j) = a[i] * b[j]
// 此时问题就变成了以 j 结尾的子序列的最小开销
// 这是一个重复的子问题，可以用动态规划解决
// 定义 f[i] 为以 i 结尾的子序列的最小开销，转移方程为 f[i] = min{j = [0, i - 1]}(f[j] + a[i] * b[j])
// 由于式子跟每次的 i 有关，我们不能维护 f[j] + a[i] * b[j] 作为变量更新 f[i]，且无法分离 a[i]
// 此时，我们需要采用凸包优化DP

// 凸包优化DP
// 将 f[j] + a[i] * b[j] 的方程写成向量乘积的形式，我们可以得到几何意义上的观察
// f[j] + a[i] * b[j] = Vec(a[i], 1) * Vec(b[j], f[j])，其中，等式右边的乘号意义是点积
// 这样我们就能单独维护 Vec(b[j], f[j]) 而不受 i 的影响了
// 由于求 min，维护下凸包
// 由于 a[i] 单调递增，向量的旋转方向固定，可以用双端队列维护凸包
// 更新 f[i]，用 f[i] 计算新的向量
// 然后用 Andrew 算法 维护下凸包

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
    __int128 det(const Vec& other) const {
        return (__int128) x * other.y - (__int128) y * other.x;
    }
};

// 下凸包（求最小值）
struct LowerHull {
    deque<Vec> hull;
    void add(const Vec& p) {
        while (hull.size() > 1 && (p - hull.back()).det(hull.back() - hull[hull.size() - 2]) >= 0) {
            hull.pop_back();
        }

        hull.push_back(p);
    }

    // 单调查询，这里假设 p 的 x 单调递增，最小值点单调右移
    // 复杂度 O(n)
    long long query_monotonic(const Vec& p) {
        while (hull.size() > 1 && p.dot(hull[0]) >= p.dot(hull[1]))
            hull.pop_front();
        return p.dot(hull.front());
    }

    // 二分查询最小值，复杂度 O(nlogn)
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

void solve() {
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
    LowerHull q;
    q.add(Vec(-b[0], 0));
    for (int i = 1; i < n; ++i) {
        Vec v0(-a[i], 1);

        f[i] = q.query_monotonic(v0);
        Vec v1(-b[i], f[i]);

        q.add(v1);
    }

    cout << f[n - 1] << endl;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    solve();
    
    return 0;
}