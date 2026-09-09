// https://codeforces.com/problemset/problem/1900/C
// 1300
// 递归遍历 + 贪心
// 如果是 'U'，一定要改，没有必要回头
// 不回头的话，要么向左要么向右
// 两种写法，传参或不传参

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;

// 传参 cost
void solve() {
    int n; cin >> n;
    string dir; cin >> dir;
    vector<vector<int>> tree(n, vector<int>(2, -1));
    for (int i = 0; i < n; ++i) {
        int l, r; cin >> l >> r;
        tree[i][0] = l - 1;
        tree[i][1] = r - 1;
    }
    int ans = n + 1;
    auto dfs = [&](this auto&& dfs, int node, int cost) -> void {
        if (tree[node][0] < 0 && tree[node][1] < 0) {
            ans = min(ans, cost);
            return;
        }
        if (tree[node][0] >= 0) {
            int n_cost = cost + (dir[node] == 'L' ? 0 : 1);
            dfs(tree[node][0], n_cost);
        }
        if (tree[node][1] >= 0) {
            int n_cost = cost + (dir[node] == 'R' ? 0 : 1);
            dfs(tree[node][1], n_cost);
        }
    };
    dfs(0, 0);
    cout << ans << endl;
}

// 不传参 cost
void solve() {
    int n; cin >> n;
    string dir; cin >> dir;
    vector<vector<int>> tree(n, vector<int>(2, -1));
    for (int i = 0; i < n; ++i) {
        int l, r; cin >> l >> r;
        tree[i][0] = l - 1;
        tree[i][1] = r - 1;
    }
    auto dfs = [&](this auto&& dfs, int node) -> int {
        if (tree[node][0] < 0 && tree[node][1] < 0) {
            return 0;
        }
        if (tree[node][0] < 0) {
            return dfs(tree[node][1]) + (dir[node] == 'R' ? 0 : 1);
        }
        if (tree[node][1] < 0) {
            return dfs(tree[node][0]) + (dir[node] == 'L' ? 0 : 1);
        }
        return min(dfs(tree[node][0]) + (dir[node] == 'L' ? 0 : 1), dfs(tree[node][1]) + (dir[node] == 'R' ? 0 : 1));
    };
    int ans = dfs(0);
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