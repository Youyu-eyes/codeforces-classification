// https://codeforces.com/problemset/problem/1857/D
// 1300
// 式子变形 + 贪心

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;

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

    int mx = INT_MIN;
    vector<int> ans;
    for (int i = 0; i < n; ++i) {
        if (a[i] - b[i] > mx) {
            mx = a[i] - b[i];
            ans = {i + 1};
        }
        else if (mx == a[i] - b[i]) {
            ans.push_back(i + 1);
        }
    }

    cout << ans.size() << endl;
    for (int i = 0; i < ans.size(); ++i) {
        cout << ans[i] << ' ';
    }
    cout << '\n';
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