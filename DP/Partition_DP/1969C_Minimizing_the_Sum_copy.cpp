// https://codeforces.com/problemset/problem/1969/C
// 1700
// 划分型DP

// 本题关键是思考两个问题
// 首先，一个数只可能被覆盖一次，因为如果相邻的数能使这个数变得更小，那么会用更小的数更新，而不会重复更新
// 其次，假设一个 m 属于 [j, i]，且 m 是这个区间内最小数的下标
// 那么从 m 更新右边一个数需要 1 步，更新右边的右边又需要一步，以此类推
// 结论：从 m 更新 [j, i] 区间内所有数需要 i - j - 1 步

// 对于划分型DP，依旧定义 f[i][t] 为前 i 个数进行 t 步可以达到的最小和
// 对于 i 考虑选或不选
// 不选的话：f[i][t] = f[i - 1][t] + nums[i - 1]
// 选：考虑枚举选哪个
// 假设 j 下标 j < i，对于 [j, i] 中的数，用区间最小值更新所有数，需要 i - j - 1 步
// 剩余 t - (i - j - 1) 步，区间和为 (i - j) * min{i, j}，其中 min{i, j} 表示 nums 从 j 到 i - 1 (左闭右开)区间最小值
// j 的下界：需要保证 t - (i - j - 1) >= 0 => 解得 j >= i - t - 1，上界为 i - 2，因为 i - 1 就不需要操作了
// f[i][t] = min{j = [i - t - 1, i - 2]}(f[j][t - (i - j - 1)] + (i - j) * min{i, j})
// 补一句，注意乘法溢出
// 最后两者取 min

// 由于 t <= k <= 10，枚举 j 的开销很小，时间复杂度 O(n*k^2)
// min{i, j} 可用 ST 表 O(1) 维护
// 预处理 O(nlogn)
// 总复杂度 O(nlogn + n*k^2)，实际上 O(nlogn) 才是瓶颈

// 补一句，如果用刷表法应该可以用单调队列优化掉枚举 j 的过程
// 作者没写出来，交给读者思考

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
const int inf = INT_MAX;
const ll ll_inf = 1e18;

class SparseTable {
    vector<vector<int>> st_min;
    vector<vector<int>> st_max;

public:
    // 时间复杂度 O(n * log n)
    SparseTable(const vector<int>& nums) {
        int n = nums.size();
        int w = bit_width((uint32_t) n);
        st_min.resize(w, vector<int>(n));
        st_max.resize(w, vector<int>(n));

        for (int j = 0; j < n; j++) {
            st_min[0][j] = nums[j];
            st_max[0][j] = nums[j];
        }

        for (int i = 1; i < w; i++) {
            for (int j = 0; j + (1 << i) <= n; j++) {
                st_min[i][j] = min(st_min[i - 1][j], st_min[i - 1][j + (1 << (i - 1))]);
                st_max[i][j] = max(st_max[i - 1][j], st_max[i - 1][j + (1 << (i - 1))]);
            }
        }
    }

    // [l, r) 左闭右开，下标从 0 开始
    // 必须保证 l < r
    // 时间复杂度 O(1)
    int query_min(int l, int r) const {
        int k = bit_width((uint32_t) r - l) - 1;
        return min(st_min[k][l], st_min[k][r - (1 << k)]);
    }

    // [l, r) 左闭右开，下标从 0 开始
    // 必须保证 l < r
    // 时间复杂度 O(1)
    int query_max(int l, int r) const {
        int k = bit_width((uint32_t) r - l) - 1;
        return max(st_max[k][l], st_max[k][r - (1 << k)]);
    }
};

void solve() {
    int n, k; cin >> n >> k;
    vector<int> nums(n);
    for (auto& x : nums) cin >> x;

    vector<vector<ll>> f(n + 1, vector<ll>(k + 1, ll_inf));
    
    SparseTable st(nums);
    ll s = nums[0];

    if (k >= n - 1) {
        cout << 1LL * st.query_min(0, n) * n << endl;
        return;
    }

    f[0][0] = 0;
    for (int i = 1; i <= k; ++i) {
        f[i][0] = s;
        s += nums[i];
        for (int t = i; t <= k; ++t) {
            f[i][t] = 1LL * st.query_min(0, i) * i;
        }
    }

    for (int i = 0; i <= n; ++i) {
        for (int t = 0; t <= k && t < i; ++t) {
            f[i][t] = f[i - 1][t] + nums[i - 1];
            for (int j = i - t - 1; j <= i - 2; ++j) {
                f[i][t] = min(f[i][t], f[j][t - (i - j - 1)] + 1LL * (i - j) * st.query_min(j, i));
            }
        }
    }

    ll ans = ll_inf;
    for (int j = 0; j <= k; ++j) {
        ans = min(ans, f[n][j]);
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