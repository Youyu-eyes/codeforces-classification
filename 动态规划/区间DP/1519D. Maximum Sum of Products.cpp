// https://codeforces.com/contest/1519/problem/D
// 1600
// 前缀和 + DP (中心扩展法) 属于区间DP的优化

// 假设左端点 l，右端点 r，双闭区间
// 定义前缀和 pre[i] = sum(a[j] * b[j] for j in range(i)) pre[0] = 0
// 则 l 左边的和为 pre[l]，右边的和为 pre[n] - pre[r + 1]
// 问题就变成了中间的反转数组，求对应积的和
// 对于不同的左端点和右端点，反转后的数组对应的可能不一样
// 比如：
// 0 1 2 3 4 5
// 5 4 3 2 1 0
// 如果取
// 1 2 3 4
// 4 3 2 1
// 是一样对应的，但是如果取
// 2 3 4
// 直接对应的话就是
// 3 2 1
// 但是我们想要的是
// 4 3 2
// 所以不能预处理反转

// 观察反转以后的数组，我们可以得到一个有趣的性质，比如
// 2 3 4
// 4 3 2
// 两边同时扩展一个数
// 1 2 3 4 5
// 5 4 3 2 1
// 可以发现，这样计算中间的数，就可以保留 2 * 4 + 3 * 3 + 4 * 2 的结果，用于更新之后的翻转和
// 参考最长回文子串的中心扩展法，中心可以有一个数，也可以两个
// 用一次遍历整合

// 复杂度：
// 共 n^2 个子数组，每次更新 ans 是 O(1) 的，共 O(n^2) 复杂度，具体可见 leetcode 5.最长回文子串

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;

void solve() {
    int n; cin >> n;
    vector<ll> a(n);
    vector<ll> b(n);
    for (int i = 0; i < n; ++i) {
        cin >> a[i];
    }
    vector<ll> pre(n + 1);
    for (int i = 0; i < n; ++i) {
        cin >> b[i];
        pre[i + 1] = pre[i] + a[i] * b[i];
    }

    ll ans = 0;
    for (int i = 0; i < 2 * n - 1; ++i) {
        int l = i / 2;
        int r = (i + 1) / 2;
        ll cur = 0;
        while (l >= 0 && r < n) {
            cur += a[l] * b[r] + (l == r ? 0 : a[r] * b[l]);
            ans = max(ans, cur + pre[l] + pre[n] - pre[r + 1]);
            l--;
            r++;
        }
    }
    cout << ans << endl;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);
    
    solve();
    
    return 0;
}