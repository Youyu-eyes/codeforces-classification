# from math import inf
# from math import gcd
# from math import lcm
# from math import sin
# from math import pi

# from heapq import heapify
# from heapq import heappush
# from heapq import heappop
# from heapq import heapreplace
# from heapq import nsmallest
# from heapq import nlargest

# from heapq import heapify_max
# from heapq import heappush_max
# from heapq import heappop_max
# from heapq import heapreplace_max
# from heapq import nsmallest_max
# from heapq import nlargest_max

# from functools import cache

# from itertools import accumulate
# from itertools import zip_longest
# from itertools import pairwise
# from itertools import combinations
# from itertools import permutations

# from bisect import bisect_left

# from collections import defaultdict
# from collections import Counter
# from collections import deque

from datetime import *
# import re
# from typing import List, Tuple

# MOD = 1_000_000_007

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

for _ in range(II()):
    todo_date, todo_time = input().split()
    d0 = datetime.strptime(todo_date, "%Y/%m/%d").date()
    hour, minute, second = map(int, todo_time.split(':'))
    total = hour * 10000 + minute * 100 + second
    d1 = d0 + timedelta(total // 86400)
    total %= 86400
    t1 = time(total // 3600, (total - total // 3600 * 3600) // 60, (total - total // 3600 * 3600) % 60)
    ans = datetime.combine(d1, t1)
    print(ans.strftime("%Y/%m/%d %H:%M:%S"))