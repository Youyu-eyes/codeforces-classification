# https://codeforces.com/contest/189/problem/A
# 1300 
# 远古题，个人感觉高了 大概 1000
# 编程入门的三种小鸡问题

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

n, a, b, c = MII()
a, b, c = sorted([a, b, c])
ans = 0
flag = False
for x in range(n // a, -1, -1):
    for y in range((n - a * x) // b, -1, -1):
        if (n - a * x - b * y) % c == 0:
            ans = max(ans, (x + y + (n - a * x - b * y) // c))
print(ans)