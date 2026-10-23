# https://codeforces.com/contest/2145/problem/C
# 1300

import sys
from math import inf
from collections import defaultdict
input = lambda: sys.stdin.readline().strip()
si = lambda: input()
ii = lambda: int(input())
many_ii = lambda: map(int, input().split())
li = lambda: list(many_ii())
MOD = 1_000_000_007

def solve(s):
    d = s.count('a') - s.count('b')
    if d == 0:
        return 0
    cnt = defaultdict(int)
    cnt[d - 1] = -1
    pre = 0
    ans = inf
    for i, c in enumerate(s):
        if c == 'b':
            pre += 1
        x = i - 2 * pre
        if x in cnt:
            ans = min(ans, i - cnt[x])
        cnt[d + x] = i
    return ans if ans < n else -1


T = ii()
for _ in range(T):
    n = ii()
    s = si()
    print(solve(s))