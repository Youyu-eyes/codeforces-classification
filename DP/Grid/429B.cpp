// https://codeforces.com/contest/429/problem/B
// 1600 估计现在 1400
// 网格图DP + 前后缀分解

// 题目有一点没有说明清楚，两个人的路径只能在一个格子上重叠，不能有两个格子重叠
// 可以参考官解里的图
// 定义 f1[i][j] 表示从 (0, 0) 出发到 (i, j) 的最大路径和
// 定义 f2[i][j] 表示从 (m - 1, 0) 出发到 (i, j) 的最大路径和
// 定义 f3[i][j] 表示从 (0, n - 1) 出发到 (i, j) 的最大路径和
// 定义 f4[i][j] 表示从 (m - 1, n - 1) 出发到 (i, j) 的最大路径和
// 问题转化为从 grid[i][j] 出发，到四个点的最大路径和，注意只有两种情况

// 由于我们最后要枚举断点，而每个DP矩阵所需的虚拟边界条件都不一样，所以需要在DP矩阵的四周都添加一个虚拟边界，方便后续的计算

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;

void solve() {
    int m, n;
    cin >> m >> n;
    vector<vector<int>> grid(m, vector<int>(n));
    for (int i = 0; i < m; ++i)
        for (int j = 0; j < n; ++j)
            cin >> grid[i][j];

    vector<vector<ll>> f1(m + 2, vector<ll>(n + 2, 0));
    vector<vector<ll>> f2(m + 2, vector<ll>(n + 2, 0));
    vector<vector<ll>> f3(m + 2, vector<ll>(n + 2, 0));
    vector<vector<ll>> f4(m + 2, vector<ll>(n + 2, 0));

    for (int i = 1; i <= m; ++i)
        for (int j = 1; j <= n; ++j)
            f1[i][j] = max(f1[i - 1][j], f1[i][j - 1]) + grid[i - 1][j - 1];

    for (int i = m; i >= 1; --i)
        for (int j = 1; j <= n; ++j)
            f2[i][j] = max(f2[i + 1][j], f2[i][j - 1]) + grid[i - 1][j - 1];

    for (int i = 1; i <= m; ++i)
        for (int j = n; j >= 1; --j)
            f3[i][j] = max(f3[i - 1][j], f3[i][j + 1]) + grid[i - 1][j - 1];

    for (int i = m; i >= 1; --i)
        for (int j = n; j >= 1; --j)
            f4[i][j] = max(f4[i + 1][j], f4[i][j + 1]) + grid[i - 1][j - 1];

    ll ans = 0;
    for (int i = 2; i < m; ++i) {
        for (int j = 2; j < n; ++j) {
            ll val1 = f1[i][j - 1] + f4[i][j + 1] + f2[i + 1][j] + f3[i - 1][j];
            ll val2 = f1[i - 1][j] + f4[i + 1][j] + f2[i][j - 1] + f3[i][j + 1];
            ans = max(ans, max(val1, val2));
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