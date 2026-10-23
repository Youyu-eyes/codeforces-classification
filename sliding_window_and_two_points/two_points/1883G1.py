# https://codeforces.com/problemset/problem/1883/G1
# 1400
# 双指针

si = lambda: input()
ii = lambda: int(input())
many_ii = lambda: map(int, input().split())
many_si = lambda: input().split()
lii = lambda: list(many_ii())
lsi = lambda: list(many_si())
MOD = 1_000_000_007

def solve(n, a, b):
    a.sort()
    b.sort()
    p1 = p2 = 0
    while p2 < n:
        if a[p1] < b[p2]:
            p1 += 1
            p2 += 1
        else:
            p2 += 1
    return n - p1

T = ii()
for _ in range(T):
    n, m = many_ii()
    a = lii()
    b = lii()
    a = [1] + a
    print(solve(n, a, b))