from typing import List, Tuple

# 对于每个右端点 i，计算所有子数组的或值，打印这些或值的分布范围（子数组左端点范围）
# 时间复杂度 O(nlogU)，其中 U = max(nums)
def logTrick(nums: List[int]) -> None:
    or_left = []  # (子数组或值，最小左端点)
    for i, x in enumerate(nums):
        # 计算以 i 为右端点的子数组或值
        for p in or_left:
            p[0] |= x  # **根据题目修改**
        # x 单独一个数作为子数组
        or_left.append([x, i])

        # 原地去重（相同或值只保留最左边的）
        idx = 1
        for j in range(1, len(or_left)):
            if or_left[j][0] != or_left[j - 1][0]:
                or_left[idx] = or_left[j]
                idx += 1
        del or_left[idx:]

        print(i, x)
        for k, (or_val, left) in enumerate(or_left):
            right = or_left[k + 1][1] - 1 if k < len(or_left) - 1 else i
            # 对于左端点在 [left, right]，右端点为 i 的子数组，OR 值都是 or_val
            print(left, right, or_val)