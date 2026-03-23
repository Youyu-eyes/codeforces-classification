// https://codeforces.com/problemset/problem/2070/D
// 1600
// 树上DP + BFS

// 假设我们在树的第 i 层，节点为 x 的位置
// 思考上一行，假设上一行的节点列表为 q，除了 x 的父节点 fa 以外，所有的节点都可以向 x 走
// 因此 能走到 x 路径总数 = 能走到 q 中所有节点的路径总数 - 能走到 fa 的路径总数
// 问题就变成了求 q 中每一个节点的路径，对于 q 中一个单独的节点，这是一个相似的子问题，可以用动态规划解决
// 定义 能走到 x 的路径总数为 f[x]
// f[x] = sum(f[y] for y in q) - f[fa]
// 对于这个状态转移，我们是从上一层转移到下一层，因此我们采用 BFS 方法遍历这棵树
// 而在遍历 x 的过程中如果还要暴力遍历上一层节点求 sum(f[y] for y in q) 的话就太慢了
// 注意到这个式子与 x 没有任何关系，即遍历当前层时，对于当前层的每个 x，对应的 sum(f[y] for y in q) 是一样的
// 我们在求 f[x] 的时候，可以顺便用一个变量 s 记录这一层的 f[x] 的和，这样遍历下一行的时候就不需要重新遍历
// 更新 q 的时候顺便更新 pre_s = s 保存 s，这样保证 s 能存储下一行，上一行的和就是 pre_s

// DP 初始化
// 根据题目，f[1] = 1，f[x for x in g[1]] = 1，pre_s = len(g[1])
// 答案就是到达所有节点的路径总和，sum(f)，记得取模，注意模数不是 1e9 + 7

// 复杂度分析：图上遍历的复杂度 O(n + m)，其中 m 为边数
// 对于树来说，m = n - 1，O(n + n - 1) = O(2*n-1)，O 记号下为 O(n)

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 998244353;

void solve() {
    int n; cin >> n;
    vector<vector<int>> g(n + 1);
    for (int i = 2; i <= n; ++i) {
        int fa; cin >> fa;
        g[fa].push_back(i);
    }
    ll pre_s = g[1].size();
    vector<ll> f(n + 1);
    vector<int> q;
    f[1] = 1;
    for (int x : g[1]) {
        f[x] = 1;
        q.push_back(x);
    }
    vector<int> cur;
    while (!q.empty()) {
        ll s = 0;
        for (int fa : q) {
            for (int x : g[fa]) {
                cur.push_back(x);
                f[x] = (pre_s - f[fa] + MOD) % MOD;
                s = (s + f[x]) % MOD;
            }
        }
        q = cur;
        cur.clear();
        pre_s = s;
    }
    ll ans = 0;
    for (int i = 1; i <= n; ++i) {
        ans = (ans + f[i]) % MOD;
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