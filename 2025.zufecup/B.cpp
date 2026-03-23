// https://cpc.zufe.ac.cn/contest/7/sub/5/problem/B
// 单调栈
// min(nums[l], nums[r]) < all(nums[l + 1, r - 1])
// 形象的说，像一个碗，如果两端比中间都大，就能盛水
// 分类讨论
// 假设 min(nums[l], nums[r]) = nums[r]
// 对于每一个 i，找到 i 左端第一个大于等于 nums[i] 的数的下标 j，则在 nums[j, i] 之间的数都满足小于 nums[i] 的条件
// 而对于 j 左端的数，假设下标 k，在 nums[k, i] 的区间内，有 nums[j] > nums[i] 不符合题意
// 因此对于 i，j 是 i 左端点的最大答案
// 同理，假设 min(nums[l], nums[r]) = nums[l]
// 则找右端点第一个大于等于 nums[i] 的下标

// 代码实现时可以把正向的单调栈封装成函数，然后将 nums 反转后再调用

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
const int inf = 0x3f3f3f3f;
const ll ll_inf = 1e18;

void solve() {
    int n; cin >> n;
    vector<int> nums(n);
    for (int i = 0; i < n; ++i) {
        cin >> nums[i];
    }

    auto f = [&](this auto&& f, vector<int> nums) -> int {
        int res = 0;
        stack<int> st;
        for (int i = 0; i < n; ++i) {
            while (!st.empty() && nums[st.top()] < nums[i]) {
                st.pop();
            }
            if (!st.empty()) {
                res = max(res, i - st.top() + 1);
            }
            st.push(i);
        }
        return res;
    };

    int ans = 0;
    ans = max(ans, f(nums));
    reverse(nums.begin(), nums.end());
    ans = max(ans, f(nums));
    cout << ans << endl;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);
    
    solve();
    
    return 0;
}