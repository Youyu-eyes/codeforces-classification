# https://codeforces.com/contest/1526/problem/B
# 1400

si = lambda: input()
ii = lambda: int(input())
many_ii = lambda: map(int, input().split())
many_si = lambda: input().split()
lii = lambda: list(many_ii())
lsi = lambda: list(many_si())
MOD = 1_000_000_007

def solve(nums):
    if nums % 11 == 0 or nums % 111 == 0:
        return True
    n = nums // 111
    for i in range(n + 1):
        if (nums - i * 111) % 11 == 0:
            return True
    return False

T = ii()
for _ in range(T):
    n = ii()
    print('YES' if solve(n) else 'NO')