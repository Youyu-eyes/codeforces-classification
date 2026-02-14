// https://codeforces.com/problemset/problem/2167/G
// 1600
// DP
// 式子转化，cost = sum(c) - sum_non-decreasing(c)
// 求 min(cost)，即求 max(sum_non-decreasing(c))
// 转化为求 sum(c[i]) 最大的递增子序列
// 相邻相关，考虑枚举选哪个
// 状态转移方程为 f[i] = max_{j = 0}^{i - 1 && a[j] <= a[i]}(f[j]) + c[i]
// 由于转移不是连续区间最大值，跟具体数据有关，无法用一个遍历维护最大值，二维DP可过

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