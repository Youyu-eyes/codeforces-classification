# https://codeforces.com/problemset/problem/2116/B
# 1300
# 贪心，位运算思想
# 比较 (2 ** a + 2 ** b) 与 (2 ** c + 2 ** d)
# 相当于比较 (1 << a + 1 << b) 与 (1 << c + 1 << d)
# 类比十进制比大小，从最高位开始比，二进制也是，哪个数的最高位高，哪个数大

# 如果预处理二进制，时间复杂度 O(n)
# 不预处理二进制数，pow 函数自带快速幂计算，复杂度 O(nlogn)，实测 Python 也能过

MOD = 998_244_353

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

MX = 100_007
pow2 = [1] * MX
for i in range(1, MX):
    pow2[i] = pow2[i - 1] * 2 % MOD

for _ in range(II()):
    n = II()
    p = LII()
    q = LII()
    j = k = 0
    ans = [0] * n
    for i in range(n):
        if p[i] > p[j]:
            j = i
        if q[i] > q[k]:
            k = i
        if p[j] > q[k]:
            ans[i] = str((pow2[p[j]] + pow2[q[i - j]]) % MOD)
        elif p[j] < q[k]:
            ans[i] = str((pow2[q[k]] + pow2[p[i - k]]) % MOD)
        else:
            ans[i] = str((pow2[p[j]] + pow2[max(q[i - j], p[i - k])]) % MOD)
    print(' '.join(ans))