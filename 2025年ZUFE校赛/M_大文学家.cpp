#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
const int inf = 0x3f3f3f3f;
const ll ll_inf = 1e18;

void solve() {
    int n, k; cin >> n >> k;
    vector<int> f(n + 1, inf);
    deque<int> q;
    f[1] = 0;
    q.push_back(1);
    for (int i = 2; i <= n; ++i) {
        if ((i & 1) == 0) { 
            f[i] = f[i / 2]; 
        }
        int left = max(i - k, (i + 1) / 2);
        while (!q.empty() && q.front() < left) {
            q.pop_front();
        }
        f[i] = min(f[i], f[q.front()]) + 1;
        while (!q.empty() && f[q.back()] >= f[i]) {
            q.pop_back();
        }
        q.push_back(i);
    }
    cout << f[n] << endl;
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