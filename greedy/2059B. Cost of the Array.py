# https://codeforces.com/problemset/problem/2059/B
# 1300
# 贪心 + 构造

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
    n, k = MII()
    nums = LII()
    if n == k:
        for i in range(1, n, 2):
            if nums[i] != i // 2 + 1:
                print(i // 2 + 1)
                break
        else:
            print(k // 2 + 1)
        continue
    for i in range(1, n - k + 2):
        if nums[i] != 1:
            print(1)
            break
    else:
        print(2)