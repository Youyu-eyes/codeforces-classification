// https://codeforces.com/problemset/problem/835/C
// 1600 大概 1400
// 二维前缀和

// 原亮度 p，时间 t 后对应的亮度为 (p + t) % (c + 1)
// 问题就变成了在 (x1, y1) 到 (x2, y2) 的矩阵里，原亮度为 p 的恒星个数 * ((p + t) % (c + 1))
// 在一个矩阵里 O(1) 的求元素和，用二维前缀和
// 二位前缀和，见 leetcode 灵茶山艾府 数据结构题单

// 值得注意的是，本题没有明说，但是星星是可以重叠的，比如 (1, 1) 上可能有两颗恒星，所以不能只用 map<pair<x, y>, s> 存储星星的亮度
// 可以直接开一个 3 维数组，而且这样看前缀和的计算也更直观

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;

void solve() {
    int n, q, c; cin >> n >> q >> c;
    vector<vector<vector<int>>> mp(c + 1, vector<vector<int>>(101, vector<int>(101)));
    for (int i = 0; i < n; ++i) {
        int x, y, s; cin >> x >> y >> s;
        mp[s][x][y]++;
    }

    vector<vector<vector<int>>> f(c + 1, vector<vector<int>>(102, vector<int>(102)));
    for (int p = 0; p <= c; ++p) {
        for (int x = 0; x <= 100; ++x) {
            for (int y = 0; y <= 100; ++y) {
                f[p][x + 1][y + 1] = f[p][x][y + 1] + f[p][x + 1][y] - f[p][x][y] + mp[p][x][y];
            }
        }
    }
    
    for (; q > 0; q--) {
        int t, x1, y1, x2, y2; cin >> t >> x1 >> y1 >> x2 >> y2;
        int ans = 0;
        for (int p = 0; p <= c; ++p) {
            int light = (p + t) % (c + 1);
            ans += (f[p][x2 + 1][y2 + 1] - f[p][x1][y2 + 1] - f[p][x2 + 1][y1] + f[p][x1][y1]) * light;
        }
        cout << ans << endl;
    }
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);
    
    solve();
    
    return 0;
}