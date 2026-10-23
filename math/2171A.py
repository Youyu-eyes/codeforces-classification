# https://codeforces.com/contest/2171/problem/A
# 800

import sys
input = lambda: sys.stdin.readline().strip()
si = lambda: input()
ii = lambda: int(input())
many_ii = lambda: map(int, input().split())
li = lambda: list(many_ii)
MOD = 1_000_000_007

def solve(n):
    if n % 2:
        return 0
    return n // 4 + 1

T = ii()
for _ in range(T):
    n = ii()
    print(solve(n))