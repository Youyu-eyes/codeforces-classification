# https://codeforces.com/contest/2091/problem/E
# 1300
# 埃氏筛
# 10^7 数据 Python 埃氏筛跑不动，python 最大跑到 10^6
# 必须用 C++，具体代码看题解

import sys
input = lambda: sys.stdin.readline().split()
si = lambda: input()[0]
ii = lambda: int(si())
lii = lambda: list(map(int, input()))
lsi = lambda: input()

MX = 10_000_001
is_prime = [False] * 2 + [True] * (MX - 2)
primes = []
for i in range(2, MX):
    if is_prime[i]:
        primes.append(i)
        for j in range(i * i, MX, i):
            is_prime[j] = False

def solve():
    T = ii()
    out = []
    for _ in range(T):
        n = ii()
        ans = 0
        i = 0
        while primes[i] <= n:
            ans += n // primes[i]
            i += 1
        out.append(str(ans))
    sys.stdout.buffer.write('\n'.join(out).encode())

if __name__ == "__main__":
    solve()