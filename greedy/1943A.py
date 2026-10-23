# https://codeforces.com/contest/1943/problem/A
# 1300
# 贪心 + 博弈

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
    a = LII()

    cnt = [0] * (n + 1)
    for x in a:
        cnt[x] += 1
    
    c = 0
    for i in range(n + 1):
        c += int(cnt[i] == 1)
        if c == 2 or cnt[i] == 0:
            print(i)
            break