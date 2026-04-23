#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
const int inf = 0x3f3f3f3f;
const ll ll_inf = 1e18;
const int MX = 1 << 12;

void solve() {
    int n; cin >> n;
    vector<int> nums(n);
    vector<int> pre_xor(n + 1);
    unordered_map<int, int> cnt;
    cnt[0]++;
    for (int i = 0; i < n; ++i) {
        cin >> nums[i];
        pre_xor[i + 1] = pre_xor[i] ^ nums[i];
        cnt[pre_xor[i + 1]]++;
    }
    
    int ans = 0;
    for (int t = 0; t <= MX; ++t) {
        for (int x : pre_xor) {
            if (cnt[t ^ x] > 0) {
                if (t == 0 && cnt[x] < 2) continue;
                ans += 1;
                break;
            }
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