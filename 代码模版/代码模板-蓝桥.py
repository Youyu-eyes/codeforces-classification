# from math import inf
# from math import gcd
# from math import sin
# from math import pi

# from heapq import heapify
# from heapq import heappush
# from heapq import heappop
# from heapq import heapreplace
# from heapq import nsmallest
# from heapq import nlargest

# from itertools import accumulate
# from itertools import zip_longest

# from bisect import bisect_left

# from collections import defaultdict
# from collections import Counter
# from collections import deque

# import datetime
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
    return list(MII())

for _ in range(II()):
    print()