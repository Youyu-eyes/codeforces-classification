# https://codeforces.com/contest/2119/problem/C
# 1300
# 需要比较熟悉位运算

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
    n, l, r, k = MII()
    if n & 1:
        print(l)
        continue
    if n == 2:
        print(-1)
        continue
    ans = 1
    while ans <= r:
        if ans > l:
            print(ans if k > n - 2 else l)
            break
        ans <<= 1
    else:
        print(-1)