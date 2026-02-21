from functools import cache

def digitDP(low, high):
    low = list(map(int, str(low)))
    high = list(map(int, str(high)))
    low = [0] * (len(high) - len(low)) + low

    @cache
    def dfs(i, limit_low, limit_high, is_num):
        if i == len(high):
            return 1
        
        res = 0
        if not is_num and low[i] == 0:
            res = dfs(i + 1, True, False, False)

        # 不动产！！！
        lo = low[i] if limit_low else 0
        hi = high[i] if limit_high else 9

        for d in range(max(1, lo), hi + 1):
            res += dfs(i + 1, limit_low and d == lo, limit_high and d == hi, True)
        
        return res
    return dfs(0, True, True, False)