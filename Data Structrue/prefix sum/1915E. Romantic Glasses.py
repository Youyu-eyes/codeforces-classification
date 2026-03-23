# https://codeforces.com/problemset/problem/1915/E
# 1300
# 前缀和 + 哈希表 + 枚举右维护左
# Python 过不了，但是原理是一样的，见 C++ 代码

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
    s = 0
    cnt = set()
    cnt.add(0)
    for i, x in enumerate(a):
        s += x if i & 1 == 0 else -x
        if s in cnt:
            print('YES')
            break
        cnt.add(s)
    else:
        print('NO')