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
input = lambda: sys.stdin.readline().split()
si = lambda: input()[0]
ii = lambda: int(si())
lii = lambda: list(map(int, input()))
lsi = lambda: input()

def solve():
    T = ii()
    out = []
    for _ in range(T):
        ans = 0
    
        out.append(str(ans))
    sys.stdout.buffer.write('\n'.join(out).encode())

if __name__ == "__main__":
    solve()