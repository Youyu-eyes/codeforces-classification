# https://codeforces.com/contest/2172/problem/M
# 1300

import sys
from collections import deque

input = lambda: sys.stdin.readline().strip()
si = lambda: input()
ii = lambda: int(input())
many_ii = lambda: map(int, input().split())
li = lambda: list(many_ii())
MOD = 1_000_000_007

def solve(n, g, types, k):
    dist = [-1] * (n + 1)
    dist[1] = 0
    q = deque([1])
    
    while q:
        u = q.popleft()
        for v in g[u]:
            if dist[v] == -1:
                dist[v] = dist[u] + 1
                q.append(v)

    max_dist = [-1] * (k + 1)
    for i in range(1, n + 1):
        t = types[i - 1]
        max_dist[t] = max(max_dist[t], dist[i])
    
    return max_dist[1:]

n, m, k = many_ii()
a = li()
g = [[] for _ in range(n + 1)]
for _ in range(m):
    x, y = many_ii()
    g[x].append(y)
    g[y].append(x)
ans = solve(n, g, a, k)
print(' '.join(map(str, ans)))