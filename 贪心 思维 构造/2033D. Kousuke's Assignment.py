# https://codeforces.com/problemset/problem/2033/D

from collections import defaultdict

si = lambda: input()
ii = lambda: int(input())
many_ii = lambda: map(int, input().split())
many_si = lambda: input().split()
lii = lambda: list(many_ii())
lsi = lambda: list(many_si())
MOD = 1_000_000_007

# 贪心解法
# 只有贪心能过，不知道为何
# DP 解法另见 动态规划-合法子数组DP

def solve_greed():
    n = ii()
    nums = lii()
    ans = s = 0
    has = set()
    has.add(0)
    for x in nums:
        s += x
        if s in has:
            ans += 1
            has.clear()
        has.add(s)

    print(ans)

T = ii()
for _ in range(T):
    solve_greed()