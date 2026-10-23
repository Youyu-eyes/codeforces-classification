# https://codeforces.com/problemset/problem/2086/C
# 1300
# 模拟，链表，图论
# 数据结构题单没有放链表，所以归在图论
# 置换环问题，是一个链式置换
# 比较难想的是，一次遍历会把这条链上所有的数全部置换一边
# 如果遍历到一个原来访问过的数，那么这个数在之前的操作中就被置换了
# 新数一定对应一条新链，相当于连通块

import sys
input = lambda: sys.stdin.readline().rstrip('\r\n')
print = lambda *args, end='\n', sep=' ': sys.stdout.write(
    sep.join(map(str, args)) + end
)

def II():
    return int(input())

def MII(b = 1):
    return map(lambda x: int(x) - b, input().split())

def LII():
    return list(MII())

for _ in range(II()):
    n = II()
    p = LII()
    d = LII()
    st = set()
    ans = [0] * n
    for i, x in enumerate(d):
        while x not in st:
            st.add(x)
            x = p[x]
        ans[i] = str(len(st))
    print(' '.join(ans))