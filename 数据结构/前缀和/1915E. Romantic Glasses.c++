// https://codeforces.com/problemset/problem/1915/E
//  1300
//  前缀和 + 哈希表 + 枚举右维护左
//  C++ 版本，
// 注意此题 哈希表 unorder_map 过不了，会哈希冲突，必须用红黑树

#include <bits/stdc++.h>
using namespace std;
#define ll long long

void solve() {
    int n;
    cin >> n;
    vector<int> a(n);
    for (int i = 0; i < n; ++i) {
        cin >> a[i];
    }

    map<ll, ll> cnt;
    cnt[0] = 1;
    ll s = 0;
    bool flag = false;
    for (int i = 0; i < n; ++i) {
        s += (i & 1) ? a[i] : -a[i];
        if (cnt[s]) {
            flag = true;
            break;
        }
        cnt[s] = 1;
    }
    if (flag) {
        cout << "YES\n";
    }
    else {
        cout << "NO\n";
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