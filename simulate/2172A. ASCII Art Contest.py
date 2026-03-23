# https://codeforces.com/contest/2172/problem/A
# 800

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
input = lambda: sys.stdin.readline().strip()

lst = sorted(list(map(int, input().split())))
if lst[-1] - lst[0] < 10:
    print(f'final {lst[1]}')
else:
    print('check again')