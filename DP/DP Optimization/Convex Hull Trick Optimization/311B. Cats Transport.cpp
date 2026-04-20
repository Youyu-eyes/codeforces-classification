// https://codeforces.com/problemset/problem/311/B
// 2400
// 式子变形 + 划分型 DP + 凸包优化 DP

// 首先，我们需要先定义 0-base 情况下的数据
// 定义猫的编号为 0 ~ m-1，山的编号为 0 ~ n-1
// 定义 d 从 d[0] 到 d[n - 1]，其中 d[0] = 0
//      h 从 h[0] 到 h[m - 1]
//      t 从 t[0] 到 t[m - 1]
// 定义 pre_d[i + 1] = pre_d[i] + d[i]，pre_d[0] = 0
// 前缀和的意义：pre_d[i] 为第 0 座山到第 i - 1 座山的距离
// 假设一个饲养员从 T 时间出发，要带走第 i 只猫，则：
// T + pre_d[h[i] + 1] >= t[i]
// T >= t[i] - pre_d[h[i] + 1]
// 定义 a[i] = t[i] - pre_d[h[i] + 1] 表示如果要接到第 i 只猫，饲养员所需的最早出发时间
// 等待时间为 T - a[i]
// 将 a 降序排序，我们可以得到一个长为 m 的数组，其中每个饲养员需要负责连续子数组的猫


// 交换论证法：
// 假设 a = [a0, a1, a2, a3, a4, a5, a6]，本来划分为 a0 ~ a3 给饲养员 1，划分 a4 ~ a6 给饲养员 2，a 降序
// 此时饲养员 1 的最优方案是在 a0 时间出发，T = a0
// T - a0 + T - a1 + T - a2 + T - a3 = 4 * a0 - (a0 + a1 + a2 + a3)
// 同理饲养员 2 的最优方案得到的最小等待时间为
// 3 * a4 - (a4 + a5 + a6)

// 不失一般性的讨论交换，一定存在一种交换是相邻交换
// 讨论 a3 和 a4 交换，饲养员 1 负责 a0, a1, a2, a4，饲养员 2 负责 a3, a5, a6
// 饲养员 1 的最优方案：4 * a0 - (a0 + a1 + a2 + a4)
// 饲养员 2 的最优方案：3 * a3 - (a3 + a5 + a6)

// 未交换总等待时间 - 交换后总等待时间：4 * a0 + 3 * a4 - 4 * a0 - 3 * a3 = 3 * (a4 - a3) < 0
// 因此未交换一定大于交换后的总等待时间


// 问题转变为将 a 数组划分为 p 段所需的最小开销
// 这是非常经典的划分型 DP

// 定义 f[k][i] 为将前 i 个数划分成 k 段的最小开销
// f[k][i] = min{j = [k - 1, i - 1]}(f[k - 1][j] + cost(j, i))

// 计算 cost(j, i)，现在 a 是降序排序的，a[j] 是最大的，即最晚的出发时间
// cost(j, i) = (i - j) * a[j] - (a[j] + a[j + 1] + …… + a[i - 1])
// 定义 s[i + 1] = s[i] + a[i], s[0] = 0
// 则 cost(j, i) = (i - j) * a[j] - s[i] + s[j]
// f[k - 1][j] + a[j] * i - a[j] * j - s[i] + s[j]
// 将式子写成点积的形式，可以得到几何意义上的观察
// (-i, 1).dot(-a[j], f[k - 1][j] - a[j] * j + s[j])

// 其中 v0 = (-i, 1) 为查询向量
// v1 = (-a[j], f[k - 1][j] - a[j] * j + s[j])
// 为了方便用凸包模板，v0.y > 0, v1.x 单调递增，因此定义为 -a[j] 

// 维护 v1 构成的下凸包，用 v0 查询
// v0 是单调的，因此可以用单调队列维护
// 代码实现的时候，采用空间优化

// 在空间优化的时候需要注意
// 新手可以采用 n_f 的滚动数组，不容易出错
// 如果一定想用一个数组，则需要先用 f[i] 构造 v1，再更新 f[i]
// 因为更新 f[i] 后 f[i] 相当于 f[k][i]，但是更新 v1 时我们需要 f[k - 1][i]

// 时间复杂度 O(mlogm + p * m)
// 空间复杂度 O(m) 不算题目数据带来的开销，这里计算的是额外空间

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

