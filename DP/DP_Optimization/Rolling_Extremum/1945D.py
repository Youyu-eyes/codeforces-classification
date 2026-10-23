# https://codeforces.com/contest/1945/problem/D
# 前缀最小值优化 DP
# 估计 1400，有 1300 贪心做法，详见贪心题单

# 由题意易得状态转移方程：f[i] = min(f[j] + a[j] + sum(b[k] for k in range(j + 1, i)) for j in range(i))
# 定义 b 的前缀和数组 s，s[0] = 0，s[1] = b[0]
# 根据前缀和定义，sum(b[k] for k in range(j + 1, i)) = s[i] - s[j + 1]
# f[i] = min(f[j] + a[j] + s[i] - s[j + 1] for j in range(i))
# s[i] 与 j 取值无关
# f[i] = s[i] + min(f[j] + a[j] - s[j + 1] for j in range(i))
# 记 g[j] = min(f[j] + a[j] - s[j + 1] for j in range(i))
# 0 <= j < m 时，f[j] = 0
# f[i] = s[i] + min(a[j] - s[j + 1] for j in range(i))
# 记 min1 = min(a[j] - s[j + 1] for j in range(i))
# f[i] = s[i] + min1
# m <= j < i 时
# 记 min2[i] = min(f[j] + a[j] - s[j + 1] for j in range(m, i))
# 则 g[j] = min(g[j1] for j1 in range(m), g[j2] for j2 in range(m, i)) = min(min1, min2[i])

# 关键递推式
# min2[i + 1] = min(min2[i], f[i] + a[i] - s[i + 1])
# 则 min2 可用一个变量动态维护

from math import inf
from itertools import accumulate

import sys
input = lambda: sys.stdin.readline().rstrip('\r\n')
print = lambda *args, end='\n', sep=' ': sys.stdout.write(
    sep.join(map(str, args)) + end
)

def II():
    return int(input())

def MII(b = 0):
    return map(lambda x: int(x) - b, input().split())

def LII(b = 0):
    return list(MII(b))

for _ in range(II()):
    n, m = MII()
    a = LII()
    b = LII()
    s = list(accumulate(b, initial = 0))

    min1 = min(a[j] - s[j + 1] for j in range(m))
    min2 = inf
    f = 0
    for i in range(m, n + 1):
        f = s[i] + min(min1, min2)
        if i < n:
            min2 = min(min2, f + a[i] - s[i + 1])
    print(f)