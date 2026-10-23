// https://codeforces.com/problemset/problem/2049/B
// 1300
// 贪心 + 数学
// 关键是尝试一下为什么不能构造

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;

void solve() {
    int n; cin >> n;
    string s; cin >> s;
    if (s[0] == 's') s[0] = '.';
    if (s.back() == 'p') s.back() = '.';
    optional<char> target;
    bool flag = true;
    for (char c : s) {
        if (target) {
            if (c != target && c != '.') {
                flag = false;
                break;
            }
        }
        else if (c != '.') {
            target = c;
        }
    }
    if (flag) {
        cout << "YES" << endl;
    }
    else {
        cout << "NO" << endl;
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