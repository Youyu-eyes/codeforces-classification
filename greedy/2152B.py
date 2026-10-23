# https://codeforces.com/contest/2152/problem/B
# 1300

import sys
input = lambda: sys.stdin.readline().strip()
si = lambda: input()
ii = lambda: int(input())
many_ii = lambda: map(int, input().split())
li = lambda: list(many_ii())
MOD = 1_000_000_007

def judge(n, D, K):
    if D < K:
        return n - D
    elif K < D:
        return D
    else:
        return 0

def solve(n, rK, cK, rD, cD):
    return max(judge(n, rD, rK), judge(n, cD, cK))

T = ii()
for _ in range(T):
    n, rK, cK, rD, cD = many_ii()
    print(solve(n, rK, cK, rD, cD))