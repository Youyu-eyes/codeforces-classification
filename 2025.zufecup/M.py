# https://cpc.zufe.ac.cn/contest/7/sub/5/problem/M
# 划分型DP + 单调队列优化
# 详细题解见 cpp 版本

from math import inf

from collections import deque

import sys
input = lambda: sys.stdin.readline().rstrip('\r\n')
print = lambda *args, end='\n', sep=' ': sys.stdout.write(
    sep.join(map(str, args)) + end
)

def II():
    return int(input())

def MII(b = 0):
    return map(lambda x: int(x) - b, input().split())

def LII(b = 0):
    return list(MII(b))

for _ in range(II()):
    n, k = MII()
    q = deque()
    q.append(1)
    f = [inf] * (n + 1)
    f[1] = 0
    for i in range(2, n + 1):
        if i & 1 == 0:
            f[i] = f[i // 2]
        left = max(i - k, (i + 1) // 2)
        while q and q[0] < left:
            q.popleft()

        f[i] = min(f[i], f[q[0]]) + 1

        while q and f[q[-1]] >= f[i]:
            q.pop()
        q.append(i)
    print(f[n])