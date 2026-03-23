# https://codeforces.com/problemset/problem/1999/E
# 1300
# 贪心 + 预处理前缀和

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

def f(n):
    res = 0
    while n:
        n //= 3
        res += 1
    return res

MX = 200_007

s = [0] * (MX + 1)
for i in range(1, MX):
    s[i + 1] = s[i] + f(i)

for _ in range(II()):
    l, r = MII()
    print(s[r + 1] + s[l + 1] - 2 * s[l])