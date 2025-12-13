# https://codeforces.com/problemset/problem/2121/D

from math import inf
from math import gcd
from math import lcm
from math import sin
from math import pi
from heapq import heapify
from heapq import heappush
from heapq import heappop
from heapq import heapreplace
from heapq import nsmallest
from heapq import nlargest
from functools import cache
from itertools import accumulate
from itertools import zip_longest
from itertools import pairwise
from bisect import bisect_left
from bisect import bisect_right
from collections import defaultdict
from collections import Counter
from collections import deque
import datetime
import re
from typing import List, Tuple
si = lambda: input()
ii = lambda: int(input())
many_ii = lambda: map(int, input().split())
many_si = lambda: input().split()
lii = lambda: list(many_ii())
lsi = lambda: list(many_si())
MOD = 1_000_000_007

# 官解先排序再交换，但是有可能先交换的次数最小，但是题目也没说最小
# 比较一下先交换还是先排序，暴力的思想直接模拟
# 按理来说两种方案一个能过就行，但是只有一个能过

def solve1(a, b):
    n = len(a)
    ans = []
    for i, (x, y) in enumerate(zip(a, b)):
        if x > y:
            a[i], b[i] = b[i], a[i]
            ans.append((3, i + 1))

    def bubblesort(nums, k):
        nonlocal ans
        for i in range(n - 1):
            for j in range(n - 1, i, -1):
                if nums[j - 1] > nums[j]:
                    nums[j - 1], nums[j] = nums[j], nums[j - 1]
                    ans.append((k, j))
    
    bubblesort(a, 1)
    bubblesort(b, 2)
    
    return ans

def solve2(a, b):
    n = len(a)
    ans = []

    def bubblesort(nums, k):
        nonlocal ans
        for i in range(n - 1):
            for j in range(n - 1, i, -1):
                if nums[j - 1] > nums[j]:
                    nums[j - 1], nums[j] = nums[j], nums[j - 1]
                    ans.append((k, j))

    bubblesort(a, 1)
    bubblesort(b, 2)

    for i, (x, y) in enumerate(zip(a, b)):
        if x > y:
            ans.append((3, i + 1))
    
    return ans

T = ii()
for _ in range(T):
    n = ii()
    a = lii()
    b = lii()
    a1 = a.copy()
    b1 = b.copy()
    ans1 = solve1(a1, b1)
    ans2 = solve2(a, b)

    if len(ans1) > len(ans2):
        ans = ans2
    else:
        ans = ans1
    print(len(ans))
    for k, i in ans:
        print(k, i)