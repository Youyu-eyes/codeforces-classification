// https://codeforces.com/problemset/problem/1107/G
// 2400
// 单调栈 + 前缀和 + ST表

// 数据结构综合题
// 假设在区间 [l, r] 中间一点 k
// max{i = [l, r - 1]}(diff[i]^2) = diff[k]^2
// 在这个区间内，left = [l, k - 1]，right = [k + 1, r - 1]
// 每一个 [left, right] 的区间的 max{i = [left, right - 1]}(diff[i]^2) = diff[k]^2
// 因此，我们枚举 k 为中心点，就能够确定 gap 的大小
// 而 sum{i = [left, right - 1]}(a - c[i])，可用前缀和优化成 s[right] - s[left]
// 在 k 确定，及 gap 确定时，要使 ans 最大，s[right] - s[left] 最大
// 则找 max{right = [k + 1, r - 1]}(s[right]) - min{left = [l, k - 1]}(s[left])
// 可用 st 表优化 RMQ 过程
// 最后一个问题是 l 和 r 的取值
// 好比一座山峰，k 为最高点，那 l 是左端第一个严格大于 diff[k] 的数，右端是第一个大于等于 diff[k] 的数
// 可用单调队列优化

// 非常综合，复杂，考察数据结构的离线处理

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
const int inf = 0x3f3f3f3f;
const ll ll_inf = 1e18;

pair<vector<int>, vector<int>> nearestGreater(vector<ll>& nums) {
    int n = nums.size();
    // left[i] 是 nums[i] 左侧最近的严格大于 nums[i] 的数的下标，若不存在则为 -1
    vector<int> left(n);
    vector<int> st{-1}; // 哨兵
    for (int i = 0; i < n; i++) {
        ll x = nums[i];
        while (st.size() > 1 && nums[st.back()] <= x) { // 如果求严格小于，改成 >=
            st.pop_back();
        }
        left[i] = st.back();
        st.push_back(i);
    }

    // right[i] 是 nums[i] 右侧最近的大于等于 nums[i] 的数的下标，若不存在则为 n
    vector<int> right(n);
    st = {n}; // 哨兵
    for (int i = n - 1; i >= 0; i--) {
        ll x = nums[i];
        while (st.size() > 1 && nums[st.back()] < x) { // 如果求严格小于，改成 >=
            st.pop_back();
        }
        right[i] = st.back();
        st.push_back(i);
    }

    return {left, right};
}

class SparseTable {
    vector<vector<ll>> st_min;
    vector<vector<ll>> st_max;

public:
    // 时间复杂度 O(n * log n)
    SparseTable(const vector<ll>& nums) {
        int n = nums.size();
        int w = bit_width((uint32_t) n);
        st_min.resize(w, vector<ll>(n));
        st_max.resize(w, vector<ll>(n));

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
    ll query_min(int l, int r) const {
        int k = bit_width((uint32_t) r - l) - 1;
        return min(st_min[k][l], st_min[k][r - (1 << k)]);
    }

    // [l, r) 左闭右开，下标从 0 开始
    // 必须保证 l < r
    // 时间复杂度 O(1)
    ll query_max(int l, int r) const {
        int k = bit_width((uint32_t) r - l) - 1;
        return max(st_max[k][l], st_max[k][r - (1 << k)]);
    }
};

void solve() {
    int n; cin >> n;
    ll a; cin >> a;
    ll pre_d = 0;
    vector<ll> s(n + 1, 0);
    vector<ll> diff(n - 1);
    ll ans = 0;
    for (int i = 0; i < n; ++i) {
        ll d, c; cin >> d >> c;
        ans = max(ans, a - c);
        s[i + 1] = s[i] + a - c;
        if (i > 0) diff[i - 1] = (d - pre_d) * (d - pre_d);
        pre_d = d;
    }

    auto [L, R] = nearestGreater(diff);
    SparseTable st(s);

    for (int k = 0; k < n - 1; ++k) {
        int left_min = max(0, L[k] + 1);
        int left_max = k;

        int r_max = (R[k] == n) ? n - 1 : R[k];
        int right_min = k + 2;
        int right_max = r_max + 1;

        if (left_min > left_max || right_min > right_max) continue;

        ll min_left = st.query_min(left_min, left_max + 1);
        ll max_right = st.query_max(right_min, right_max + 1);
        ans = max(ans, max_right - min_left - diff[k]);
    }

    cout << ans << "\n";
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    solve();
    
    return 0;
}