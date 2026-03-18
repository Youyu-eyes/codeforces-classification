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