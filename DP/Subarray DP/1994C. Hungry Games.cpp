// https://codeforces.com/contest/1994/problem/C
// 1600
// 子数组DP + 滑窗优化
// 由于本题的毒性都是大于 0 的，因此后缀和递增
// 可采用二分(O(nlogn))或滑窗(O(n))优化
// 官解中采用二分，本解中采用滑窗优化

// 先思考简单的问题：如果自身毒性 g > x，直接停止
// 显然，是一个越短越合法滑窗
// 假设 sum(a[i: j + 1]) <= x，sum(a[i: j + 2]) > x
// 即选 j + 1 后重新开始，那么在 i 到 j 的区间内所有下标作为右端点都可，方案为 j - i + 1

// 本题 g > x 后，g = 0
// 相当于 g > x 后，重新开始吃
// 问题被转化成了规模更小的子问题
// 定义左端点为 i 时的方案数为 f[i]
// f[i] 由两部分组成，一是 i~j 的每一个下标，二是选到 j + 1，g 归零，则由 f[j + 2] 转移
// f[i] = f[j + 2] + j - i + 1
// 官解中 j 用二分求得，其实可以直接由滑窗（双指针）求得

// 复杂度 O(n)

// 为什么从后往前遍历，留给读者思考
// 提示：与吃的方向有关，本题 从左向右吃 和 从右向左吃 是不一样的

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;

void solve() {
    int n;
    ll x;
    cin >> n >> x;
    vector<ll> a(n);
    for (int i = 0; i < n; ++i) {
        cin >> a[i];
    }
    int j = n - 1;
    vector<ll> f(n + 2);
    ll s = 0;
    ll ans = 0;
    for (int i = n - 1; i >= 0; --i) {
        s += a[i];
        while (s > x) {
            s -= a[j--];
        }
        f[i] = f[j + 2] + j - i + 1;
        ans += f[i];
    }
    cout << ans << endl;
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