// https://codeforces.com/contest/1196/problem/D1
// 1500 暴力做法估计 1200
// 定长滑窗

// 本题和 D2 是一样的，请看下一个题 1196D2

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
const int inf = INT_MAX;

const string t = "RGB";

void solve() {
    int n, k; cin >> n >> k;
    string s; cin >> s;
    int ans = inf;
    for (int d = 0; d < 3; ++d) {
        int cnt = 0;
        for (int i = 0; i < n; ++i) {
            cnt += (s[i] != t[(i + d) % 3]);
            int l = i - k + 1;
            if (l < 0) {
                continue;
            }
            ans = min(ans, cnt);
            cnt -= (s[l] != t[(l + d) % 3]);
        }
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