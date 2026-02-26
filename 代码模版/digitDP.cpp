#include <bits/stdc++.h>
using namespace std;


long long digitDP(long long low, long long high, int target) {
    string low_s = to_string(low);
    string high_s = to_string(high);

    // 二进制
    // string low_s; for (int x = low; x; x >>= 1) low_s.push_back('0' + (x & 1)); reverse(low_s.begin(), low_s.end()); if (low == 0) low_s = "0";
    // string high_s; for (int x = high; x; x >>= 1) high_s.push_back('0' + (x & 1)); reverse(high_s.begin(), high_s.end()); if (high == 0) high_s = "0";

    int n = high_s.size();
    int diff_lh = n - low_s.size();
    vector memo(n, vector<long long>(target + 1, -1));

    auto dfs = [&](this auto&& dfs, int i, int cnt0, bool limit_low, bool limit_high) -> long long {
        if (cnt0 > target) {
            return 0; // 不合法
        }
        if (i == n) {
            return cnt0 == target;
        }

        if (!limit_low && !limit_high && memo[i][cnt0] >= 0) {
            return memo[i][cnt0];
        }

        int lo = limit_low && i >= diff_lh ? low_s[i - diff_lh] - '0' : 0;
        int hi = limit_high ? high_s[i] - '0' : 9;  // 二进制填 1

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
            res += dfs(i + 1, 
                       cnt0 + (d == 0), 
                       limit_low && d == lo, 
                       limit_high && d == hi);
            // res %= MOD;
        }

        if (!limit_low && !limit_high) {
            memo[i][cnt0] = res;
        }
        return res;
    };

    return dfs(0, 0, true, true);
}