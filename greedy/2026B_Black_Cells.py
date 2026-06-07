# https://codeforces.com/problemset/problem/2026/B
# 1300
# 贪心 + 枚举
# 主要想明白为什么要取 (a1, a2), (a3, a4) …… 这样的元组
# 假设涂 (a2, a3)，那么势必要涂 (a1, a4)， 在 a 有序的前提下，a4 - a1 > (a4 - a3 or a2 - a1)
# 取最大值一定会取 a1-a4，a2 和 a3 的距离就不重要了

from math import inf

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

def cal(a):
    s_a = sorted(a)
    global ans
    ans = min(ans, max(s_a[i + 1] - s_a[i] for i in range(0, n, 2)))   

for _ in range(II()):
    n = II()
    a = LII()
    ans = inf
    st = set(a)
    if n % 2 == 0:
        cal(a)
        print(ans)
        continue
    for x in a:
        if x + 1 not in st:
            a.append(x + 1)
            cal(a)
            a.pop()
    print(ans)