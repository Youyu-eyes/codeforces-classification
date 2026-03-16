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


# 离散化
def discretize(arr):
    sorted_unique = sorted(set(arr))
    value_to_index = {value: idx for idx, value in enumerate(sorted_unique)}
    discretized = [value_to_index[x] for x in arr]
    return discretized


# 凸包
class Vec:
    __slots__ = 'x', 'y'

    def __init__(self, x: int, y: int):
        self.x = x
        self.y = y

    def __sub__(self, b: "Vec") -> "Vec":
        return Vec(self.x - b.x, self.y - b.y)

    def det(self, b: "Vec") -> int:
        return self.x * b.y - self.y * b.x

# Andrew 算法，计算 points 的凸包（逆时针顺序）
# 时间复杂度 O(n log n)，其中 n = len(points)
def convexHull(points: List[Vec]) -> List[Vec]:
    if len(points) <= 1:
        return points

    points.sort(key=lambda p: (p.x, p.y))

    q = []

    # 计算下凸包（从左到右）
    for p in points:
        # 新来的点 p，能否让旧的点变成在凸包内的点？ ->  需要判断向量左右关系  ->  det
        while len(q) > 1 and (q[-1] - q[-2]).det(p - q[-1]) <= 0:
            q.pop()
        q.append(p)

    # 计算上凸包（从右到左）
    # 注意下凸包的最后一个点，已经是上凸包的（右边）第一个点了，所以从 n-2 开始遍历
    lower_size = len(q)
    for i in range(len(points) - 2, -1, -1):
        p = points[i]
        while len(q) > lower_size and (q[-1] - q[-2]).det(p - q[-1]) <= 0:
            q.pop()
        q.append(p)

    # 此时首尾是同一个点 points[0]，需要去掉
    q.pop()

    return q

# 快速幂
def myPow(x: float, n: int) -> float:
    ans = 1
    if n < 0:
        n = -n
        x = 1 / x
    while n:
        if n & 1:
            ans *= x
        x *= x
        n >>= 1
    return ans

# 矩阵快速幂
# a @ b，其中 @ 是矩阵乘法
def mul(a: List[List[int]], b: List[List[int]]) -> List[List[int]]:
    return [[sum(x * y for x, y in zip(row, col)) for col in zip(*b)]
            for row in a]

# a^n @ f0
def pow_mul(a: List[List[int]], n: int, f0: List[List[int]]) -> List[List[int]]:
    res = f0
    while n:
        if n & 1:
            res = mul(a, res)
        a = mul(a, a)
        n >>= 1
    return res