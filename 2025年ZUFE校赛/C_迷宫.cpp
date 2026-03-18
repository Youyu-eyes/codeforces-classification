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