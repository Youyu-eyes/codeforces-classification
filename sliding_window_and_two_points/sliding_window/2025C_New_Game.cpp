// https://codeforces.com/contest/2025/problem/C
// 1300
// 排序 + 滑窗
// 用 哈希表 统计可能会比较慢，可以参考官解，这里给出初学者易懂的写法

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;

void solve() {
    int n, k; cin >> n >> k;
    vector<int> a(n);
    for (int i = 0; i < n; ++i) {
        cin >> a[i];
    }

    ranges::sort(a);
    int ans = 0;
    unordered_map<int, int> cnt;

    int l = 0;
    for (int r = 0; r < n; ++r) {
        if (r > 0 && a[r] > a[r-1] + 1) {
            cnt.clear();
            l = r;
        }
        
        cnt[a[r]]++;

        while (cnt.size() > k) {
            cnt[a[l]]--;
            if (cnt[a[l]] == 0) {
                cnt.erase(a[l]);
            }
            l++;
        }
        
        ans = max(ans, r - l + 1);
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