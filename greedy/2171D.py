# https://codeforces.com/contest/2171/problem/D
# 1400

from math import inf
si = lambda: input()
ii = lambda: int(input())
many_ii = lambda: map(int, input().split())
many_si = lambda: input().split()
lii = lambda: list(many_ii())
lsi = lambda: list(many_si())
MOD = 1_000_000_007

def solve(n, nums):
    suf = [0] * n
    suf[-1] = nums[-1]
    for i in range(n - 2, -1, -1):
        suf[i] = max(suf[i + 1], nums[i])
    
    pre = inf
    for i in range(n - 1):
        pre = min(pre, nums[i])
        if pre > suf[i + 1]:
            return False
    return True

T = ii()
for _ in range(T):
    n = ii()
    nums = lii()
    if solve(n, nums):
        print('Yes')
    else:
        print('No')