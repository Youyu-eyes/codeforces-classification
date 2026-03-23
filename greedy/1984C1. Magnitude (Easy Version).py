# https://codeforces.com/problemset/problem/1984/C1
# 1300
# 贪心，只进行一次操作二

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
    a = LII()
    s = mn = 0
    for x in a:
        s += x
        mn = min(mn, s)
    print(s - 2 * mn)