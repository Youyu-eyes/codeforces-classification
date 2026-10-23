# https://codeforces.com/contest/1948/problem/C
# 1300
# 广搜 BFS
# DFS 也可，但是卡递归，详见 DFS 部分

import sys
input = lambda: sys.stdin.readline().rstrip('\r\n')
print = lambda *args, end='\n', sep=' ': sys.stdout.write(
    sep.join(map(str, args)) + end
)

def II():
    return int(input())

def MII(b = 0):
    return map(lambda x: int(x) - b, input().split())

def LII():
    return list(MII())

for _ in range(II()):
    n = II()
    grid = []
    for _ in range(2):
        grid.append(input())
    vis = [[False] * n for _ in range(2)]

    q = [(0, 0)]
    while q:
        cur = q
        q = []
        for i, j in cur:
            for x, y in (i, j - 1), (i, j + 1), (i - 1, j), (i + 1, j):
                if 0 <= x < 2 and 0 <= y < n and not vis[x][y]:
                    vis[x][y] = True
                    if grid[x][y] == '<':
                        y -= 1
                    else:
                        y += 1
                    vis[x][y] = True
                    q.append((x, y))
    print('YES' if vis[-1][-1] else 'NO')