# https://codeforces.com/contest/2165/problem/A
# 1300

import sys
from itertools import pairwise
input = lambda: sys.stdin.readline().strip()
si = lambda: input()
ii = lambda: int(input())
many_ii = lambda: map(int, input().split())
li = lambda: list(many_ii())
MOD = 1_000_000_007

def solve(nums):
    nums += [nums[0]]
    ans = -max(nums)
    for x, y in pairwise(nums):
        ans += max(x, y)
    return ans
T = ii()
for _ in range(T):
    n = ii()
    nums = li()
    ans = solve(nums)
    print(ans)