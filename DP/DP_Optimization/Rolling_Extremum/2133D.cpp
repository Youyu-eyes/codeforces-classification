// https://codeforces.com/contest/2133/problem/D
// 1900
// 贪心 + 划分型DP
// 我会从 O(n^2) 的DP，优化到 O(n)，再进一步带你理解官解的答案
// 对于官解的解释我不认可，我会从另一个角度带你理解

// 首先，我们需要确定一下划分的方向
// 思考 (0 < j < i < n and i - j > 1) 的两个数 h[j] 与 h[i]
// 如果先删除 j 再删除 i
// 首先受到伤害的是 h[j + 1]，受到 j + 1 点伤害
// 然后再是 h[i + 1]，此时底部变成了原来的 h[j + 1]， 总高度为 i - j
// 总共是 i + 1 点伤害

// 再思考先删除 i 再删除 j
// 首先受到伤害的是 h[i + 1]，受到 i + 1 点伤害
// 其次是 h[j + 1]，受到 j + 1 点伤害
// 总共 i + j + 2 点伤害

// 发现差别在 h[i + 1] 上
// 分类讨论，令 a1 为先删 j 后删 i 的摔落伤害，a2 为先删 i 后删 j 的摔落伤害，显然有 i - j < i + 1
// if   h[i + 1] >= i + 1:        a1 = i - j, a2 = i + 1    => a1 < a2
// elif i - j < h[i + 1] < i + 1: a1 = i - j, a2 = h[i + 1] => a1 < a2
// else h[i + 1] <= i - j:        a1 = a2 = h[i + 1]        => a1 = a2

// 根据严格证明，在不考虑特殊情况的条件下，先杀上面的怪物一定比先杀下面的怪物更优
// 但是，有个例外：如果 j = i - 1，先杀 j 更优，因为此时 h[j + 1] = h[i]
// 如果先杀 i，再杀 j 时，没有 h[i] 掉落，没有 j + 1 伤害的相抵，前面的分类讨论也就没有意义

// 注意到，如果我们从上往下划分，划到 j 时，假设上一次是 i，那么新产生的怪物堆就是 [j + 1, i - 1]
// 根据上述贪心理论，在确定挥刀斩掉 j 时，j 到 i 中不可能在 [j + 2, i - 1] 中挥出一刀
// 因为如果在 [j + 2, i - 1] 中挥出的一刀要优于 j 的话，应该先杀在这个区间里的怪物，再杀 j，这样才能使掉落伤害最大化
// 那么 [j + 1, i - 1] 区间内的怪物，我们就只能从 j + 1 开始一个一个向上删除

// 现在，整个过程分为 2 步，一步是先在 [0, i - 1] 的区间内划分出 [0, j - 1]，[j + 1, i - 1] 两个区间
// 然后从下往上删除 [j + 1, i - 1] 的怪物
// [0, j - 1] 的部分就是一个规模更小的子问题，可用动态规划解决
// 对于划分型DP，通用定义方式为定义 f[i] 为 [0, i - 1] 区间内，杀死怪物所需最少步骤

// 先算 [j + 1, i - 1] 中的开销
// 第 j + 1 个怪物受到掉落伤害，剩下血量为 max(h[j + 1] - (j + 1), 0)，需要单独减去
// [j + 2, i - 1] 的怪物，每个怪物剩下的血量为 h - 1
// 则所需开销为 sum(h[m] - 1 for m in range(j + 2, i)) = sum(h[m] for m in range(j + 2, i)) - (i - j - 2)
// 定义前缀和 s[i + 1] = s[i] + h[i]，sum(h[m] for m in range(j + 2, i)) = s[i] - s[j + 2]
// 显然，j + 2 > i 时没有意义，因此 j <= i - 2

// 对于 f[i]，要先杀 h[j]，然后杀 [j + 1, i - 1]，最后杀 [0, j - 1]
// 状态转移方程为 f[i] = min{for j in range(i - 1)}(f[j] + h[j] + max(h[j + 1] - (j + 1), 0) + s[i] - s[j + 2] - i + j + 2)
// => f[i] = min{for j in range(i - 1)}(f[j] + h[j] + max(h[j + 1] - (j + 1), 0) - s[j] + j + 2) + s[i] - i
// 对于 j = i - 1，删除 j 后没有掉落伤害，f[i] = f[i - 1] + h[i - 1]
// 令上面的 f[i] = mn，两者再取 min
// f[i] = min(mn, f[i - 1] + h[i - 1])
// 由于枚举 j 的复杂度是 O(n)，但是所求是从 0 开始的连续区间最小值，可用一个单独的变量优化到 O(n)
// 最终答案为 f[n]
// 初始化 f[0] = 0，f[1] = h[0]（最底下的数只能删除，不能享受摔落伤害）

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
const int inf = INT_MAX;
const ll ll_inf = LLONG_MAX;

void solve() {
    int n; cin >> n;
    vector<int> h(n);
    vector<ll> s(n + 1);
    for (int i = 0; i < n; ++i) {
        cin >> h[i];
        s[i + 1] = s[i] + h[i];
    }
    vector<ll> f(n + 1, ll_inf);
    f[0] = 0;
    f[1] = h[0];
    ll mn = f[0] + h[0] + h[1] - 1 - s[2] + 2;
    for (int i = 2; i <= n; ++i) {
        f[i] = min(mn + s[i] - i, f[i - 1] + h[i - 1]);
        if (i < n) {
            mn = min(mn, f[i - 1] + h[i - 1] + max(h[i] - i, 0) - s[i + 1] + i + 1);
        }
    }
    cout << f[n] << endl;
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


// 下面是对官解的解释

// 根据上面的贪心证明，枚举 j 的过程其实是多余的
// 因为我们要在 [j + 1, i - 1] 区间里依次杀掉一个怪物，最后会删到 i - 2，而先杀掉 i - 2 比先杀掉下面的怪物都更优
// 因此如果直接在 i - 2 处删除，是满足 j < i - 1 的贪心最优解
// 还需额外考虑一下 j = i - 1 的情况，两个取 min
// 官解代码如下

void solve() {
    int n; cin >> n;
    vector<int> h(n);
    vector<ll> s(n + 1);
    for (int i = 0; i < n; ++i) {
        cin >> h[i];
        s[i + 1] = s[i] + h[i];
    }
    vector<ll> f(n + 1, ll_inf);
    f[0] = 0;
    f[1] = h[0];
    ll mn = f[0] + h[0] + h[1] - 1 - s[2] + 2;
    for (int i = 2; i <= n; ++i) {
        f[i] = min(f[i - 2] + h[i - 2] + max(h[i - 1] - (i - 1), 0), f[i - 1] + h[i - 1]);
    }
    cout << f[n] << endl;
}