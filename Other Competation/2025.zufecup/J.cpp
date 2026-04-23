#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
const int inf = 0x3f3f3f3f;
const ll ll_inf = 1e18;

void solve() {
    int n; cin >> n;
    ll pos = 0;
    ll neg = 0;
    int m = 0;
    for (int i = 0; i < n; ++i) {
        int wt; cin >> wt;
        if (wt < 0) {
            neg += wt;
            m++;
        } else {
            pos += wt;
        }
    }
    if (m == 0) {
        cout << pos << endl;
        return;
    }
    double ans = (double)pos / (m + 1) + (double)neg / m;
    cout << fixed << setprecision(10) << ans << endl;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);
    solve();
    return 0;
}