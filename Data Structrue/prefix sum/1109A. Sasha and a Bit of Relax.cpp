// https://codeforces.com/problemset/problem/1109/A
// 1600
// 很古老的题，放到现在就 1300
// 枚举右维护左 + 异或前缀和

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;

void solve() {
    int n; cin >> n;
    ll s = 0, ans = 0;
    unordered_map<ll, array<ll, 2>> cnt;
    cnt[0][1] = 1;
    for (int i = 0; i < n; ++i) {
        int a; cin >> a;
        s ^= a;
        if (cnt.contains(s)) {
            ans += cnt[s][i % 2];
        }
        cnt[s][i % 2]++;
    }
    cout << ans << endl;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);
    
    solve();
    
    return 0;
}