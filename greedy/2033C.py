# https://codeforces.com/problemset/problem/2033/C
# 1400

from itertools import pairwise
si = lambda: input()
ii = lambda: int(input())
many_ii = lambda: map(int, input().split())
many_si = lambda: input().split()
lii = lambda: list(many_ii())
lsi = lambda: list(many_si())
MOD = 1_000_000_007

def solve(n, a):
    for i in range(1, n // 2):
        if a[i - 1] == a[i] or a[n - 1 - i] == a[n - i]:
            a[i], a[n - i - 1] = a[n - i - 1], a[i]
    
    return sum(x == y for x, y in pairwise(a))

T = ii()
for _ in range(T):
    n = ii()
    a = lii()
    print(solve(n, a))