// https://cpc.zufe.ac.cn/contest/7/sub/5/problem/M
// 划分型DP + 单调队列优化

// 反向思考
// 假设我们有一个长为 i 的蛋糕
// 如果 i 是偶数，我们可以切半，问题变成了长为 i / 2 的蛋糕所需分割最小次数
// 无论 i 是否为偶数，我们还可以执行的操作是切一刀，假设这刀切分 j - 1 和 j
// 切出来的蛋糕长度为 i - j，题目要求长度 <= k 且 剩下的长度 j >= i - j
// i - j <= k and 2 * j >= i
// j >= i - k and j >= (i + 1) // 2
// j >= max(i - k, (i + 1) // 2)
// 切分之后，问题就变成了长为 j 的蛋糕所需最小分割次数
// 问题就变成了一个规模更小的子问题，可用动态规划解决

// 对于划分型DP
// 定义 f[i] 为长为 i 的数组所需最小切分次数
// 这里我们假设的是从 0 到 i - 1 的数组，这样能与划分型DP的定义（划分点在剩余区间右端）契合
// 如果 i 是偶数，我们可以减半 i，上一次长度为 i // 2，f[i] = f[i // 2] + 1
// 注意这里，如果 i 不是偶数，我们不能减半划分，因为每一次的 i 都必须保证正整数
// 其余的情况，枚举 j，j 的范围为 [max(i - k, (i + 1) // 2), i - 1]
// 则 f[i] = min{j = [max(i - k, (i + 1) // 2), i - 1]}(f[j]) + 1
// 这是二维DP，复杂度 O(n^2)，超时

// 注意到 max(i - k, (i + 1) // 2) 两部分，由于 i 递增，k 为常数，因此 i - k 和 (i + 1) // 2 都递增，则 max(i - k, (i + 1) // 2) 递增
// 证明：max(a, b) 和 min(a, b) 在 a, b 递增的前提下一定递增
// 先证 max(a, b)
// 假设前一次是 max(a0, b0)，后一次是 max(a1, b1)，根据定义，a0 <= a1 且 b0 <= b1
// 假设 a0 < b0，剩下分为两种情况：a1 < b1, a1 > b1
// 1. a1 < b1 => max(a0, b0) = b0, max(a1, b1) = b1, b0 <= b1, 符合
// 2. a1 > b1 => max(a0, b0) = b0, max(a1, b1) = a1, b0 <= b1, b1 < a1 => b0 < a1，符合
// a0 > b0 时同理，根据对称性满足
// 后证 min(a, b)
// 假设前一次 min(a0, b0)，后一次 min(a1, b1)，根据定义 a0 <= a1, b0 <= b1
// 假设 a0 < b0
// 1. a1 < b1 => min(a0, b0) = a0, min(a1, b1) = a1, a0 <= a1, 符合
// 2. a1 > b1 => min(a0, b0) = a0, min(a1, b1) = b1, a0 < b0, b0 <= b1 => a0 < b1，符合
// 证毕

// 现在我们证明了随着右端点右移，左端点递增，求区间内 f[j] 的最大值，典型的滑动窗口最大值问题，可用单调队列优化
// 单调队列的介绍和学习见 b站灵茶山艾府 基础算法精讲

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
const int inf = 0x3f3f3f3f;
const ll ll_inf = 1e18;

void solve() {
    int n, k; cin >> n >> k;
    vector<int> f(n + 1, inf);
    deque<int> q;
    f[1] = 0;
    q.push_back(1);
    for (int i = 2; i <= n; ++i) {
        if ((i & 1) == 0) { 
            f[i] = f[i / 2]; 
        }
        int left = max(i - k, (i + 1) / 2);
        while (!q.empty() && q.front() < left) {
            q.pop_front();
        }
        f[i] = min(f[i], f[q.front()]) + 1;
        while (!q.empty() && f[q.back()] >= f[i]) {
            q.pop_back();
        }
        q.push_back(i);
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