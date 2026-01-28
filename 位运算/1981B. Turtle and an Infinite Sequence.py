# https://codeforces.com/problemset/problem/1981/B
# 1300
# 结论题，见 结论文件夹 -> 区间按位或

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
    n, m = MII()
    l = max(n - m, 0)
    r = m + n
    if l == r:
        print(n)
        continue
    p = 0
    diff = l ^ r
    while diff > 0:
        diff >>= 1
        p += 1
    print(r | ((1 << p) - 1))