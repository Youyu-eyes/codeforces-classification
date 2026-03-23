# https://codeforces.com/contest/2065/problem/C2
# 1300
# 从最小/最大贪心 + 二分
# 实现是最小开始，也可以从 n-1 回推

from bisect import bisect_left

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
    n, m = MII()
    a = LII()
    b = sorted(LII())
    a[0] = min(a[0], b[0] - a[0])
    for i in range(1, n):
        j = bisect_left(b, a[i - 1] + a[i])
        if a[i - 1] > a[i]:
            if j == m:
                print('NO')
                break
            a[i] = b[j] - a[i]
        elif j < m:
            a[i] = min(a[i], b[j] - a[i])

    else:
        print('YES')