// https://codeforces.com/problemset/problem/2167/G
// 1600
// DP
// 式子转化，cost = sum(c) - sum_non-decreasing(c)
// 求 min(cost)，即求 max(sum_non-decreasing(c))
// 转化为求最长的非递减子序列的 sum(c)
// 这里需要证明，为什么保留 a 的最长非递减子序列一定能够将 a 转化成非递减数组
// 因为题目要求非递减，假设选择的最长非递减子序列中两个连续的数 a[i]，a[j]
// 贪心的想，我们可以将 a[i + 1: j] 中所有的数更改成 a[i](a[j]也可)

// 相邻相关，考虑枚举选哪个
// 状态转移方程为 f[i] = max_{j = 0}^{i - 1 && a[j] <= a[i]}(f[j]) + c[i]
// 由于转移不是连续区间最大值，跟具体数据有关，无法用一个遍历维护最大值，二维DP可过

// 时间复杂度 O(n^2)
// 空间复杂度 O(n)

// 本题需要满足 j < i 与 a[j] < a[i] 两个条件，是二维偏序问题，由于第一维度天然满足，不需要排序
// 可用树状数组或线段树优化至 O(n log n)
// 详见线段树优化 DP

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;

void solve() {
    int n; cin >> n;
    vector<int> a(n);
    for (int i = 0; i < n; ++i) {
        cin >> a[i];
    }
    vector<int> c(n);
    ll ans = 0;
    for (int i = 0; i < n; ++i) {
        cin >> c[i];
        ans += c[i];
    }
    vector<ll> f(n);
    ll sum_max = 0;
    for (int i = 0; i < n; ++i) {
        f[i] = c[i];
        ll mx = 0;
        for (int j = 0; j < i; ++j) {
            if (a[j] <= a[i]) {
                mx = max(mx, f[j]);
            }
        }
        f[i] += mx;
        sum_max = max(sum_max, f[i]);
    }
    cout << ans - sum_max << endl;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);
    
    int T;
    cin >> T;
    
    while (T--) {
        solve();
    }
    
    return 0;
}