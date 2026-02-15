// https://codeforces.com/contest/1875/problem/D
// 1600
// 子序列DP

// 正常不优化也可以通过，官解用的刷表，这里用查表
// 定义 f[i] 为以 i 为 MEX 时一直减到 0 所需要的最小花费
// 思路跟官解一致，认为要不断减少 i
// 一个细节：当最后一次减少 i 时，MEX = i，应当将 i 添加到答案中，而在此之前减少的都是 j
// 因此，状态转移方程：f[i] = min_{j = i + 1}^{m}(f[j] + (cnt[i] - 1) * j + j) = min_{j = i + 1}^{m}(f[j] + cnt[i] * j + j - i)
// 从 m 减到 0，j - i 一共对答案贡献 -m，在结尾减去即可，答案为 f[0] - m
// 状态转移方程：f[i] = min_{j = i + 1}^{m}(f[j] + cnt[i] * j)
// 复杂度 O(m^2)

// 由状态转移方程可知，可用斜率优化DP优化到 O(mlogm)，另见斜率优化DP题单

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
ll inf = LLONG_MAX;

void solve() {
    int n; cin >> n;
    unordered_map<ll, int> cnt;
    for (int i = 0; i < n; ++i) {
        ll x; cin >> x;
        cnt[x]++;
    }
    ll m = 0;
    while (cnt[m]) m++;
    vector<ll> f(m + 1, inf);
    f[m] = 0;
    for (ll i = m - 1; i >= 0; --i) {
        for (ll j = i + 1; j <= m; ++j) {
            f[i] = min(f[i], f[j] + cnt[i] * j);
        }
    }
    cout << f[0] - m << endl;
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