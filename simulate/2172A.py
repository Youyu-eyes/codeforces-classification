# https://codeforces.com/contest/2172/problem/A
# 800

import sys
input = lambda: sys.stdin.readline().strip()

lst = sorted(list(map(int, input().split())))
if lst[-1] - lst[0] < 10:
    print(f'final {lst[1]}')
else:
    print('check again')