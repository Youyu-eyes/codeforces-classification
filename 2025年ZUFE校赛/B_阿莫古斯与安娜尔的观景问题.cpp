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