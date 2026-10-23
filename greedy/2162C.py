# https://codeforces.com/contest/2162/problem/C
# 1100

import sys
input = lambda: sys.stdin.readline().strip()
si = lambda: input()
ii = lambda: int(input())
many_ii = lambda: map(int, input().split())
li = lambda: list(many_ii)
MOD = 1_000_000_007

def solve(num, target):
    if num.bit_length() < target.bit_length():
        return 0, (0, 0)
    if num == target:
        return 1, (0, 0)
    if num.bit_length() == target.bit_length():
        return 2, (num ^ target, 0)
    return 3, ((num - pow(2, num.bit_length() - 1)) ^ target, pow(2, num.bit_length() - 1))
    

T = ii()
for _ in range(T):
    num, target = many_ii()
    flag, ans = solve(num, target)
    if flag == 0:
        print(-1)
    elif flag == 1:
        print(0)
    elif flag == 2:
        print(1)
        print(ans[0])
    else:
        print(2)
        print(ans[0], ans[1])