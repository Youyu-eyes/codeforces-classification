from collections import deque
from typing import List, Tuple

# 找到 i 左端第一个严格大于 nums[i] 的数的下标
def nearestGreater(nums: List[int]) -> Tuple[List[int], List[int]]:
    n = len(nums)
    # left[i] 是 nums[i] 左侧最近的严格大于 nums[i] 的数的下标，若不存在则为 -1
    left = [-1] * n
    st = []
    for i, x in enumerate(nums):
        while st and nums[st[-1]] <= x:  # 如果求严格小于，改成 >=
            st.pop()
        if st:
            left[i] = st[-1]
        st.append(i)

    # right[i] 是 nums[i] 右侧最近的严格大于 nums[i] 的数的下标，若不存在则为 n
    right = [n] * n
    st = []
    for i in range(n - 1, -1, -1):
        x = nums[i]
        while st and nums[st[-1]] <= x:  # 如果求严格小于，改成 >=
            st.pop()
        if st:
            right[i] = st[-1]
        st.append(i)

    return left, right

# 找到离 i 最远的，比 nums[i] 小的数
def longestWPI(self, hours: List[int]) -> int:
    n = len(hours)
    s = [0] * (n + 1)  # 前缀和
    st = [0]  # s[0]
    for j, h in enumerate(hours, 1):
        s[j] = s[j - 1] + (1 if h > 8 else -1)
        if s[j] < s[st[-1]]: st.append(j)  # 感兴趣的 j
    ans = 0
    for i in range(n, 0, -1):
        while st and s[i] > s[st[-1]]:
            ans = max(ans, i - st.pop())  # [st[-1],i) 可能是最长子数组
    return ans