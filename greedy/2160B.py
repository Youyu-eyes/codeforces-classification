# https://codeforces.com/contest/2160/problem/B
# 1100

import sys
input = lambda: sys.stdin.readline().strip()
si = lambda: input()
ii = lambda: int(input())
many_ii = lambda: map(int, input().split())
li = lambda: list(many_ii())
MOD = 1_000_000_007

def solve(n, b):
    a = [0] * n
    a[0] = 1
    max_n = 2
    for i in range(1, n):
        j = i + b[i - 1] - b[i]
        if j < 0:
            a[i] = max_n
            max_n += 1
        else:
            a[i] = a[j]
    return a
        
T = ii()
for _ in range(T):
    n = ii()
    b = li()
    a = solve(n, b)
    for x in a:
        print(x, end = ' ')
    print()