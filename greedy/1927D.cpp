// https://codeforces.com/problemset/problem/1927/D
// 1300
// 贪心 + 构造
// 记录离 a[i] 最近的不同数的下标

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;



void solve() {
    int n; cin >> n;
    vector<int> a(n);
    vector<int> p(n, -1);
    for (int i = 0; i < n; ++i) {
        cin >> a[i];
        if (i > 0) {
            if (a[i - 1] == a[i]) {
                p[i] = p[i - 1];
            } else {
                p[i] = i - 1;
            }
        }
    }

    int m; cin >> m;
    for (int i = 0; i < m; ++i) {
        int l, r; cin >> l >> r;
        l--;
        r--;
        if (p[r] < l) {
            cout << -1 << ' ' << -1 << endl;
        } else {
            cout << p[r] + 1 << ' ' << r + 1 << endl;
        }
    }
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