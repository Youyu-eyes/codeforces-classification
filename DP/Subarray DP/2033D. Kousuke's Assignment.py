# https://codeforces.com/problemset/problem/2033/D
# DP 解法
# 1300

import sys
input = lambda: sys.stdin.readline().split()
si = lambda: input()[0]
ii = lambda: int(si())
lii = lambda: list(map(int, input()))
lsi = lambda: input()

# 动态规划解法
# 动态规划会超时，不知为何，可能是哈希表常数太大
# PyPy 也过不了，PyPy 的哈希表常数比 Python 还大
# 能过的方法见贪心题单

# 优化前
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
    nums = LII()
    s = 0
    cnt = {0: -1}
    f = [0] * (n + 1)
    for i, x in enumerate(nums):
        s += x
        f[i + 1] = f[i]
        if s in cnt:
            f[i + 1] = max(f[i + 1], f[cnt[s] + 1] + 1)
        cnt[s] = i
    print(f[-1])

# 滚动数组优化，哈希表直接统计 max(f[cnt[s] + 1])
# 这实际上已经是一个贪心了，见贪心题单
def II():
    return int(input())

def MII(b = 0):
    return map(lambda x: int(x) - b, input().split())

def LII():
    return list(MII())

for _ in range(II()):
    n = II()
    nums = LII()
    s = 0
    cnt = {0: 0}
    f = 0
    for x in nums:
        s += x
        if s in cnt:
            f = max(f, cnt[s] + 1)
        cnt[s] = f
    print(f)