MOD = 1_000_000_007

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

for i in range(1):
    n, K = MII()
    if K == 0:
        print(pow(2, n, MOD))
        break

    K -= 1
    f_0 = [[0] * n for _ in range(K + 1)]
    s_1 = [[0] * (n + 1) for _ in range(K + 1)]
    f_0[0][0] = 1

    if K >= 1:
        s_1[1][1] = 1

    for i in range(n):
        for k in range(K + 1):
            s_1[k][i + 1] = s_1[k][i]
            for j in range(i - 1, -1, -1):
                c = (i - j) * (i - j + 1) // 2
                if c > k:
                    break
                s_1[k][i + 1] = (s_1[k][i + 1] + f_0[k - c][j]) % MOD
            c_all = (i + 1) * (i + 2) // 2
            if k == c_all:
                s_1[k][i + 1] = (s_1[k][i + 1] + 1) % MOD

            f_0[k][i] = s_1[k][i]

            if k == 0:
                f_0[k][i] = (f_0[k][i] + 1) % MOD

    total_less = 0
    for k in range(K + 1):
        total_less = (total_less + f_0[k][n - 1]) % MOD          # 以 '0' 结尾
        f1_last = (s_1[k][n] - s_1[k][n - 1]) % MOD
        total_less = (total_less + f1_last) % MOD              # 以 '1' 结尾

    ans = (pow(2, n, MOD) - total_less) % MOD
    print(ans)