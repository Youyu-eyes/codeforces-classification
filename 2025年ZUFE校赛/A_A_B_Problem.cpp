// https://cpc.zufe.ac.cn/contest/7/sub/5/problem/A
// 模拟即可

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
const int inf = INT_MAX;
const ll ll_inf = 1e18;

void solve() {
    int n; cin >> n;
    vector<int> a(n);
    vector<int> b(n);
    for (int i = 0; i < n; ++i) {
        cin >> a[i];
    }
    for (int i = 0; i < n; ++i) {
        cin >> b[i];
    }
    ll ans = 0;
    for (int i = 0; i < n; ++i) {
        if (a[i] >= b[i]) {
            ans += abs(a[i] + b[i] - a[i] * b[i]);
        }
    }
    cout << ans << endl;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);
    
    solve();
    
    return 0;
}