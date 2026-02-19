// https://codeforces.com/contest/1741/problem/E
// 1600
// 划分型DP + 刷表查表同时进行
// 别样的 DP 思考角度

// 对于划分型DP，通用状态定义方式为定义 f[i] 为前 i 个数是否可以分割
// 本题比较不同的是，要从 b[i] 开始考虑，而不是从 f[i] 开始考虑
// 说明：如果从 f[i] 开始考虑
// 从 f[j] 转移，前一个划分的数组有两种情况
// b[j: i - 1]，b[i - 1] 为长度，即 b[i - 1] = i - j - 1
// => j = i - b[i - 1] - 1
// b[j + 1: i]，b[j] 为长度，即 b[j] = i - j - 1
// 到这里发现，无法从 b[j] = i - j - 1 中单独算出 j
// 由于数据范围 2 * 10 ^ 5，枚举 j 复杂度 O(n^2) 超时
// 从 b[i] 的角度考虑
// 假设 b[i] 为长度，则有两个可选数组，
// 考虑 b[i] 在右端，子数组为 b[i - b[i]: i]
// 考虑 b[i] 在左端，子数组为 b[i + 1: i + b[i] + 1]
// 前者从查表角度考虑，用 f[i - b[i]] 更新 f[i + 1]
// 后者从刷表角度考虑，用 f[i] 更新 f[i + b[i] + 1]
// 为什么更新的是 f[i + 1] 而不是 f[i]：
// 我们定义的是前 i 个数，思考的 i 为左端点，对于子数组 b[i - b[i]: i] 来说，右端点是 i - 1
// 而 b[i] 是长度，也要被划分进去，前面就有 0 到 i 共 i + 1 个数，所以更新的是 f[i + 1]
// 用 f[i] 更新同理，画图即可明白

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;

void solve() {
    int n; cin >> n;
    vector<int> b(n);
    for (int i = 0; i < n; ++i) {
        cin >> b[i];
    }

    vector<int> f(n + 1, 0);
    f[0] = 1;
    for (int i = 0; i < n; ++i) {

        // 查表
        if (i - b[i] >= 0) {
            f[i + 1] |= f[i - b[i]];
        }

        // 刷表
        if (i + b[i] + 1 <= n) {
            f[i + b[i] + 1] |= f[i];
        }
    }
    cout << (f.back() ? "YES" : "NO") << endl;
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