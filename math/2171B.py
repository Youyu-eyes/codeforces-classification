# https://codeforces.com/contest/2171/problem/B
# 900

import sys
input = lambda: sys.stdin.readline().strip()
MOD = 1_000_000_007

def solve(nums):
    diff = abs(nums[0] - nums[-1])
    if nums[0] == -1 and nums[-1] >= 0:
        nums[0] = nums[-1]
        diff = 0
    if nums[-1] == -1 and nums[0] >= 0:
        nums[-1] = nums[0]
        diff = 0
    for i, x in enumerate(nums):
        if x == -1:
            nums[i] = 0
    return nums, diff

T = int(input())
for _ in range(T):
    n = int(input())
    nums, diff = solve(list(map(int, input().split())))
    print(diff)
    for x in nums:
        print(x, end = ' ')
    print()