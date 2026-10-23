# https://codeforces.com/problemset/problem/2150/A
# 1300
# 堆模拟

from heapq import heapify
from heapq import heappop
from copy import deepcopy
si = lambda: input()
ii = lambda: int(input())
many_ii = lambda: map(int, input().split())
many_si = lambda: input().split()
lii = lambda: list(many_ii())
lsi = lambda: list(many_si())
MOD = 1_000_000_007

def solve(s, black):
    h = deepcopy(black)
    heapify(h)
    if h[0] == 1:
        heappop(h)
    pos = 1
    for c in s:
        if c == 'A':
            pos += 1
            if h and pos == h[0]:
                heappop(h)
            else:
                black.append(pos)
        else:
            pos += 1
            while h and pos == h[0]:
                heappop(h)
                pos += 1
            black.append(pos)
            pos += 1
            while h and pos == h[0]:
                heappop(h)
                pos += 1
    
    black.sort()
    return len(black), ' '.join(map(str, black))

T = ii()
for _ in range(T):
    n, m = many_ii()
    s = si()
    black = lii()
    ans = solve(s, black)
    print(ans[0])
    print(ans[1])