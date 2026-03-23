// https://cpc.zufe.ac.cn/contest/7/sub/5/problem/F
// 构造题

// 记录每个数出现的次数
// 如果只有一个 0，则构造 0
// 如果长度为偶数但是有一个数出现次数为奇数，则无法构造
// 如果长度为奇数但是有两个以上的数出现次数为奇数，则无法构造
// 然后倒序枚举 9 到 0 的出现次数，从两边开始填
// 注意，答案的长度与原数组一定相同
// 最终构造出来的数，如果开头为 0，因为提前特判了只有一个 0 的情况，此时直接返回 -1
// 最后构造的数就是答案

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
const int inf = 0x3f3f3f3f;
const ll ll_inf = 1e18;

void solve() {
    string s; cin >> s;
    int n = s.size();
    if (n == 1 && s[0] - '0' == 0) {
        cout << 0 << endl;
        return;
    }
    vector<int> cnt(10);
    for (char c : s) {
        cnt[c - '0']++;
    }
    vector<int> ans(n, -1);
    int l = 0, r = n - 1;
    for (int x = 9; x >= 0; --x) {
        if (cnt[x] == 0) continue;
        if (cnt[x] % 2 && ((n & 1) == 0 || ans[n / 2] >= 0)) {
            cout << -1 << endl;
            return;
        }

        if (cnt[x] % 2) {
            ans[n / 2] = x;
            cnt[x] -= 1;
        }
        while (cnt[x] > 0) {
            ans[l++] = x;
            ans[r--] = x;
            cnt[x] -= 2;
        }
    }
    if (!ans[0]) {
        cout << -1 << endl;
        return;
    }
    string res;
    for (int x : ans) {
        res += x + '0';
    }
    cout << res << endl;
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