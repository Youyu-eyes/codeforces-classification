# https://codeforces.com/problemset/problem/2064/C
# 1300
# 提前声明，这一版过不了
# 区间DP，用 DP 的话难度分大概 1200 左右，会低一点
# 别试了，C++ 也会超内存（狗头）
# 写递推的时候要注意一点，l 要倒序枚举
# 能过的方法见贪心题单

from functools import cache

import sys
input = lambda: sys.stdin.readline().rstrip('\r\n')
print = lambda *args, end='\n', sep=' ': sys.stdout.write(
    sep.join(map(str, args)) + end
)

def II():
    return int(input())

def MII(b = 0):
    return map(lambda x: int(x) - b, input().split())

def LII():
    return list(MII())

for _ in range(II()):
    n = II()
    a = LII()

    # 记忆化搜索
    @cache
    def dfs(l, r):
        if l == r:
            return 0
        res = 0
        if a[i] > 0:
            res = max(res, dfs(i + 1, r) + a[i])
        else:
            res = max(res, dfs(l, i) - a[i])
        return res
    print(dfs(0, n))

    # 递推
    f = [[0] * (n + 1) for _ in range(n + 1)]
    for l in range(n - 1, -1, -1):
        for r in range(l + 1, n + 1):
            res = 0
            for i in range(l, r):
                if a[i] > 0:
                    res = max(res, f[i + 1][r] + a[i])
                else:
                    res = max(res, f[l][i] - a[i])
            f[l][r] = res
    print(f[0][n])