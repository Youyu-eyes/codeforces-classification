// https://acm.timus.ru/problem.aspx?space=1&num=1401
// 递归分治

// 分治法经典题，结论也可以直接记住
// 在 2^n * 2^n 的正方形网格内，将任意的格子置为空，剩下的格子一定能被形如
// xx  xx  x    x
// x    x  xx  xx
// 这样的 L 形图形填满

// 递归分治的过程就是该结论的构造性证明

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;
const int inf = 0x3f3f3f3f;
const ll ll_inf = 1e18;

template<typename T = long long>
T myPow(long long x, long long n, long long mod = 0) {
    long long ans = 1;
    long long base = x;
    long long exp = n;

    // 如果模数非零，先对底数取模
    if (mod) base %= mod;

    // 处理负指数
    if (exp < 0) {
        if (!mod) return T(0);   // 无模数模式下负指数无意义，返回0
        // 模质数下求逆元（费马小定理）
        long long inv = 1, b = base, p = mod - 2;
        while (p) {
            if (p & 1) inv = (inv * b) % mod;
            b = (b * b) % mod;
            p >>= 1;
        }
        base = inv;
        exp = -exp;
    }

    while (exp) {
        if (exp & 1) {
            ans *= base;
            if (mod) ans %= mod;
        }
        base *= base;
        if (mod) base %= mod;
        exp >>= 1;
    }
    return T(ans);
}

void solve() {
    int n; cin >> n;
    int pos_i, pos_j; cin >> pos_i >> pos_j; --pos_i; --pos_j;
    int L = myPow<int>(2, n);
    vector ans(L, vector<int>(L));
    int cur = 0;
    auto dfs = [&](this auto&& dfs, int i, int j, int x, int y, int len) {
        if (len == 1) return;

        int half = len / 2;
        vector<pair<int, int>> certen(4);
        certen[0] = {i + half - 1, j + half - 1};
        certen[1] = {i + half - 1, j + half};
        certen[2] = {i + half, j + half - 1};
        certen[3] = {i + half, j + half};

        vector<pair<int, int>> start(4);
        start[0] = {i, j};
        start[1] = {i, j + half};
        start[2] = {i + half, j};
        start[3] = {i + half, j + half};
        
        cur++;
        int id = cur;
        for (int p = 0; p < 4; ++p) {
            int r = start[p].first;
            int c = start[p].second;
            if (r <= x && x < r + half && c <= y && y < c + half) {
                dfs(r, c, x, y, half);
                continue;
            }
            int mx = certen[p].first;
            int my = certen[p].second;
            ans[mx][my] = id;
            dfs(r, c, mx, my, half);
        }
    };

    dfs(0, 0, pos_i, pos_j, L);
    for (int i = 0; i < L; ++i) {
        for (int j = 0; j < L; ++j) {
            cout << ans[i][j] << ' ';
        }
        cout << endl;
    }
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);
    
    solve();
    
    return 0;
}