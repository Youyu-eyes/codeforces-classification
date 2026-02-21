from functools import cache

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