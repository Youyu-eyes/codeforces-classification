// https://codeforces.com/problemset/problem/165/C
// 1600 久远的题 预期 1300
// 恰好型滑动窗口

// 恰好有 k 个 1 的子数组数量
// 转化成至少 k 个 1 子数组数量 - 至少 k + 1 个 1 子数组数量
// 比如 k = 3，至少 k 个 1 的子数组数量会计算出有 3 4 5 …… s.count('1') 个 1 的子数组数量
// 至少 k + 1 个 1 的子数组数量会计算出有 4 5 6 …… s.count('1') 个 1 的子数组数量
// 二者相减 即为有 3 个 1 的子数组数量

// 至少 k 个 1 子数组数量，越长越合法型滑窗
// 包装一个滑窗函数 cal，传入 k 和 k + 1，就不用重复写两边代码

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;

void solve() {
    int k; cin >> k;
    string s; cin >> s;
    auto cal = [&](int limit) -> ll {
        int cnt = 0;
        int l = 0;
        ll ans = 0;
        for (int i = 0; i < s.size(); ++i) {
            cnt += s[i] - '0';
            
            while (l <= i && cnt >= limit) {
                cnt -= s[l] - '0';
                l++;
            }

            ans += l;
        }
        return ans;
    };
    cout << cal(k) - cal(k + 1) << endl;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);
    
    solve();
    
    return 0;
}