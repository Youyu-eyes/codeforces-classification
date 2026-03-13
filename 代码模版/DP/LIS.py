from bisect import bisect_left
from typing import List, Tuple

class Solution:
    def lengthOfLIS(self, nums: List[int]) -> int:
        g = []
        for x in nums:
            j = bisect_left(g, x)
            if j == len(g):  # >=x 的 g[j] 不存在
                g.append(x)
            else:
                g[j] = x
        return len(g)