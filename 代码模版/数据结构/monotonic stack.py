from collections import deque
from typing import List, Tuple

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