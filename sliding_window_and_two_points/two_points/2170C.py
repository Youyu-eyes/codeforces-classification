# https://codeforces.com/problemset/problem/2170/C
# 好题
# 1300
# 贪心 + 式子变形：
# x = q(r + 1) + r <= k
# ==> qr + q + r <= k
# ==> (q + 1)(r + 1) <= (k + 1) 
# ==> (r + 1) <= (k + 1) // (q + 1)

# 排序 + 双指针 + 二分: 
# 在 r 数组中二分，如果 r[p2] 不满足题意 q[p1]，q[p1] 变大，r[p2] 更不满足条件，可以直接缩小二分范围

from bisect import bisect_left

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
    n, k = MII()
    lq = sorted(LII())
    lr = sorted(LII(b = -1))
    p1 = 0

    # p2 开区间，因为 bisect_left 右端点参数默认开区间
    p2 = n
    ans = 0
    while p1 < n and p2 >= 0:
        q = lq[p1]
        p2 = bisect_left(lr, (k + 1) // (q + 1) + 1, 0, p2) - 1
        if p2 >= 0:
            ans += 1
        p1 += 1
    print(ans)