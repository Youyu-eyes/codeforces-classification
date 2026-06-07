// https://codeforces.com/problemset/problem/2167/F
// 1600
// 换根DP

// 先思考一个简单的问题，如果这颗树一开始以 1 为根，怎么求 S1
// 如果暴力枚举 k 个节点，一共需要枚举 C(k, n) 显然超时
// 思考一下，如果想要这 k 个节点的 LCA 是根节点，这 k 个节点要满足什么性质
// 显然，假设 1 的孩子节点为 2 3 4，分别对应一个子树，要使 k 个节点的 LCA 是根节点的必要条件是 k 个节点分别分布在其中的至少两个子树里
// 比如 2 子树选 1 个节点，3 子树选 k - 1 个节点，这 k 个节点的 LCA 依旧是 1
// 思考这个命题的对应命题，要使 k 个节点的 LCA 不是根节点的必要条件是什么
// 根据上面的推论，要使 k 个节点的 LCA 不是根节点，则这 k 个节点需要满足全部在一颗子树上
// 比如全在 2 子树上，则这 k 个节点的 LCA 就是 2，而非整颗树根节点
// 要使 k 个节点全部分布在 2 这颗子树上，显然需要满足 2 这颗子树的大小 >= k
// 定义在全树以 1 为根节点的情况下，以 x 节点为根的子树的大小为 sz[x]
// 则 x 可以作为 LCA 的必要条件为 sz[x] >= k
// 还是 1 -> 2 3 4 这颗树，整颗树的 S1 是以 2 3 4 为根的子树的 LCA 个数加上根节点
// 定义以 1 为根节点的情况下，以 x 节点为根的子树的 LCA 个数为 dfs(x)
// 上述文字可以表述为 dfs(1) = dfs(2) + dfs(3) + dfs(4) + 1
// 对于这个根节点能不能给答案作贡献，我们需要讨论 sz[x] >= k，sz[1] = n，题目确定 n >= k，因此 1 节点一定能做一次贡献
// 对于 1 的每一个子树，问题就变成了这颗子树中有多少的子树的根节点可以作为 LCA，这是一个规模更小的子问题，可以递归解决
// 而对于一般的情况，假设 2 的孩子节点列表为 g[2]，dfs(2) = sum(dfs(y) for y in g[2]) + (sz[2] >= k)
// 更一般的情况，假设 x 的孩子节点列表为 g[x]，dfs(x) = sum(dfs(y) for y in g[x]) + (sz[x] >= k)
// 注意，这个题看似能提前判断 sz[x] < k 的时候推出递归，不影响 f 数组的答案
// 但是我们求 sz[x] 势必需要递归子树 y，用子树的 size 和求出 sz[x]

// 定义 f[x] 为以 x 为根节点，整颗子树的 LCA，即 SX
// 显然 f[1] = S[1]
// 问题变成了，我们求出了 f[1]，怎么 O(1) 地更新其他的 f 值
// 这是很经典的换根DP，rerooting
// 在这里我很难画图，请读者画出这样一棵树
// g = [[], [2, 3], [4, 5]], k = 3
// 差不多是这样一棵树

//    1
//  2   3
// 4 5

// 我们很容易看出 f[1] = 2
// 换根为 2
//    2
// 4  5  1
//         3

// 显然，答案为 1
// 先思考 2 这个节点对答案的贡献
// 对于 f[1]，2 节点贡献了 1，因为子树 2 的大小 sz[2] >= 3
// 换根到 2 时，2 节点作为整颗树的根节点，一定会对答案做出贡献 1，因此前后贡献不变

// 如果换根到 3
//       3
//     1
//   2
//  4 5

// 原来的 3 节点对答案贡献是 0，因为 sz[3] < 3
// 而以 3 为根的树中，由于根节点一定会对答案做出贡献 1，因此整颗树的 LCA 节点相比 f[x] 多了 1
// 综合来看，从 x 换到以 y 的节点，y 对 f[y] 的贡献与 f[x] 的贡献的差值为 sz[y] < k

// 现在考虑 1 对答案的贡献的差值
// 对于 f[1]，1 节点一定会贡献 1，这是因为 1 是根节点的缘故
// 换完根后，1 节点对应子树的大小为 n - sz[2]
// 如果这颗子树的大小 n - sz[2] < k 的话，1 就不能作为 LCA，即对答案 f[2] 的贡献为 0
// 之前贡献是 1，现在贡献为 0，前后差了 1
// 对于一般的情况，换到以 y 为根后，1 的子树大小为 n - sz[y]
// 1 原来的贡献为 1，如果 n - sz[y] < k，1 对 f[y] 贡献为 0，需要减去

// 可以证明，换根只与 x 和 y 节点有关，其余子树对 f[x] 和 f[y] 的贡献不变

// 综上，f[y] = f[x] + (sz[y] < k)(y 是否会对 f[y] 做出额外贡献) - ((n - sz[y]) < k)(x 是否在 y 为根的树中不能作为 LCA)
// 可以跑两个 DFS，也可以跑 DFS + BFS，个人觉得递归更好写

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;

void solve() {
    int n, k; cin >> n >> k;
    vector<vector<int>> g(n);
    for (int i = 0; i < n - 1; ++i) {
        int u, v; cin >> u >> v; --u, --v;
        g[u].push_back(v);
        g[v].push_back(u);
    }
    
    vector<int> sz(n);
    auto dfs = [&](this auto&& dfs, int x, int fa) -> pair<int, int> {
        int size = 0;
        int mask = 0;
        for (int y : g[x]) {
            if (y != fa) {
                pair<int, int> p = dfs(y, x);
                size += p.first;
                mask += p.second;
            }
        }
        sz[x] = size + 1;
        mask += (sz[x] >= k);
        return pair(sz[x], mask);
    };
    vector<ll> f(n);
    f[0] = dfs(0, -1).second;

    ll ans = f[0];
    auto rerooting = [&](this auto&& rerooting, int x, int fa) -> void {
        for (int y : g[x]) {
            if (y != fa) {
                f[y] = f[x] + (sz[y] < k) - ((n - sz[y]) < k);
                ans += f[y];
                rerooting(y, x);
            }
        }  
    };
    rerooting(0, -1);

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