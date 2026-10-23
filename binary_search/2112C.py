# https://codeforces.com/contest/2112/problem/C
# 1300
# 复杂度 O(n^2logn)

from bisect import bisect_left

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
    return list(MII())

for _ in range(II()):
    n = II()
    nums = LII()
    mx = nums[-1]
    ans = 0
    for i in range(2, n):
        for j in range(1, i):
            k = bisect_left(nums, max(mx, 2 * nums[i]) - nums[i] - nums[j] + 1, 0, j)
            ans += j - k
    print(ans)