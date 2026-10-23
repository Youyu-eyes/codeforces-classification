# https://codeforces.com/problemset/problem/2144/C
# 1300

import sys
from functools import cache
input = lambda: sys.stdin.readline().strip()
si = lambda: input()
ii = lambda: int(input())
many_ii = lambda: map(int, input().split())
many_si = lambda: input().split()
lii = lambda: list(many_ii())
lsi = lambda: list(many_si())
MOD = 998244353

# 记忆化搜索
def solve1(a, b):
    @cache
    def dfs(i, j):
        if i == 0:
            return 1
        res = 0
        if a[i] >= b[i - 1] and b[i] >= a[i - 1]:
            res += dfs(i - 1, j ^ 1)
        if a[i] >= a[i - 1] and b[i] >= b[i - 1]:
            res += dfs(i - 1, j)
        return res % MOD
    return (dfs(len(a) - 1, 0) + dfs(len(a) - 1, 1)) % MOD

# 递推
def solve2(a, b):
    n = len(a)
    f = [[0] * 2 for _ in range(n)]
    f[0][0] = f[0][1] = 1
    for i in range(1, n):
        for j in range(2):
            if a[i] >= b[i - 1] and b[i] >= a[i - 1]:
                f[i][j] += f[i - 1][j ^ 1]
            if a[i] >= a[i - 1] and b[i] >= b[i - 1]:
                f[i][j] += f[i - 1][j]
            f[i][j] %= MOD
    return sum(f[-1]) % MOD

# 空间优化
def solve3(a, b):
    n = len(a)
    f0 = f1 = 1
    for i in range(1, n):
        n_f0 = n_f1 = 0
        if a[i] >= b[i - 1] and b[i] >= a[i - 1]:
            n_f0 += f1
            n_f1 += f0
        if a[i] >= a[i - 1] and b[i] >= b[i - 1]:
            n_f0 += f0
            n_f1 += f1
        f0 = n_f0 % MOD
        f1 = n_f1 % MOD
    return (f0 + f1) % MOD

T = ii()
for _ in range(T):
    n = ii()
    a = lii()
    b = lii()
    print(solve3(a, b))