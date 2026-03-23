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
                       cnt0 + (d == 0), // 根据题目更改
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


// 满足条件数字的价值总和
// 计算在 [low, high] 中的整数 x 的数位和，满足 x 中的不同数字个数不超过 k
long long digitDPContribution(long long low, long long high, int k) {
    string low_s = to_string(low);
    string high_s = to_string(high);
    int n = high_s.size();
    int diff_lh = n - low_s.size();
    vector memo(n, vector<pair<long long, long long>>(1 << 10, {-1, -1}));

    // dfs 返回两个数：子树合法数字个数，子树数位总和
    auto dfs = [&](this auto&& dfs, int i, int mask, bool limit_low, bool limit_high) -> pair<long long, long long> {
        if (i == n) {
            // 如果没有特殊约束，那么能递归到终点的都是合法数字
            return {1, 0};
        }

        if (!limit_low && !limit_high && memo[i][mask].first >= 0) {
            return memo[i][mask];
        }

        int lo = limit_low && i >= diff_lh ? low_s[i - diff_lh] - '0' : 0;
        int hi = limit_high ? high_s[i] - '0' : 9;

        long long cnt = 0, sum = 0;
        int d = lo;

        // 如果前导零不影响答案，去掉这个 if block
        if (limit_low && i < diff_lh) {
            // 不填数字，上界不受约束
            tie(cnt, sum) = dfs(i + 1, 0, true, false);
            d = 1;
        }

        for (; d <= hi; d++) {
            int new_mask = mask | 1 << d;
            if (popcount((uint32_t) new_mask) > k) { // 不满足要求
                continue;
            }
            auto [sub_cnt, sub_sum] = dfs(i + 1, new_mask, limit_low && d == lo, limit_high && d == hi);
            cnt += sub_cnt; // 累加子树的合法数字个数
            sum += sub_sum; // 累加子树的数位总和
            sum += d * sub_cnt; // d 会出现在 sub_cnt 个数中（贡献法）
            // cnt %= MOD; sum %= MOD;
        }

        pair<long long, long long> res = {cnt, sum};
        if (!limit_low && !limit_high) {
            memo[i][mask] = res;
        }
        return res;
    };

    return dfs(0, 0, true, true).second;
}