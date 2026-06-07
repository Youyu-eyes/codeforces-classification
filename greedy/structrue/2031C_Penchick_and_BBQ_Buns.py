# https://codeforces.com/contest/2031/problem/C
# 1300
# 构造，完全平方数

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
    n = II()
    if n & 1 == 0:
        ans = []
        for i in range(1, n // 2 + 1):
            ans.extend([i, i])
        print(' '.join(map(str, ans)))
        continue
    
    if n <= 25:
        print(-1)
        continue

    ans = [1, 3, 3, 4, 4, 5, 5, 6, 6, 1, 2, 7, 7, 8, 8, 9, 9, 10, 10, 11, 11, 12, 12, 13, 13, 1, 2]
    for i in range(1, (n - 27) // 2 + 1):
        ans.extend([i + 13, i + 13])
    print(' '.join(map(str, ans)))