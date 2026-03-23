# https://codeforces.com/contest/1945/problem/D
# 1300
# 贪心做法

# 将数组画成以 m 为分割线的两部分，[0, m) 和 [m, n)
# 对于右边，对于任意一个 j，如果 i 与 j 交换，则代价是 a[j]，如果不交换，代价是 b[j]
# 没有规定最多交换距离的情况下，对于任意的 j，均可以自由选择交换 或 不交换
# 贪心的想，如果 a[j] > b[j]，不交；a[j] < b[j]，交换
# 即对于 j in [m, n)，代价为 min(a[j], b[j])，求和 sum(min(a[j], b[j]))，记为 min2
# 对于 j in [0, m)
# 我们只做一次交换，代价最小。因为如果做第二次交换，一定比做完一次交换的代价大
# 交换的代价是 min(a[j] + sum(b[k] for k in range(j + 1, m)))
# 前缀和优化后：min(a[j] + s[m] - s[j + 1])，记为 min1
# 答案为 min1 + min2

# 另有 1400 连续最小值优化 DP，见 DP 题单 -> 优化 DP 章节

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

    min1 = inf
    for i in range(m):
        min1 = min(min1, a[i] + s[m] - s[i + 1])
    ans = sum(min(a[j], b[j]) for j in range(m, n))
    print(ans + min1)