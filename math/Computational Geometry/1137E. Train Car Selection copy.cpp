// https://codeforces.com/problemset/problem/1137/E
// 2700
// 贪心 + 凸包优化

// 一一讨论本题用的 3 种操作
// 下文假设我们从 0 下标开始记录车厢下标，则 Ai = s * (i - 1) + b 就会变成 Ai = s * i + b，最后输出的时候 +1 即可 

// 1. 在头部添加
// 此时，右边的所有车厢都没用了，因为新添加的车厢价值为 0，其余车厢价值 >= 0
// 且当 3 操作发生的时候，车厢增量为 s * i + b，由于 s, b >= 1 因此随着 i 增大，车厢价值增量越大
// 头部新添加的车厢价值最小，增加量最小，因此右边的车厢无用，可以直接忽略

// 2. 在尾部添加
// 此时我们需要维护所有的信息，但是和头部添加一样，比如我们要添加 1000 个车厢，实际上可能成为答案的只有开头的车厢
// 所以只用维护下标为 n 的车厢信息即可

// 3. 计算价值
// Ai = s * i + b
// 但是如果对所有的已知信息全部做修改，需要 O(n) 时间，显然不可以
// 假设在 3 操作之前，Ai = s * i + b
// 下一次的操作，s = s1, b = b1
// Ai = s * i + b + s1 * i + b1 = (s + s1) * i + (b + b1)
// 这与原值 Ai 分离
// 我们可以动态维护 t_s, t_b 表示 s 和 b 的和，这样我们不需要知道 Ai，只需要知道 i 就能直接求出 Ai
// 但是，比如先执行了一次 3 操作，再在末尾添加车厢，那新添加的车厢的 t_s, t_b 与之前车厢就不同了
// 我们维护一个 diff，作为偏移量，修改式子为 Ai = t_s * i + t_b + diff
// 初始状态 Ai = 0，则 diff = -t_s * i - t_b
// 只有在末尾添加的时候才需要，因为在开头加车会把所有车全部删除，因此执行 2 操作时，i = n，带入计算后 diff 就是定值

// 但是这样还是要 O(n) 时间更新，有没有什么方法能把 i 也分离出来，这样我们就能直接用 t_s, t_b 求了
// Ai = t_s * i + t_b + diff
// 把式子写成向量点积的形式，可以带来几何意义上的观察
// Ai = (t_s, 1).dot(i, diff) + t_b
// 这样我们维护 (i, diff) 信息当作备选车厢，然后用 (t_s, 1) 向量可以动态查询 Ai 的最值

// 现在描述一下代码
// 定义 q 为需要维护的下凸包，p 为操作类型，t_s 为 s 总和 t_b 为 b 总和，n 为当前车厢总长度
// 对于第一种操作，我们清空凸包，将 t_s 和 t_b 重新置为 0，更新 n 车厢长度
// 对于第二种操作，我们在凸包末尾添加 Vec(n, -t_s * n – t_b)，更新 n 车厢长度
// 对于第三种操作，我们将 t_s 与 t_b 进行累加
// 处理完之后，查询 v0 = Vec(t_s, 1) 的点积结果，维护单调队列从末尾删除

// 时间复杂度 O(m) 
// 空间复杂度 O(U) U 表示凸包中维护最多的信息长度

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
    Vec query_monotonic(const Vec& p, int dir) {
        if (dir < 0) {
            while (hull.size() > 1 && p.dot(hull[0]) >= p.dot(hull[1])) {
                hull.pop_front();
            }
            return hull.front();
        } else {
            while (hull.size() > 1 && p.dot(hull.back()) >= p.dot(hull[hull.size() - 2])) {
                hull.pop_back();
            }
            return hull.back();
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

void solve() {
    ll n, m; cin >> n >> m;
    LowerHull q;
    q.add(Vec(0, 0));
    ll t_s, t_b; t_s = 0LL; t_b = 0LL;
    for (int M = 0; M < m; ++M) {
        int p; cin >> p;
        if (p == 1) {
            q.clear();
            q.add(Vec(0, 0));
            t_s = 0LL;
            t_b = 0LL;
            ll cur; cin >> cur;
            n += cur;
        }
        
        else if (p == 2) {
            q.add(Vec(n, -t_s * n - t_b));
            ll cur; cin >> cur;
            n += cur;
        }

        else {
            ll b, s; cin >> b >> s;
            t_s += s;
            t_b += b;
        }

        Vec v0(t_s, 1);
        Vec point = q.query_monotonic(v0, 1);
        ll ans_x = point.x;
        ll ans_v = v0.dot(point) + t_b;
        cout << ans_x + 1 << ' ' << ans_v << endl;
    }
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);
    
    int T;
    T = 1;
    
    while (T--) {
        solve();
    }
    
    return 0;
}