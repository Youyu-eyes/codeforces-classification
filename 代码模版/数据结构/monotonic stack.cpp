#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
const int inf = INT_MAX;
const ll ll_inf = LLONG_MAX;

pair<vector<int>, vector<int>> nearestGreater(vector<int>& nums) {
    int n = nums.size();
    // left[i] 是 nums[i] 左侧最近的严格大于 nums[i] 的数的下标，若不存在则为 -1
    vector<int> left(n);
    vector<int> st{-1}; // 哨兵
    for (int i = 0; i < n; i++) {
        int x = nums[i];
        while (st.size() > 1 && nums[st.back()] <= x) { // 如果求严格小于，改成 >=
            st.pop_back();
        }
        left[i] = st.back();
        st.push_back(i);
    }

    // right[i] 是 nums[i] 右侧最近的严格大于 nums[i] 的数的下标，若不存在则为 n
    vector<int> right(n);
    st = {n}; // 哨兵
    for (int i = n - 1; i >= 0; i--) {
        int x = nums[i];
        while (st.size() > 1 && nums[st.back()] <= x) { // 如果求严格小于，改成 >=
            st.pop_back();
        }
        right[i] = st.back();
        st.push_back(i);
    }

    return {left, right};
}

// 在 i 左侧离 i 最远的 下标 j
int longestWPI(vector<int> &hours) {
    int n = hours.size(), ans = 0, s[n + 1]; // 前缀和
    stack<int> st;
    st.push(s[0] = 0);
    for (int j = 1; j <= n; ++j) {
        s[j] = s[j - 1] + (hours[j - 1] > 8 ? 1 : -1);
        if (s[j] < s[st.top()]) st.push(j); // 感兴趣的 j
    }
    for (int i = n; i; --i)
        while (!st.empty() && s[i] > s[st.top()]) {
            ans = max(ans, i - st.top()); // [栈顶,i) 可能是最长子数组
            st.pop();
        }
    return ans;
}