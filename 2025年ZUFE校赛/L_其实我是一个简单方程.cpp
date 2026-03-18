#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
const int inf = 0x3f3f3f3f;
const ll ll_inf = 1e18;

void solve() {
    int n; cin >> n;
    int mid = trunc(sqrt(n));
    if (mid * mid == n) {
        cout << mid << ' ' << 0 << endl;
        return;
    }
    if (n & 1) {
        cout << (1 + n) / 2 << ' ' << (n - 1) / 2 << endl;
        return;
    }
    if (n % 4 == 0) {
        cout << (2 + n / 2) / 2 << ' ' << (n / 2 - 2) / 2 << endl;
        return;
    }
    cout << -1 << endl;
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