// https://codeforces.com/problemset/problem/1151/B
// 1600
// 异或性质 + 贪心

// 假设我们都取第一个元素 grid[i][0]，求得异或和为 ans
// 如果 ans != 0，直接输出
// 如果 ans == 0，则枚举每一行，假设 x = grid[i][j] != grid[i][0]
// ans = 0 ^ grid[i][0] ^ x = grid[i][0] ^ x
// 因为 grid[i][0] != x
// 所以 grid[i][0] ^ x != 0
// 如果一行元素中有至少一个元素与 grid[i][0] 不同，就能构造出答案

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;

void solve() {
    int m, n; cin >> m >> n;
    vector<vector<int>> grid(m, vector<int>(n));
    int ans = 0;
    for (int i = 0; i < m; ++i) {
        for (int j = 0; j < n; ++j) {
            cin >> grid[i][j];
            if (!j) {
                ans ^= grid[i][0];
            }
        }
    }

    if (ans) {
        cout << "TAK" << endl;
        for (int i = 0; i < m; ++i) {
            cout << 1 << ' '; 
        }
        cout << endl;
        return;
    }

    int x = -1, y = -1;
    for (int i = 0; i < m; ++i) {
        for (int j = 1; j < n; ++j) {
            if (grid[i][j] != grid[i][0]) {
                x = i;
                y = j;
                break;
            }
        }
        if (x >= 0) break;
    }

    if (x >= 0) {
        cout << "TAK" << endl;
        for (int i = 0; i < m; ++i) {
            if (i == x) {
                cout << y + 1 << ' ';
            } else {
                cout << 1 << ' ';
            }
        }
        return;
    } else {
        cout << "NIE" << endl;
    }
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);
    
    solve();
    
    return 0;
}