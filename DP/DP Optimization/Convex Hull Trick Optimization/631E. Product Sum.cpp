// https://codeforces.com/contest/631/problem/E
// 2600
// 前后缀分解 + 前缀和 + 凸包优化

// 严格意义上讲，这个题的转移过程没有用到前一个状态，即 f[i] 的取值与 f[j] 无关，只与 j 的取值有关，从递归角度应该不能成为动态规划
// 但是本题向左移动和向右移动分开计算，然后取最大的做法是标准的前后缀分解，因此归于 DP

// 首先，我们规定，本题的下标从 0 开始，定义 a[0] = 0，题目中的 n++，这样能将下标一一对应
// 定义把 a[i] 移到的 j 下标，即 i 为我们选择移动的数的下标，j 为移动的目标下标
// 分类讨论：
// 1. i = j
// 题目中说恰好移动一次，且可以放回原处，因此答案至少为 sum(a[i] * i) 即增量为 0

// 2. j < i
// 此时，a[i] 左移
// 2 4 3 6 7 5 3
// 7 2 4 3 6 5 3
// 会发现移动后，原本 7 * 5 变成了 7 * 1，而在 [1, 4] 中的所有数右移一位，2 * 1 变成了 2 * 2，增加了 2 * (2 - 1) = 2
// 4 * 2 -> 4 * 3 增加了 4 * (3 - 2) = 4
// 一般的，在 [j, i - 1] 的数 a[k]，会增加 a[k] * ((k + 1) - k) = a[k]
// 因此，从 i 左移到 j，变化量为 sum{k = [j, i - 1]}(a[k]) + a[i] * (j - i)
// 那么对于定值 i，定义 i 左移的最大变化量为 best_left，则：
// best_left[i] = max{j = [1, i - 1]}(sum{k = [j, i - 1]}(a[k]) + a[i] * (j - i))
// 显然，我们需要枚举 j，这样复杂度就是 O(n^2) 这是不被接受的
// 首先，我们用前缀和优化，定义 s[i + 1] = s[i] + a[i]，s[0] = 0
// sum{k = [j, i - 1]}(a[k]) = s[i] - s[j]
// best_left[i] = max{j = [1, i - 1]}(s[i] - s[j] + a[i] * (j - i)) = max{j = [1, i - 1]}(s[i] - s[j] + a[i] * j - a[i] * i)
// = s[i] - a[i] * i + max{j = [1, i - 1]}(a[i] * j - s[j])
// 将里面的式子写成向量点积的形式，可以带来几何意义上的观察
// a[i] * j - s[j] = (a[i], 1) * (j, -s[j]) 其中 * 代表 dot
// 则我们维护上凸包 (j, -s[j])
// 我们将 v0 = (a[i], 1) 称为查询向量，v1 = (j, s[j]) 称为待维护向量，要用凸包模板则需要保证：
// v0.y > 0，v1.x 单调递增
// 因此我们将 v0.y 设置成 1 而不是 -1
// 由于 a[i] 不具有单调性，我们查询最值的时候需要在凸包上二分

// 3. i < j
// 此时，a[i] 右移
// 画图与左移同理，这里我直接写转移方程为
// best_right[i] = max{j = [i + 1, n - 1]}(s[i + 1] - s[j + 1] + a[i] * (j - i))
// = s[i + 1] - a[i] * i + max{j = [i + 1, n - 1]}(a[i] * j - s[j + 1])
// 写成点积
// (a[i], 1) * (j, -s[j + 1])
// 维护这个上凸包需要倒序遍历，从后往前计算凸包
// 此时，j 是单调递减的，违背模板
// 我们将点积重写成 (-a[i], 1) * (-j, -s[j + 1])
// j 单调递减，-j 单调递增，就可以顺利维护上凸包

// 代码实现的时候，倒序预处理 best_right
// 将正序处理 best_left 和 计算答案的遍历合并
// best_right[n - 1] = -ll_inf 表示最后一个数 a[n - 1] 不能右移
// ans = max(0LL, best_right[1]) 等价于 ans = max(0LL, best_right[1], best_left[1])
// 但是 best_left[1] = -ll_inf，表示第一个数不能左移，与 0 取最大值为 0，忽略
// 0LL 表示将数放回原位，则答案就为 sum(a[i] * i) 即我们讨论的 i = j 的情况，表示增量为 0

// 时间复杂度 O(nlogn)
// 空间复杂度 O(n)


#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const ll ll_inf = 1e18;

// -- 凸包模板开始 -- //
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

// 上凸包（求最大值）
struct UpperHull {
    deque<Vec> hull;
    void add(const Vec& p) {
        while (hull.size() > 1 && (p - hull.back()).det(hull.back() - hull[hull.size() - 2]) <= 0) {
            hull.pop_back();
        }

        hull.push_back(p);
    }
    // 单调查询（假设查询向量 p 的斜率单调递增，即 p.x 递增）
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
// -- 凸包模板结束 -- //



void solve() {
    int n; cin >> n; ++n;
    vector<ll> a(n);
    vector<ll> s(n + 1);
    for (int i = 1; i < n; ++i) {
        cin >> a[i];
        s[i + 1] = s[i] + a[i];
    }
    
    // 右边：将 a[i] 移动到 j > i 的位置
    vector<ll> right_best(n);
    UpperHull hull_right;
    hull_right.add(Vec(-(n - 1), -s[n]));

    for (int i = n - 2; i > 0; --i) {
        ll best = hull_right.query_binary(Vec(-a[i], 1));

        right_best[i] = s[i + 1] - a[i] * i + best;

        hull_right.add(Vec(-i, -s[i + 1]));
    }

    ll ans = max(0LL, right_best[1]);

    // 左边：将 a[i] 移动到 j < i 的位置
    UpperHull hull_left;
    hull_left.add(Vec(1, -s[1]));

    for (int i = 2; i < n; ++i) {
        ll best = hull_left.query_binary(Vec(a[i], 1));

        ll left_best = s[i] - a[i] * i + best;

        ans = max(ans, max(left_best, right_best[i]));

        hull_left.add(Vec(i, -s[i]));
    }

    // 加上原始特征值
    for (int i = 1; i < n; ++i) {
        ans += a[i] * i;
    }
    cout << ans << endl;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    solve();

    return 0;
}