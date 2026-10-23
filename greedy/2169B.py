# https://codeforces.com/contest/2169/problem/B
# 1100 正常写法
# 1250 正则式匹配

import sys
import re
input = lambda: sys.stdin.readline().strip()
si = lambda: input()
ii = lambda: int(input())
many_ii = lambda: map(int, input().split())
li = lambda: list(many_ii)
MOD = 1_000_000_007

def solve(s):
    if s == '*':
        return 1
    match = re.match(r'^(<*)(\*?)(>*)$', s)
    if match:
        a = len(match.group(1))
        b = len(match.group(2))
        c = len(match.group(3))
        if a > 0 or c > 0:
            return max(a, c) + b
    return -1

T = ii()
for _ in range(T):
    s = si()
    print(solve(s))