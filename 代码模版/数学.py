from math import isqrt
from typing import List, Tuple
# 根号判断素数
def is_prime(n: int) -> bool:
    for i in range(2, isqrt(n) + 1):
        if n % i == 0:
            return False
    return n >= 2

# 埃氏筛
MX = 1_000_001
is_prime = [False] * 2 + [True] * (MX - 2)
primes = []
for i in range(2, MX):
    if is_prime[i]:
        primes.append(i)
        for j in range(i * i, MX, i):
            is_prime[j] = False

# 分解质因数
MX = 1_000_001
lpf = [0] * MX
for i in range(2, MX):
    if lpf[i] == 0:  # i 是质数
        for j in range(i, MX, i):
            if lpf[j] == 0:  # 首次访问 j
                lpf[j] = i

def prime_factorization(x: int) -> List[Tuple[int, int]]:
    res = []
    while x > 1:
        p = lpf[x]
        e = 1
        x //= p
        while x % p == 0:
            e += 1
            x //= p
        res.append((p, e))
    return res

# 预处理每个数的因子
MX = 1_000_001  # **根据题目调整**
divisors = [[] for _ in range(MX)]
for i in range(1, MX):
    for j in range(i, MX, i):  # 枚举 i 的倍数 j
        divisors[j].append(i)  # i 是 j 的因子