struct LowerHull {
    deque<Vec> hull;
    void add(const Vec& p) {
        while (hull.size() > 1 && (hull.back() - hull[hull.size() - 2]).det(p - hull.back()) <= 0) {
            hull.pop_back();
        }

        hull.push_back(p);
    }

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
    int n, m, p; cin >> n >> m >> p;
    vector<ll> pre_d(n + 1);
    for (int i = 2; i <= n; ++i) {
        ll d; cin >> d;
        pre_d[i] = pre_d[i - 1] + d;
    }

    if (p >= m) {
        cout << 0 << endl;
        return;
    }
    
    vector<ll> a(m);
    for (int i = 0; i < m; ++i) {
        int h; ll t; cin >> h >> t;
        a[i] = t - pre_d[h];
    }
    ranges::sort(a, greater<>());
    vector<ll> s(m + 1);
    for (int i = 0; i < m; ++i) {
        s[i + 1] = s[i] + a[i];
    }

    vector<ll> f(m + 1, ll_inf);
    f[0] = 0;
    LowerHull q;
    for (int k = 1; k <= p; ++k) {
        q.add(Vec(-a[k - 1], f[k - 1] - a[k - 1] * (k - 1) + s[k - 1]));
        for (int i = k; i <= m + k - p; ++i) {
            Vec v0(-i, 1);
            ll best = q.query_monotonic(v0);
            
            // 先构造 v1，相当于用 f[k - 1][i] 更新 v1
            if (i < m) {
                Vec v1(-a[i], f[i] - a[i] * i + s[i]);
                q.add(v1);
            }

            // 再更新 f[i]，相当于更新 f[k][i]
            f[i] = best - s[i];
        }
        q.clear();
    }
    cout << f[m] << endl;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    solve();
    
    return 0;
}


// 有一个小 trick，虽然无法提升复杂度
// 细心的同学可能在交换论证的时候观察到了
// 在我们交换后，前后总等待时间加总，观察减号右边
// (a0 + a1 + a2 + a3) + (a4 + a5 + a6) = a0 + …… + a6
// (a0 + a1 + a2 + a4) + (a3 + a5 + a6) = a0 + …… + a6
// 我们相减的时候直接抵消，但是可以观察到，如果要把猫全部带走，右边减号部分必然为 a 数组的和
// 现在我们可以简化式子
// 定义 f[k][i] 前 i 个数划分为 k 段的总出发时间（注意不是等待时间），即 sum(T)
// f[k][i] = min{j = [k - 1, i - 1]}(f[k - 1][j] + a[j] * i - a[j] * j)
// 最终等待时间为 f[k][m] - sum(a)

// 这样虽然不能降低复杂度，但是能简化式子
// 不好的地方在于原式可以缩小 f[k][i] 的大小
// 如果最后一起减，f[k][i] 会比较大
// 测试的时候，下面的代码时空开销会略大一点，不过无伤大雅，能够简单化自己的方程也是一种优势

void solve() {
    int n, m, p; cin >> n >> m >> p;
    vector<ll> pre_d(n + 1);
    for (int i = 2; i <= n; ++i) {
        ll d; cin >> d;
        pre_d[i] = pre_d[i - 1] + d;
    }

    if (p >= m) {
        cout << 0 << endl;
        return;
    }
    
    vector<ll> a(m);
    for (int i = 0; i < m; ++i) {
        int h; ll t; cin >> h >> t;
        a[i] = t - pre_d[h];
    }
    ranges::sort(a, greater<>());
    ll s = 0LL;
    for (int i = 0; i < m; ++i) {
        s += a[i];
    }

    vector<ll> f(m + 1, ll_inf);
    f[0] = 0;
    LowerHull q;
    for (int k = 1; k <= p; ++k) {
        q.add(Vec(-a[k - 1], f[k - 1] - a[k - 1] * (k - 1)));
        for (int i = k; i <= m + k - p; ++i) {
            Vec v0(-i, 1);
            ll best = q.query_monotonic(v0);
            
            // 先构造 v1，相当于用 f[k - 1][i] 更新 v1
            if (i < m) {
                Vec v1(-a[i], f[i] - a[i] * i);
                q.add(v1);
            }

            // 再更新 f[i]，相当于更新 f[k][i]
            f[i] = best;
        }
        q.clear();
    }
    cout << f[m] - s << endl;
}