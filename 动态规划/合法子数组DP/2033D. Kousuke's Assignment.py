# https://codeforces.com/problemset/problem/2033/D

from collections import defaultdict
from math import inf

si = lambda: input()
ii = lambda: int(input())
many_ii = lambda: map(int, input().split())
many_si = lambda: input().split()
lii = lambda: list(many_ii())
lsi = lambda: list(many_si())
MOD = 1_000_000_007

# 动态规划解法
# 动态规划会超时，不知为何，可能是哈希表常数太大
# 能过的方法见贪心题单

def solve_DP():
    n = int(input())
    nums = list(map(int, input().split()))
    s = 0
    cnt = defaultdict(lambda: inf)
    cnt[0] = -1
    f = [0] * (n + 1)
    for i, x in enumerate(nums):
        s += x
        f[i + 1] = f[i]
        if cnt[s] < inf:
            f[i + 1] = max(f[i + 1], f[cnt[s] + 1] + 1)
        cnt[s] = i

    print(f[-1])

T = ii()
for _ in range(T):
    solve_DP()