# https://codeforces.com/problemset/problem/2064/C
# 1300
# 贪心 + 前后缀分解
# 首先发现每次只取最左边正数和最右边负数，如果取最左边第二个正数，第一个正数会被浪费
# 其次是发现取到最后形成的分割线，左边全是正数，右边全是负数
# 求左边正数前缀和 + 右边负数后缀和绝对值，左边正数前缀和可以合并到最终循环（前后缀分解技巧）

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
    suf = [0] * n
    suf[-1] = -min(a[-1], 0)
    for i in range(n - 2, -1, -1):
        suf[i] = suf[i + 1] - min(a[i], 0)
    
    ans = suf[0]
    pre = 0
    for i, x in enumerate(a):
        pre += max(x, 0)
        ans = max(ans, pre + suf[i])
    print(ans)