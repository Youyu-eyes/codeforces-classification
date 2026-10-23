# https://codeforces.com/contest/1919/problem/C
# 1400

from math import inf
si = lambda: input()
ii = lambda: int(input())
many_ii = lambda: map(int, input().split())
many_si = lambda: input().split()
lii = lambda: list(many_ii())
lsi = lambda: list(many_si())
MOD = 1_000_000_007

# 另一解法见动态规划：线段树优化DP
# 贪心做法：只会出现 3 种情况，令 x <= y
# 1. a <= x，放在 x 后面最优，不会增加
# 2. a > y，放在 x 后面最优，会增加
# 3. x < a <= y，放在 y 后面最优，不会增加 （要证明只需假设后面一个数 a1，枚举 a1 和 a，y 的大小关系和 a，x 的大小关系）

def solve(nums):
    x = y = inf
    ans = 0
    for a in nums:
        if a <= x:
            x = a
        elif a > y:
            ans += 1
            x = y
            y = a
        else:
            y = a
    return ans

T = ii()
for _ in range(T):
    n = ii()
    nums = lii()
    print(solve(nums))