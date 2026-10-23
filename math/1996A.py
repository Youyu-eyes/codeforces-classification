# https://codeforces.com/contest/1996/problem/A
# 800

import sys
input = lambda: sys.stdin.readline().strip()
si = lambda: input()
ii = lambda: int(input())
many_ii = lambda: map(int, input().split())
li = lambda: list(many_ii)
MOD = 1_000_000_007

def solve(n):
    ans, n = divmod(n, 4)
    ans += n // 2
    return ans

T = ii()
for _ in range(T):
    n = ii()
    print(solve(n))