// https://codeforces.com/problemset/problem/1575/L
// 2100
// 二维偏序 + 最长递增子序列

// 先来解释以下二维偏序
// 给定 N 个具有两个属性的元素 (x_i, y_i)，我们需要找到满足 x_i < x_j 且 y_i < y_j 的点对关系（具体关系看题目）
// 并在此基础上求最大长度（最长链）、数量统计（逆序对/支配对） 或 最值
// 简单来说就是在两个约束条件下寻找子序列或者下标对的问题

// 本题是经典的二维偏序求最长链问题
// 本题需要满足 3 个条件
// 1. i < j
// 2. a[i] < a[j]
// 3. i - a[i] <= j - a[j]

// 又由三角形不等式，可以得知在 2, 3 式满足的前提下，1 式一定满足
// 这里注意，1, 2 或 1, 3 都是推不出剩下的那条式子的
// 因此，转化为二维偏序问题

// 题目转化为满足 2, 3 式的最长子序列
// 定义 nums[i] = pair(a[i], i - a[i])
// 将 nums 数组用 2 式进行排序后
// 2 式一定满足要求，则可对 3 式求最长非递减子序列
// 但是对于本题，相同的 a[i] 只能放在一个 i 中，不能选择多个 a[i]
// 因此对 2 式排序的同时，如果 a[i] = a[j]，则需要 i - a[i] > j - a[j]
// 这样在求非递减子序列时，就不会选择相同的 a[i] 放在 i 下标处

// 时间复杂度 O(nlogn)
// 空间复杂度 O(n)

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
const int inf = 0x3f3f3f3f;
const ll ll_inf = 1e18;

void solve() {
    int n; cin >> n;
    vector<pair<int, int>> nums = {{0, 0}};
    for (int i = 1; i <= n; ++i) {
        int x; cin >> x;
        if (i >= x) {
            nums.push_back({x, i - x});
        }
    }
    ranges::sort(nums, {}, [](auto& e) { return pair(e.first, -e.second); });
    vector<int> g;
    for (auto [_, x] : nums) {
        auto it = ranges::lower_bound(g, x + 1);
        if (it == g.end()) {
            g.push_back(x);
        } else {
            *it = x;
        }
    }
    int ans = g.size();
    cout << ans - 1 << endl;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    solve();
    
    return 0;
}