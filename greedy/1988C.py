# https://codeforces.com/contest/1988/problem/C
# 1300
# 位运算构造
# 位运算 + 贪心 + 计数
# 拆位法
# 如果只有一个置位，则只能填这个数本身
# 否则，从第二位开始贪心
# 第一次把正数第二位置 1，其余位置置 0
# 第二个数为第一个数对于 n 的反转，即 0 -> 1, 1 -> 0
# 模拟一遍，发现中间的数是 n 去掉除最高置位，每次去掉一个 1 产生的数
# 一共有 C(s - 2, s - 1) = s - 1，然后包括第一个数和最后一个数，共 s + 1 个，其中 s = n.bit_count()
# 可以证明没有更长的序列满足题意
# 最后一个数是 n

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
    st = []
    i = 0
    while i < n.bit_length():
        if n >> i & 1:
            st.append(i)
        i += 1
    size = len(st)
    if size == 1:
        print(1)
        print(n)
        continue
    print(size + 1)
    ans = [0] * (size + 1)
    ans[0] = str(1 << st[-2])
    ans[-1] = str(n)
    j = 1
    for i in range(len(st) - 2, -1, -1):
        ans[j] = str(n - (1 << st[i]))
        j += 1
    print(' '.join(ans))