// https://codeforces.com/problemset/problem/2184/D
// 1600
// 贪心 + 拆位法 + 数位DP
// 官解是组合数学，给出不同做法，数位DP 更简单

// 从二进制角度考虑，题意就是 Alice 将一个二进制数转换为一位 0 的最小操作次数
// 贪心的想，Alice 最优操作方式是如果是偶数就 除2，如果是奇数就 减1
// x / 2 < x - 1 在 x >= 2 时恒成立，而 x = 1 是奇数，只能 减1

// 对于一个二进制数
// 如果末尾是 0，则 除2 就能去掉末尾的 0，共 1 步
// 如果末尾是 1，则执行 减1 后 除2 能去掉末尾的 1，共 2 步
// 特殊情况：只剩一个 1，只需要执行 减1 一次操作

// 因此对于 n 的二进制数来说，总操作次数为 cnt0 + 2 * cnt1 - 1（开头的 1 多算了一次操作次数）
// 在 数位DP 过程中记录 cnt = cnt0 + cnt1 * 2，最后判断边界时 cnt - 1 > target，则数字合法

// 优化：n <= 10 ** 9 < 2 ** 30，最多 30 位，如果每一位都是 1，最多 60 次操作
// 如果 k >= 60，一定能把 n 置为 0，提前判断

#include <bits/stdc++.h>
using namespace std;

using ll = long long;
const int MOD = 1'000'000'007;

long long digitDP(long long low, long long high, int target) {
    string low_s; for (int x = low; x; x >>= 1) low_s.push_back('0' + (x & 1)); reverse(low_s.begin(), low_s.end()); if (low == 0) low_s = "0";
    string high_s; for (int x = high; x; x >>= 1) high_s.push_back('0' + (x & 1)); reverse(high_s.begin(), high_s.end()); if (high == 0) high_s = "0";

    int n = high_s.size();
    int diff_lh = n - low_s.size();
    vector memo(n, vector<long long>(2 * (n - 1), -1));

    auto dfs = [&](this auto&& dfs, int i, int cnt, bool limit_low, bool limit_high) -> long long {
        if (i == n) {
            return cnt - 1 > target;
        }

        if (!limit_low && !limit_high && memo[i][cnt] >= 0) {
            return memo[i][cnt];
        }

        int lo = limit_low && i >= diff_lh ? low_s[i - diff_lh] - '0' : 0;
        int hi = limit_high ? high_s[i] - '0' : 1;  // 二进制填 1

        long long res = 0;
        int d = lo;

        // 通过 limit_low 和 i 可以判断能否不填数字，无需 is_num 参数
        // 如果前导零不影响答案，去掉这个 if block
        if (limit_low && i < diff_lh) {
            // 不填数字，上界不受约束
            res = dfs(i + 1, 0, true, false);
            d = 1; // 下面填数字，从 1 开始填
        }

        for (; d <= hi; d++) {
            // d = 0: cnt += 1; d = 1: cnt += 2; 合并就是 cnt += d + 1
            res += dfs(i + 1, cnt + d + 1, limit_low && d == lo, limit_high && d == hi);
            // res %= MOD;
        }

        if (!limit_low && !limit_high) {
            memo[i][cnt] = res;
        }
        return res;
    };

    return dfs(0, 0, true, true);
}

void solve() {
    ll n; cin >> n;
    int k; cin >> k;
    if (k >= 60) {
        cout << 0 << endl;
        return;
    }
    ll ans = digitDP(1LL, n, k);
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