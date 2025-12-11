# https://codeforces.com/contest/2162/problem/C

import sys
from math import inf
from math import gcd
from math import lcm
from math import sin
from math import pi
import heapq as h
from functools import cache
from itertools import accumulate
from itertools import zip_longest
from bisect import bisect_left
from bisect import bisect_right
from collections import defaultdict
from collections import Counter
from collections import deque as q
import datetime
import re
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