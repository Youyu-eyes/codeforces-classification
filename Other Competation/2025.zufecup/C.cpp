// https://cpc.zufe.ac.cn/contest/7/sub/5/problem/C
// 网格图DP

// 记忆化搜索实现
// 定义 dfs(i, j, v) 为当前在 grid[i][j] 上，连续向右走了 v 步所需的最小开销
// 分类讨论
// 如果 v = 0，则一定是从 i - 1 往 i 转移，即向下走了一步
// 此时，前一次的 v 可以任意取到 [0, k] 步
// 如果前一次是 k，则这一次没有额外开销
// 如果前一次小于 k，则有额外开销
// 值得注意的是，假设从 0 开始向右走一步，v = 1，j = 1；假设从 1 开始向右走一步，v = 1，j = 2
// 以此类推，可知 v <= j
// 则 v 取值为 [0, min(k, j)]
// 如果 v = 1，则从 j - 1 转移，前一次步数可能为 0，也可能为 k，同理为 k 不加额外开销
// 剩下的情况就是从 dfs(i, j - 1, v - 1) 转移，有额外开销

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
const int inf = 0x3f3f3f3f;
const ll ll_inf = 1e18;

void solve() {
    int m, n, k; cin >> m >> n >> k;
    vector<vector<int>> grid(m, vector<int>(n));
    for (int i = 0; i < m; ++i) {
        for (int j = 0; j < n; ++j) {
            cin >> grid[i][j];
        }
    }

    if (grid[0][0] == 0) {
        cout << -1 << endl;
        return;
    }

    vector<vector<vector<int>>> memo(m, vector<vector<int>>(n, vector<int>(k + 1, inf)));
    auto dfs = [&](this auto&& dfs, int i, int j, int v) -> int {
        if (i == 0 && j == 0) {
            return (v == 0 ? grid[0][0] : inf);
        }
        if (i < 0 || j < 0) {
            return inf;
        }
        if (j < v) {
            return inf;
        }
        int& res = memo[i][j][v];
        if (res < inf) {
            return res;
        }
        
        int x = grid[i][j];
        if (x == 0) {
            return inf;
        }
        
        if (v > 1) {
            res = dfs(i, j - 1, v - 1) + x;
        } else if (v == 0) {
            res = dfs(i - 1, j, k);
            for (int t = 0; t < min(k, j + 1); ++t) {
                res = min(res, dfs(i - 1, j, t) + x);
            }
        } else {
            res = min(dfs(i, j - 1, 0) + x, dfs(i, j - 1, k));
        }
        
        return res;
    };

    int ans = inf;
    for (int u = 0; u <= k; ++u) {
        ans = min(ans, dfs(m - 1, n - 1, u));
    }

    cout << (ans < inf ? ans : -1) << endl;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    solve();
    
    return 0;
}