# https://codeforces.com/contest/1948/problem/C
# 1300
# 图的遍历，递推 DFS
# 同样，递归会被卡
# 广搜版本见 BFS

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
    
    stack = [(0, 0)]    
    while stack:
        i, j = stack.pop()
        
        if vis[i][j]:
            continue
        vis[i][j] = True
        
        for dx, dy in [(-1, 0), (1, 0), (0, -1), (0, 1)]:
            x, y = i + dx, j + dy
            if 0 <= x < 2 and 0 <= y < n and not vis[x][y]:
                if grid[x][y] == '<':
                    next_y = y - 1
                else:
                    next_y = y + 1
                stack.append((x, next_y))

    print('YES' if vis[-1][-1] else 'NO')