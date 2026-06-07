from math import inf, gcd, lcm, sin, pi, sqrt, isqrt

from heapq import heapify, heappush, heappop, heapreplace, nsmallest, nlargst

from heapq import heapify_max, heappush_max, heappop_max, heapreplace_max, nsmallest_max, nlargest_max

from functools import cache, cmp_to_key

from itertools import accumulate, zip_longest, pairwise, combinations, permutations, groupby

from bisect import bisect_left, bisect_right

from collections import defaultdict, Counter, deque

from typing import List, Tuple

MOD = 1_000_000_007

from sortedcontainers import SortedSet, SortedDict, SortedList
from datetime import *
t1 = date(year = 2025, month = 3, day = 1)
t2 = date(year = 2025, month = 3, day = 17)
print(t2 - t1)  # 16 days, 0:00:00
print((t2 - t1).days)  # 16

# 参数也可简化
t1 = date(2025, 1, 1)
t2 = date(2025, 3, 17)
print((t2 - t1).days)  # 75

# 获取当前日期
print(date.today())

t1 = date(2025, 3, 17)
print(t1 + timedelta(days = 2))  # 2025-03-19
print(t1 + timedelta(1))  # 2025-03-18
print(t1.weekday())  # 0-6, 0代表周一

# bitwise 注意点
# Python  ^ 运算 比 & 运算低一级

# 自定义排序
# https://leetcode.cn/problems/maximum-value-of-concatenated-binary-segments/description/
# compare 函数需要接受 a, b 两个参数，然后返回 -1/0/1
nums = []
def compare(a, b):
    if a < b:
        return -1 # 表示 a 排在 b 之前
    if a == b:
        return 0 # 表示 a 与 b 顺序无所谓
    if a > b:
        return 1 # 表示 a 排在 b 之后
nums.sort(key = cmp_to_key(compare))