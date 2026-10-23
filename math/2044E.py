# https://codeforces.com/problemset/problem/2044/E
# 1300
# 纯数学

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
    k, l1, r1, l2, r2 = MII()
    kn = 1; ans = 0
    while r2 // kn >= l1:
        ans += max(0, min(r2 // kn, r1) - max((l2 - 1) // kn + 1, l1) + 1)
        kn *= k
    print(ans)