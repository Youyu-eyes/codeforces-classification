from functools import cache
from collections import deque
from typing import List, Tuple

def digitDP(low: int, high: int, target: int) -> int:
    low_s = list(map(int, str(low)))
    high_s = list(map(int, str(high)))

    # 二进制情况
    # low_s = list(map(int, bin(low)[2:]))
    # high_s = list(map(int, bin(high)[2:]))

    n = len(high_s)
    diff_lh = n - len(low_s)

    @cache
    def dfs(i: int, cnt0: int, limit_low: bool, limit_high: bool) -> int:
        if cnt0 > target:
            return 0
        if i == n:
            return 1 if cnt0 == target else 0

        lo = low_s[i - diff_lh] if limit_low and i >= diff_lh else 0
        hi = high_s[i] if limit_high else 9  # 二进制填 1

        res = 0
        start = lo

        # 通过 limit_low 和 i 可以判断能否不填数字，无需 is_num 参数
        # 如果前导零不影响答案，去掉这个 if block
        if limit_low and i < diff_lh:
            # 不填数字，上界不受约束
            res = dfs(i + 1, 0, True, False)
            start = 1  # 下面填数字，从 1 开始填

        for d in range(start, hi + 1):
            res += dfs(i + 1,
                       cnt0 + (1 if d == 0 else 0),
                       limit_low and d == lo,
                       limit_high and d == hi)

        # res %= MOD
        return res

    return dfs(0, 0, True, True)

# 满足条件数字的价值总和
# 计算在 [low, high] 中的整数 x 的数位和，满足 x 中的不同数字个数不超过 k
def digitDPContribution(low: int, high: int, k: int) -> int:
    low_s = list(map(int, str(low)))  # 避免在 dfs 中频繁调用 int()
    high_s = list(map(int, str(high)))
    n = len(high_s)
    diff_lh = n - len(low_s)

    # dfs 返回两个数：子树合法数字个数，子树数位总和
    @cache
    def dfs(i: int, mask: int, limit_low: bool, limit_high: bool) -> Tuple[int, int]:
        if i == n:
            # 如果没有特殊约束，那么能递归到终点的都是合法数字
            return 1, 0

        lo = low_s[i - diff_lh] if limit_low and i >= diff_lh else 0
        hi = high_s[i] if limit_high else 9

        cnt = res = 0
        start = lo

        # 如果前导零不影响答案，去掉这个 if block
        if limit_low and i < diff_lh:
            # 不填数字，上界不受约束
            cnt, res = dfs(i + 1, 0, True, False)
            start = 1

        for d in range(start, hi + 1):
            new_mask = mask | 1 << d
            if new_mask.bit_count() > k:  # 不满足要求
                continue
            sub_cnt, sub_sum = dfs(i + 1,
                                   new_mask,
                                   limit_low and d == lo,
                                   limit_high and d == hi)
            cnt += sub_cnt  # 累加子树的合法数字个数
            res += sub_sum  # 累加子树的数位总和
            res += d * sub_cnt  # d 会出现在 sub_cnt 个数中（贡献法）

        # cnt %= MOD; res %= MOD
        return cnt, res

    return dfs(0, 0, True, True)[1]